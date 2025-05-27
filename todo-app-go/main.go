package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID          int
	Description string
	isCompleted bool
}

var tasks []Task
var nextID = 1

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("\nCommands:")
		fmt.Println("  add <description>")
		fmt.Println("  list")
		fmt.Println("  complete <id>")
		// fmt.Println("  delete <id>")
		fmt.Println("  exit")
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.SplitN(input, " ", 2)
		command := parts[0]

		switch command {
		case "add":
			if len(parts) < 2 || parts[1] == "" {
				fmt.Println("Error: Please provide a task description.")
				continue
			}
			description := parts[1]
			addTask(description)

		case "list":
			listTasks()

		case "complete":
			if len(parts) < 2 {
				fmt.Println("Error: Please provide a task description.")
				continue
			}

			idStr := parts[1]
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Error: Invalid task ID. Please enter a number.")
				continue
			}
			completeTask(id)

		case "delete":
			if len(parts) < 2 {
				fmt.Println("Error: Please provide a task description.")
				continue
			}

			idStr := parts[1]
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Error: Invalid task ID. Please enter a number.")
				continue
			}
			deleteTask(id)

		case "exit":
			fmt.Println("Exiting task manager")

		default:
			fmt.Println("Unknown command, Please try again")
		}
	}

}

func addTask(description string) {
	task := Task{
		ID:          nextID,
		Description: description,
		isCompleted: false,
	}

	tasks = append(tasks, task)
	nextID++
	// fmt.Println("Task added")
	fmt.Printf("Task added: \"%s\" (ID: %d)\n", description, task.ID)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks in the list")
		return
	}
	fmt.Println("Your tasks")
	for _, task := range tasks {
		status := " "
		if task.isCompleted {
			status = "x"
		}
		fmt.Printf("[%s] ID: %d - %s\n", status, task.ID, task.Description)
	}
}

func completeTask(id int) {
	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].isCompleted = true
			fmt.Printf("Task %d marked as complete: \"%s\"\n", task.ID, task.Description)
			found = true
			break
		}

	}
	if !found {
		fmt.Printf("Task with ID %d not found.\n", id)

	}
}

func deleteTask(id int) {
	foundIndex := -1
	for i, task := range tasks {
		if task.ID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		fmt.Printf("Task with ID %d not found.\n", id)
		return
	}

	tasks = append(tasks[:foundIndex], tasks[foundIndex+1:]...)
	fmt.Printf("Task ID %d deleted.\n", id)

}
