package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of this",
	Run:   SayHello,
}

func SayHello(cmd *cobra.Command, args []string) {
	fmt.Println("v0.0.1")
}
