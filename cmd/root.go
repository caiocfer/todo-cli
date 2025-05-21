package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
}

var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "Todocli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from Cobra")

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)

	}

}
