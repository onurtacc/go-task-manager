package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"go-task-manager/cmd"
	"go-task-manager/db"
	"os"
	"path/filepath"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
