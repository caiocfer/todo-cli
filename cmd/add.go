package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/caiocfer/todo_cli/localfile"
	"github.com/caiocfer/todo_cli/todo"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add item to todo list",
	Run:   AddItem,
}

func AddItem(cmd *cobra.Command, args []string) {

	var todoName string

	in := bufio.NewReader(os.Stdin)
	fmt.Println("Add a task to your todo item")
	fmt.Print("Task Name: ")
	todoName, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	todoName = strings.TrimSpace(todoName)

	if todoName == "" {
		fmt.Println("Task name can't be avaliable be empty")
		return
	}

	taskItem := createTodoItem(todoName)
	localfile.WriteToJson(taskItem)

}

func createTodoItem(taskName string) todo.TodoItem {
	task := todo.TodoItem{
		Id:        0,
		Name:      taskName,
		Completed: false,
	}

	return task

}
