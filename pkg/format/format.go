/*
go-osv-scanner - CLI client to discover vulnerable modules.

@author: siroa

@license: https://github.com/siroa/go-osv-scanner/blob/main/LICENSE
*/
package format

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
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
	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRenderer(renderer.NewBlueprint(tw.Rendition{
			Settings: tw.Settings{Separators: tw.Separators{BetweenRows: tw.On}},
		})),
		tablewriter.WithConfig(tablewriter.Config{
			Row: tw.CellConfig{
				Formatting:   tw.CellFormatting{AutoWrap: tw.WrapNormal}, // Wrap long content
				Alignment:    tw.CellAlignment{Global: tw.AlignLeft},     // Left-align rows
				ColMaxWidths: tw.CellWidth{Global: 25},
			},
		}),
	)
	table.Header([]string{"advisoryKey", "title", "url", "cvss3Score", "cvss3Vector"})
	table.Bulk(data)
	table.Render()
}
