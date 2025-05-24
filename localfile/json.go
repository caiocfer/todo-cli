package localfile

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/caiocfer/todo_cli/constants"
	"github.com/caiocfer/todo_cli/todo"
)

func WriteToJson(todoItem todo.TodoItem) {

	readItems := ReadJson()
	if len(readItems) > 0 {
		lastId := readItems[len(readItems)-1].Id
		newId := lastId + 1

		todoItem.Id = newId
	}

	readItems = append(readItems, todoItem)

	data, err := json.MarshalIndent(readItems, "", "")
	if err != nil {
		fmt.Println("Failed to marshal file")
		return
	}

	err = os.WriteFile(constants.FILE_NAME, data, 0644)
	if err != nil {
		fmt.Printf("Failed to write file: %v\n", err)
		return
	}

	fmt.Println("Task added successfully")
}

func ReadJson() []todo.TodoItem {
	createFileIfNotExists()

	var readItems []todo.TodoItem
	file, err := os.ReadFile(constants.FILE_NAME)
	if err != nil {
		fmt.Errorf("Failed to read file", err)
	}

	err = json.Unmarshal(file, &readItems)
	if err != nil {
		fmt.Errorf("Failed to parse data", err)
	}

	return readItems
}

func UpdateJson(tasks []todo.TodoItem) error {
	newTaskList, err := json.MarshalIndent(tasks, "", "")
	if err != nil {
		fmt.Println("Failed to marshal file")
		return err
	}

	err = os.WriteFile(constants.FILE_NAME, newTaskList, 0644)
	if err != nil {
		fmt.Printf("Failed to write file: %v\n", err)
		return err
	}

	return nil

}

func createFileIfNotExists() {
	file, err := os.OpenFile(constants.FILE_NAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

}
