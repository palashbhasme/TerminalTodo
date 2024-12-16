package internals

import (
	"flag"
	"fmt"
	"log"
	"taskmanager/internals/utils"
)

var filePath = utils.GetFilePath()

func AddCmd(todo_list *Todos, args []string) {
	cmdAdd := flag.NewFlagSet("add", flag.ExitOnError)
	cmdAddTask := cmdAdd.String("task", "", "enter a task name you want to track")
	cmdAddCat := cmdAdd.String("cat", "Uncategorized", "enter the task category")

	cmdAdd.Parse(args)

	if len(*cmdAddTask) == 0 {
		log.Fatalln("Error: the --task flag is required for the 'add' subcommand.")
	}

	todo_list.Add(*cmdAddTask, *cmdAddCat)
	todo_list.Store(filePath)

	fmt.Println("Task added successfully!")

}
