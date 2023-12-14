package cmd

import (
	"fmt"
	"log"
	"os"
	"scanner/pkg/gomod"
	"time"

	"github.com/spf13/cobra"
)

var line = &cobra.Command{
	Use:   "scanner",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		modFile, err := cmd.Flags().GetString("mod")
		if err != nil {
			log.Fatalln(err)
		}
		if modFile == "" {
			log.Fatalln("Please specify go.mod file")
		}
		file, err := os.ReadFile(modFile)
		if err != nil {
			log.Fatalln(err)
		}
		gm := gomod.ParseGoMod(file)
		fmt.Printf("Your module name: %s\n", gm.Name)
		for _, v := range gm.Modules {
			v.SetAdvisoryKeys()
			v.PrintModule()
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
}
