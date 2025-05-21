package cmd

import (
	"os"

	"github.com/caiocfer/todo_cli/todo"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List items in your todo list",
	Run:   listItems,
}

func listItems(cmd *cobra.Command, args []string) {
	data := [][]string{
		{"ID", "Name", "Completed"},
	}
	todo.ListItems()
	table := tablewriter.NewWriter(os.Stdout)
	table.Header(data[0])
	table.Bulk(todo.ItemsList)
	table.Render()
}
