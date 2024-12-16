package internals

import (
	"fmt"
)

type MyError struct {
	action  string
	Message string
	taskId  int
}

func (err *MyError) Error() string {

	error := fmt.Sprintf("Error: %s on action: %s, task id %d", err.Message, err.action, err.taskId)
	return error
}
