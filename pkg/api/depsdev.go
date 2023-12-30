/*

go-osv-scanner - CLI client to discover vulnerable modules.

@author: siroa

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/siroa/go-osv-scanner/blob/main/LICENSE

*/

package api

import (
	"log"

	"github.com/edoardottt/depsdev/pkg/depsdev"
)

type DepsdevRepository interface {
	GetAdvisoryKeys(string, string) []depsdev.AdvisoryKeys
	GetAdvisory(string) depsdev.Advisory
}

type Depsdev struct {
}

// Get Avisorykeys from package name and version using depsdev
func (d Depsdev) GetAdvisoryKeys(name, ver string) []depsdev.AdvisoryKeys {
	client := depsdev.NewAPI()
	i, err := client.GetVersion("Go", name, ver)
	if err != nil {
		log.Fatalln(err)
	}

	return i.AdvisoryKeys
}

// ToDo: Get details on vulnerability information you find
func (d Depsdev) GetAdvisory(ad string) depsdev.Advisory {
	client := depsdev.NewAPI()
	i, err := client.GetAdvisory(ad)
	if err != nil {
		log.Fatalln(err)
	}
	return i
}
