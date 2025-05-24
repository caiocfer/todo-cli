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

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a task in your task list",
	Run:   completeItem,
}

func completeItem(cmd *cobra.Command, args []string) {
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Select the task you wish to complete")
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

	updatedTaskList := localfile.CompleteTask(taskIdInt)

	err = localfile.UpdateJson(updatedTaskList)
	if err != nil {
		fmt.Printf("Failed to save to json %s", err)
	} else {
		fmt.Println("Task completed successfully")
		localfile.ListItems()
	}

}
