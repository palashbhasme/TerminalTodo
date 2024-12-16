package main

import (
	"flag"
	"log"
	"os"
	"taskmanager/internals"
)

func main() {

	//list of todos
	var todos internals.Todos
	flag.Parse()

	switch flag.Arg(0) {
	case "init":
		internals.Init()
	case "add":
		internals.LoadData(&todos)
		internals.AddCmd(&todos, os.Args[2:])
	case "delete":
		internals.LoadData(&todos)
		internals.DelCmd(&todos, os.Args[2:])
	case "update":
		internals.LoadData(&todos)
		internals.UpdCmd(&todos, os.Args[2:])
	case "show":
		internals.LoadData(&todos)
		todos.PrintTable()

	default:
		log.Printf("Command %s not found did you mean: \n", flag.Arg(0))
		log.Printf("\n command \"init\" \n command \"add\" \n command \"delete\" \n command \"update\"")
	}

}
