/*

go-osv-scanner - CLI client to discover vulnerable modules.

@author: siroa

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/siroa/go-osv-scanner/blob/main/LICENSE

*/

package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/siroa/go-osv-scanner/pkg/api"
	"github.com/siroa/go-osv-scanner/pkg/format"
	"github.com/siroa/go-osv-scanner/pkg/gomod"

	"github.com/spf13/cobra"
)

var line = &cobra.Command{
	Use:   "scanner",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			log.Fatalln("verbose flag error: ", err)
		}
		modFile, err := cmd.Flags().GetString("mod")
		if err != nil {
			log.Fatalln("go.mod flag error: ", err)
		}
		if modFile == "" {
			log.Fatalln("Please specify go.mod file")
		}
		file, err := os.ReadFile(modFile)
		if err != nil {
			log.Fatalln("ReadFile error: ", err)
		}
		gm := gomod.ParseGoMod(file)
		fmt.Printf("Your module name: %s\n", gm.Name)
		for _, v := range gm.Modules {
			v.SetAdvisoryKeys(api.Depsdev{})
			if verbose && len(v.AdvisoryIDs) != 0 {
				v.SetAdvisory(api.Depsdev{})
				osvs := format.GetOsvs()
				format.PrintTable(osvs)
			} else {
				v.PrintModule()
			}
			time.Sleep(100 * time.Millisecond)
		}
	},
}

func Execute() {
	err := line.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	line.Flags().StringP("mod", "m", "", "Specify the path to the go.mod file")
	line.Flags().BoolP("verbose", "v", false, "Output vulnerability details")
}
