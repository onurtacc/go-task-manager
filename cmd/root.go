package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "go-task-manager",
	Short: "go-task-manager is a CLI for managing your TODOs.",
}
