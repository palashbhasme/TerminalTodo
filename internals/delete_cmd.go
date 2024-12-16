package internals

import (
	"flag"
	"fmt"
	"log"
)

func DelCmd(todo_list *Todos, args []string) {

	cmdDel := flag.NewFlagSet("delete", flag.ExitOnError)
	cmdDelTask := cmdDel.Int("id", -1, "enter the id of task you want to delete")

	cmdDel.Parse(args)

	if *cmdDelTask == -1 {
		log.Println("Error: the --del flag is required for the 'delete' subcommand.")
	}

	todo_list.Delete(*cmdDelTask)
	todo_list.Store(filePath)

	fmt.Println("Task deleted successfully!")
}
