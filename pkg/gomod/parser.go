/*
go-osv-scanner - CLI client to discover vulnerable modules.

@author: siroa

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/siroa/go-osv-scanner/blob/main/LICENSE
*/
package gomod

import (
	"fmt"
	"log"
	"strconv"

	"github.com/siroa/go-osv-scanner/pkg/api"
	"github.com/siroa/go-osv-scanner/pkg/format"

	"golang.org/x/mod/modfile"
)

type GoMod struct {
	Name    string
	Modules []Module
}

func NewGoMod(name string) *GoMod {
	return &GoMod{
		Name:    name,
		Modules: []Module{},
	}
}

func (g *GoMod) setModules(ms []Module) {
	g.Modules = ms
}

type Module struct {
	Name        string
	Version     string
	AdvisoryIDs []string
}

func NewModule(name, ver string) *Module {
	return &Module{
		Name:    name,
		Version: ver,
	}
}

func (m *Module) SetAdvisoryKeys(depsdev api.DepsdevRepository) {
	var keys []string
	adkeys := depsdev.GetAdvisoryKeys(m.Name, m.Version)
	if len(adkeys) == 0 {
		fmt.Printf("No vulnerabilities were found in %s:%s\n", m.Name, m.Version)
		return
	}
	for _, v := range adkeys {
		keys = append(keys, v.ID)
	}
	m.AdvisoryIDs = keys
}

func (m Module) SetAdvisory(depsdev api.DepsdevRepository) {
	fmt.Printf("Vulnerability Detection!: %s:%s\n", m.Name, m.Version)
	for _, ad := range m.AdvisoryIDs {
		a := depsdev.GetAdvisory(ad)
		score := strconv.FormatFloat(a.Cvss3Score, 'b', 2, 64)
		o := format.NewOsv(a.AdvisoryKey.ID, a.Title, a.URL, score, a.Cvss3Vector)
		o.AddOsvs()
	}
}

func (m Module) PrintModule() {
	if len(m.AdvisoryIDs) == 0 {
		return
	}
	fmt.Printf("Vulnerability Detection!: %s:%s\n", m.Name, m.Version)
	for _, v := range m.AdvisoryIDs {
		fmt.Println(v)
	}
}

func ParseGoMod(file []byte) *GoMod {
	f, err := modfile.Parse("go.mod", file, nil)
	if err != nil {
		log.Fatalln(err)
	}

	gm := NewGoMod(f.Module.Mod.Path)
	var ms []Module
	for _, v := range f.Require {
		ms = append(ms, *NewModule(v.Mod.Path, v.Mod.Version))
	}
	gm.setModules(ms)

	return gm
}
