package api

import "github.com/edoardottt/depsdev/pkg/depsdev"

type DepsdevRepositoryStub struct {
}

func (d DepsdevRepositoryStub) GetAdvisoryKeys(name, ver string) []depsdev.AdvisoryKeys {
	return []depsdev.AdvisoryKeys{{ID: "Key1"}, {ID: "Key2"}}
}
