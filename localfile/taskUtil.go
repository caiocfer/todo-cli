package localfile

import (
	"fmt"

	"github.com/caiocfer/todo_cli/todo"
)

func GetTaskItem(taskId int) todo.TodoItem {

	tasks := ReadJson()
	fmt.Println(taskId)
	var taskItem todo.TodoItem
	foundItem := false
	for i := range tasks {
		if taskId == i {
			taskItem = tasks[i]
			foundItem = true
		}

	}
	if foundItem {
		return taskItem
	} else {
		return todo.TodoItem{}
	}

}

func CompleteTask(taskId int) []todo.TodoItem {
	tasks := ReadJson()
	for i := range tasks {
		if taskId == i {
			tasks[i].Completed = true
		}
	}

	return tasks
}
