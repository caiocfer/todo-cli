package localfile

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func ListItems() {
	data := [][]string{
		{"ID", "Name", "Completed"},
	}
	items := ReadJson()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(data[0])
	table.Bulk(items)
	table.Render()
}
