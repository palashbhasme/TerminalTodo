package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var IsLoaded bool

func GetFilePath() string {
	var homeDir string

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("error in retrieving user home directory")
	}
	fullPath := filepath.Join(homeDir, "todos.json")

	return fullPath

}
func GetApproval() {
	var ans string
	fmt.Println("Task Manager would like access to you home dir to store .todos file Y/n?")

	fmt.Scan(&ans)

	if ans == "Y" {
		return
	} else {
		fmt.Println("Permission denied, exited")
		os.Exit(1)
	}

}
