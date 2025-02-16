package cmd

import (
	"fmt"

	"cli-file-mgmt/services"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [filename] [content]",
	Short: "Create a new file with content ",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		content := args[1]
		err := services.CreateFile(filename, content)
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
