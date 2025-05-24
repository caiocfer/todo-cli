package cmd

import (
	"github.com/caiocfer/todo_cli/localfile"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List items in your todo list",
	Run:   listItems,
}

func listItems(cmd *cobra.Command, args []string) {
	localfile.ListItems()
}
