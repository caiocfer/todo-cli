package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/caiocfer/todo_cli/localfile"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task in your task list",
	Run:   deleteItem,
}

func deleteItem(cmd *cobra.Command, args []string) {
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Select the task you wish to delete")
	localfile.ListItems()

	fmt.Print("Task Id: ")
	taskId, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	taskId = strings.TrimSpace(taskId)
	taskIdInt, err := strconv.Atoi(taskId)
	if err != nil {
		fmt.Println("Invalid Id:", err)
		return
	}

	updatedTaskList := localfile.DeleteTask(taskIdInt)

	err = localfile.UpdateJson(updatedTaskList)
	if err != nil {
		fmt.Printf("Failed to save to json %s", err)
	} else {
		fmt.Println("Task deleted successfully")
		localfile.ListItems()
	}

}
