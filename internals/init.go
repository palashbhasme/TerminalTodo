package internals

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"taskmanager/internals/utils"
)

func Init() {

	utils.GetApproval()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("Error retrieving usser home directory: ", err)
	}

	filePath := filepath.Join(homeDir, "todos.json")

	_, err = os.Stat(filePath)

	if err != nil {

		if os.IsNotExist(err) {

			_, err := os.Create(filePath)

			if err != nil {
				log.Println("Error creating File: ", err)
			}
		}
	} else {
		fmt.Println("File already exsits, previous todos are loaded")
	}

	fmt.Println("App initialized successfully!")

}

func IsEmpty(filePath string) bool {

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Println("Error retrieving file info", err)
	}

	return fileInfo.Size() == 0
}
