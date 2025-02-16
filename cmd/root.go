package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "file-manager",
	Short: "A simple CLI file manager",
	Long:  "A Go-based CLI for file operations like creating, deleting, and running files.",
	Run: func(cmd *cobra.Command, args []string) {
		printCommandTable()
	},
}

func printCommandTable() {
	color.Set(color.FgYellow)
	fmt.Println("\nAvailable Commands:")
	color.Unset()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"COMMAND", "DESCRIPTION"})
	table.SetBorder(true)
	table.Append([]string{"create", "Create a new file"})
	table.Append([]string{"delete", "Delete a file"})
	table.Append([]string{"list", "List all files"})
	table.Append([]string{"run", "Execute a script or open a file"})
	table.Append([]string{"help", "Show help for commands"})

	table.Render()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Set(color.FgRed)
		fmt.Println("❌ Error:", err)
		color.Unset()
		os.Exit(1)
	}
}
