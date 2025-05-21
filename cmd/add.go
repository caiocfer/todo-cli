package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add item to todo list",
	Run:   AddItem,
}

func AddItem(cmd *cobra.Command, args []string) {
	fmt.Println("Adding item #Todo")
}
