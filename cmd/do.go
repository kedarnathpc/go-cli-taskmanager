package cmd

import (
	"fmt"
	"strconv"

	"github.com/kedarnathpc/task/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks as task complete.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {

			// convert the string arguments into integers
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument.")
			} else {

				// append into id slice
				ids = append(ids, id)
			}
		}

		// get the list of tasks from the db
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Printf("Invalid task number: %d", id)
				continue
			}
			task := tasks[id-1]

			// delete task function from the db
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error:%s", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
