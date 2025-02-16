package cmd

import (
	"fmt"

	"cli-file-mgmt/services"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [filename]",
	Short: "Delete a file with this command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		err := services.DeleteFile(filename)
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
