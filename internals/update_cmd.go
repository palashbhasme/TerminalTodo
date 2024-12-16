package internals

import (
	"flag"
	"log"
	"taskmanager/internals/utils"
)

func UpdCmd(todo_list *Todos, args []string) {
	cmdUpdate := flag.NewFlagSet("update", flag.ExitOnError)
	cmdId := cmdUpdate.Int("id", -1, "enter the id of the task you want to update")
	cmdDone := cmdUpdate.String("done", "", "update task state (true/false)")
	cmdName := cmdUpdate.String("name", "", "update name of the task")
	cmdCat := cmdUpdate.String("cat", "", "update category of the task")

	cmdUpdate.Parse(args)

	// Validate task ID
	if *cmdId == -1 {
		log.Fatalln("Error: the --id flag is required for the 'update' subcommand.")
	}

	// Validate that at least one update field is provided
	if len(*cmdName) == 0 && len(*cmdCat) == 0 && len(*cmdDone) == 0 {
		log.Fatalln("Error: You must provide at least one of --name, --cat, or --done to update a task.")
	}

	// Validate the --done flag
	var doneVal *bool // Pointer to handle the optional --done flag
	if len(*cmdDone) > 0 {
		switch *cmdDone {
		case "true":
			trueVal := true
			doneVal = &trueVal
		case "false":
			falseVal := false
			doneVal = &falseVal
		default:
			log.Fatalln("Error: Invalid value for --done. Only 'true' or 'false' are allowed.")
		}
	}

	// Perform updates based on the provided flags
	if len(*cmdName) != 0 {
		err := todo_list.UpdateTask(*cmdId, cmdName, nil, nil)
		if err != nil {
			log.Println("Error updating name:", err)
			log.Fatalln(err)
		}
	}

	if len(*cmdCat) != 0 {
		err := todo_list.UpdateTask(*cmdId, nil, cmdCat, nil)
		if err != nil {
			log.Println("Error updating category:", err)
			log.Fatalln(err)
		}
	}

	if doneVal != nil {
		err := todo_list.UpdateTask(*cmdId, nil, nil, doneVal)
		if err != nil {
			log.Println("Error updating done status:", err)
			log.Fatalln(err)
		}
	}

	// Save changes
	todo_list.Store(utils.GetFilePath())
}
