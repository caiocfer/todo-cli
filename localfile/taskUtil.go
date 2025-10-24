package localfile

import (
	"github.com/caiocfer/todo_cli/todo"
)

func GetTaskItem(taskId int) todo.TodoItem {

	tasks := ReadJson()
	for i := range tasks {
		if taskId == tasks[i].Id {
			return tasks[i]
		}
	}

	return todo.TodoItem{}

}

func DeleteTask(taskId int) ([]todo.TodoItem, bool) {
	tasks := ReadJson()
	for i := range tasks {
		if taskId == tasks[i].Id {
			return append(tasks[:i], tasks[i+1:]...), true
		}
	}

	// If not found, return the original list unchanged and false
	return tasks, false

}

func CompleteTask(taskId int) []todo.TodoItem {
	tasks := ReadJson()
	for i := range tasks {
		if taskId == tasks[i].Id {
			tasks[i].Completed = true
			break
		}
	}

	return tasks

}
