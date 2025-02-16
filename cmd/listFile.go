package cmd

import (
	"fmt"

	"cli-file-mgmt/services"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the files and directories.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := services.ListFiles()
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
