package internals

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"taskmanager/internals/utils"
	"time"

	"github.com/alexeyco/simpletable"
)

type todo struct {
	ID          int        `json:"id"`
	Task        string     `json:"task"`
	Done        bool       `json:"done"`
	Category    string     `json:"category"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type Todos []todo

var CurrId int

// add todo to Todos slice
func (t *Todos) Add(task, category string) {

	new_todo := todo{
		ID:          CurrId,
		Task:        task,
		Done:        false,
		Category:    category,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*t = append(*t, new_todo)
	CurrId++

}

func (t *Todos) Delete(id int) {

	index := t.todoById(id)

	if index != -1 {
		*t = append((*t)[0:index], (*t)[index+1:]...)
	} else {
		log.Fatalln("No such task exists")
	}

}
func (t *Todos) UpdateTask(id int, name *string, category *string, done *bool) error {

	index := t.todoById(id)

	if name != nil {
		(*t)[index].Task = *name
		log.Println(*name)
	}

	if category != nil {
		(*t)[index].Category = *category
		log.Println(*category)

	}

	if done != nil {
		now := time.Now()
		if *done {
			(*t)[index].CompletedAt = &now
		} else {
			(*t)[index].CompletedAt = nil

		}
		(*t)[index].Done = *done
		log.Println(*done)

	}

	return nil
}

func (t *Todos) Store(filename string) {
	data, err := json.Marshal(t)
	if err != nil {
		log.Fatalln("Error in marshalling")
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatalln("Error writing to todos file")
	}

}

func LoadData(t *Todos) {

	if utils.IsLoaded {
		return
	} else {

		filePath := utils.GetFilePath()

		if emptyFlag := IsEmpty(filePath); emptyFlag {

			utils.IsLoaded = true
			return
		}

		data, err := os.ReadFile(filePath)

		if err != nil {
			log.Fatalln("Error retrieving todos file: ", err)
		}

		err = json.Unmarshal(data, t)
		if err != nil {
			log.Fatalln("Error unmarshalling todos file:", err)
		}

		if len(*t) > 0 {
			maxId := (*t)[0].ID
			for _, todo := range *t {
				if todo.ID > maxId {
					maxId = todo.ID
				}
			}
			CurrId = maxId + 1
		}

		utils.IsLoaded = true
	}
}

// return index for given todo id
func (t *Todos) todoById(id int) int {

	ls := *t

	if id < 0 {
		err := MyError{
			action:  "get task Id",
			Message: "Error getting task Id",
			taskId:  id,
		}

		log.Fatalln("Invalid taskid error", err.Error())
	}
	//current index
	inx := -1
	for index, task := range ls {
		if id == task.ID {
			inx = index
		}
	}

	if inx == -1 {
		err := MyError{
			action:  "get task Id",
			Message: "Error getting task Id",
			taskId:  id,
		}

		log.Fatalln("Invalid taskid error", err.Error())
	}

	return inx
}

func (t *Todos) PrintTable() {
	if len(*t) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Category"},
			{Align: simpletable.AlignCenter, Text: "Done"},
			{Align: simpletable.AlignCenter, Text: "Created At"},
			{Align: simpletable.AlignCenter, Text: "Completed At"},
		},
	}

	completed := 0
	notCompleted := 0

	for _, todo := range *t {
		completedAt := "N/A"
		if todo.CompletedAt != nil {
			completedAt = todo.CompletedAt.Format("02 Jan 2006 15:04:05")
		}

		row := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%d", todo.ID)},
			{Align: simpletable.AlignLeft, Text: todo.Task},
			{Align: simpletable.AlignLeft, Text: todo.Category},
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%v", todo.Done)},
			{Align: simpletable.AlignLeft, Text: todo.CreatedAt.Format("02 Jan 2006 15:04:05")},
			{Align: simpletable.AlignLeft, Text: completedAt},
		}

		table.Body.Cells = append(table.Body.Cells, row)

		if todo.Done {
			completed++
		} else {
			notCompleted++
		}
	}

	// Footer to match header columns
	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: " "}, // Align footer with header
			{Align: simpletable.AlignRight, Text: "Completed:"},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%d", completed)},
			{Align: simpletable.AlignRight, Text: "Not Completed:"},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%d", notCompleted)},
			{Align: simpletable.AlignRight, Text: " "},
		},
	}

	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())
}
