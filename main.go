package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kedarnathpc/task/cmd"
	"github.com/kedarnathpc/task/db"

	"github.com/mitchellh/go-homedir"
)

func main() {

	// gets the home directory path
	home, _ := homedir.Dir()

	// passes the home dir to make database at home dir
	dbPath := filepath.Join(home, "tasks.db")
	checkError(db.Init(dbPath))
	cmd.Execute()
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
