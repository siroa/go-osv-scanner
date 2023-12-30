/*
go-osv-scanner - CLI client to discover vulnerable modules.

@author: siroa

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/siroa/go-osv-scanner/blob/main/LICENSE
*/
package format

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type Osv struct {
	AdvisoryKey string
	Title       string
	Url         string
	Score       string
	Vector      string
}

var Osvs []Osv

func NewOsv(key, title, url, score, vec string) Osv {
	return Osv{
		AdvisoryKey: key,
		Title:       title,
		Url:         url,
		Score:       score,
		Vector:      vec,
	}
}

func (o Osv) AddOsvs() {
	Osvs = append(Osvs, o)
}

func GetOsvs() [][]string {
	if len(Osvs) == 0 {
		return nil
	}
	data := [][]string{}
	for _, v := range Osvs {
		data = append(data, []string{v.AdvisoryKey, v.Title, v.Url, v.Score, v.Vector})
	}
	Osvs = []Osv{}
	return data
}

func PrintTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"advisoryKey", "title", "url", "cvss3Score", "cvss3Vector"})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}
