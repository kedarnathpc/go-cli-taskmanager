package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// this is the root command
// It only prints out what the command does which is tells what is cli executable does
var RootCmd = &cobra.Command{
	// name of the root command
	Use: "task",
	// short description of the command
	Short: "Task is a CLI TODO manager.",
}

// function to execute the root commmand when called
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
