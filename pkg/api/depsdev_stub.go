package api

import "github.com/edoardottt/depsdev/pkg/depsdev"

type DepsdevRepositoryStub struct {
}

func (d DepsdevRepositoryStub) GetAdvisoryKeys(name, ver string) []depsdev.AdvisoryKeys {
	return []depsdev.AdvisoryKeys{{ID: "Key1"}, {ID: "Key2"}}
}

func (d DepsdevRepositoryStub) GetAdvisory(ad string) depsdev.Advisory {
	return depsdev.Advisory{
		AdvisoryKey: depsdev.AdvisoryKey{ID: "Key1"},
		URL:         "http://stub002",
		Title:       "stub003",
		Aliases: []string{
			"stub004",
			"stub005",
		},
		Cvss3Score:  1.0,
		Cvss3Vector: "stub006",
	}
}
