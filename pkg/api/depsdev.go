package api

import (
	"fmt"
	"log"

	"github.com/edoardottt/depsdev/pkg/depsdev"
)

func GetAdvisoryKeys(name, ver string) []depsdev.AdvisoryKeys {
	client := depsdev.NewAPI()
	i, err := client.GetVersion("Go", name, ver)
	if err != nil {
		log.Fatalln(err)
	}

	return i.AdvisoryKeys
}

// ToDo: Get details on vulnerability information you find
func GetAdvisory(ad string) {
	client := depsdev.NewAPI()
	i, err := client.GetAdvisory(ad)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(i)
}