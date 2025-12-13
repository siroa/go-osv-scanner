package api

import def "github.com/edoardottt/depsdev/pkg/depsdev/definitions"

type DepsdevRepositoryStub struct {
}

func (d DepsdevRepositoryStub) GetAdvisoryKeys(name, ver string) []def.AdvisoryKey {
	return []def.AdvisoryKey{{ID: "Key1"}, {ID: "Key2"}}
}

func (d DepsdevRepositoryStub) GetAdvisory(ad string) def.Advisory {
	return def.Advisory{
		AdvisoryKey: def.AdvisoryKey{ID: "Key1"},
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
