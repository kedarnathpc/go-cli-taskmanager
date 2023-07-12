package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kedarnathpc/task/db"

	"github.com/spf13/cobra"
)

// represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the task list.",

	// function to run when add command is called
	Run: func(cmd *cobra.Command, args []string) {

		// takes the arguments as a slice of strings and converts them into one string
		task := strings.Join(args, " ")

		// creates task at the db
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong.")
			os.Exit(1)
		}
		fmt.Printf("Added \"%s\" to your task list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
