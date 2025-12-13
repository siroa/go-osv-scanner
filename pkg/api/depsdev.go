/*

go-osv-scanner - CLI client to discover vulnerable modules.

@author: siroa

@license: https://github.com/siroa/go-osv-scanner/blob/main/LICENSE

*/

package api

import (
	"log"

	def "github.com/edoardottt/depsdev/pkg/depsdev/definitions"
	depsdev "github.com/edoardottt/depsdev/pkg/depsdev/v3"
)

type DepsdevRepository interface {
	GetAdvisoryKeys(string, string) []def.AdvisoryKey
	GetAdvisory(string) def.Advisory
}

type Depsdev struct {
}

// Get Avisorykeys from package name and version using depsdev
func (d Depsdev) GetAdvisoryKeys(name, ver string) []def.AdvisoryKey {
	client := depsdev.NewV3API()
	i, err := client.GetVersion("Go", name, ver)
	if err != nil {
		log.Fatalln(err)
	}
	return i.AdvisoryKeys
}

// ToDo: Get details on vulnerability information you find
func (d Depsdev) GetAdvisory(ad string) def.Advisory {
	client := depsdev.NewV3API()
	i, err := client.GetAdvisory(ad)
	if err != nil {
		log.Fatalln(err)
	}
	return i
}
