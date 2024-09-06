package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct{
	description string
	completed bool
}

const taskFile = "tasks.txt"

func main(){
	

	tasks := loadTasksFromFile(taskFile)
	if len(tasks) == 0{
		fmt.Println("No tasks found.")
	} else{
		fmt.Println("Taks loaded successfully")
	}

	for{
		
		fmt.Println("\nWelcome to the Task Manager!")
		fmt.Println("1. Add a Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Mark Task as Complete")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Choose an option: ")
		fmt.Scan(&choice)

		switch choice{
			case 1:
				addTask(&tasks)
			case 2:
				viewTasks(tasks)
			case 3:
				var taskNumber int
				fmt.Print("Enter the task number to delete: ")
				fmt.Scan(&taskNumber)
				deleteTask(&tasks, taskNumber)
			case 4:
				var taskNumber int 
				fmt.Print("Enter the task number to mark as complete: ")
				fmt.Scan(&taskNumber)
				markTaskComplete(&tasks, taskNumber)
			case 5:
				saveTasksToFile(tasks, taskFile)
				fmt.Println("Exiting the Task Manager.")
				return
			default:
				fmt.Println("Invalid choice. Please choose a valid option.")
		}
	}
}

func addTask(tasks *[]Task){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nEnter the task descrition: ")
	taskDescription, _ := reader.ReadString('\n')
	taskDescription = strings.TrimSpace(taskDescription)

	newTask := Task{description: taskDescription, completed: false}
	*tasks = append(*tasks, newTask)
	fmt.Println("Task added successfuly!")
}

func viewTasks(tasks []Task){
	fmt.Println("Your tasks:")
	for i, task := range tasks{
		status := "Incomplete"
		if task.completed{
			status = "Completed"
		}
		fmt.Printf("%d. %s - %s\n", i+1, task.description, status)
	}
}

func deleteTask(tasks *[]Task, taskNumber int){
	if taskNumber > 0 && taskNumber <= len(*tasks){
		*tasks = append((*tasks)[:taskNumber - 1], (*tasks)[taskNumber:]... )
		fmt.Println("Task deleted successfuly!")
	} else {
		fmt.Println("Invalid task number!")
	}
}

func markTaskComplete(tasks *[]Task, taskNumber int){
	if taskNumber > 0 && taskNumber <= len(*tasks){
		(*tasks)[taskNumber - 1].completed = true
		fmt.Println("Task marked as completed!")
	} else{
		fmt.Println("Invalid task number!")
	}
}

func saveTasksToFile(tasks []Task, filename string){
	file, err := os.Create(filename)
	if err != nil{
		fmt.Println("Error creating file:",err)
		return
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	for _, task := range tasks{
		status := "0"
		if task.completed{
			status = "1"
		}
		write.WriteString(fmt.Sprintf("%s | %s\n", task.description, status))
	}
	write.Flush()
}

func loadTasksFromFile(filename string) [] Task{
	var tasks []Task
	file, err := os.Open(filename)
	if err != nil{
		fmt.Println("Error opening file:",err)
		return tasks
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		parts := strings.Split(line, " | ")
		description := parts[0]
		completed := parts[1] == "1"
		tasks = append(tasks, Task{description : description, completed : completed})
	}
	return tasks
}