package cmd

import (
	"fmt"

	"cli-file-mgmt/services"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [filename] [extension]",
	Short: "Run a script with this command",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		extension := args[1]
		err := services.RunFile(filename, extension)
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
