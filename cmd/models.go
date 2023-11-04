package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Duration  time.Time `json:"deadline"`
	Completed bool      `json:"completed"`
}

func generateGoogleUUID() string {
	uuid := uuid.New()
	return uuid.String()
}

// Read tasks from a file
func loadTasksFromFile() ([]Task, error) {

	var tasks []Task

	data, err := os.ReadFile("tasks.json")

	if err != nil && !os.IsNotExist(err) {
		// Handle error, unless the file doesn't exist yet
		fmt.Println("Error reading file:", err)
		return nil, nil
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		fmt.Println("Error unmarshaling JSON data:", err)
		return nil, err
	}

	return tasks, nil

}

func readAllTasks() error {
	tasks, err := loadTasksFromFile()

	if err != nil {
		fmt.Println("Error loading file from db:", err)

		return err
	}

	for _, task := range tasks {
		fmt.Printf("ID: %s \nTitle: %s \nDeadline: %s\n\n", task.Id, task.Title, task.Duration)
	}

	return nil

}

func readSingleTaskHelper(id string) (*Task, error) {
	tasks, err := loadTasksFromFile()

	if err != nil {
		fmt.Println("Error loading file from db:", err)
		return nil, err
	}

	// small app, don't really care about optimization
	for _, task := range tasks {
		if task.Id == id {
			return &task, nil
		}
	}
	return nil, nil
}

func readSingleTask(id string) error {
	task, err := readSingleTaskHelper(id)

	if err != nil {
		fmt.Println("Error loading file from db:", err)
		return err
	}

	if task == nil {
		fmt.Printf("Could not find task with ID: %s\n", id)
		return nil
	}

	fmt.Printf("ID: %s \nTitle: %s \nDeadline: %s\n\n", task.Id, task.Title, task.Duration)

	return nil

}

func markTaskAsCompleted(id string) error {

	tasks, err := loadTasksFromFile()
	if err != nil {
		fmt.Println("Error loading file from db:", err)
		return err
	}
	currentTask, err := readSingleTaskHelper(id)

	if currentTask == nil {
		fmt.Printf("Could not find task with ID: %s\n", id)
		return nil
	}

	if err != nil {
		fmt.Println("Error finding task:", err)
		return err
	}

	result := removeTaskFromArray(tasks, id)

	currentTask.Completed = true
	result = append(result, *currentTask)

	err = writeTasksToFile(result)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Printf("Task with ID:%s has been marked as completed.", id)

	return nil

}


func editTask(id string) error {

	tasks, err := loadTasksFromFile()
	if err != nil {
		fmt.Println("Error loading file from db:", err)
		return err
	}
	currentTask, err := readSingleTaskHelper(id)

	if currentTask == nil {
		fmt.Printf("Could not find task with ID: %s\n", id)
		return nil
	}

	if err != nil {
		fmt.Println("Error finding task:", err)
		return err
	}

	result := removeTaskFromArray(tasks, id)

	currentTask.Completed = true
	result = append(result, *currentTask)

	err = writeTasksToFile(result)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Printf("Task with ID:%s has been marked as completed.", id)

	return nil

}

func deleteTask(id string) error {
	tasks, err := loadTasksFromFile()

	if err != nil {
		fmt.Println("Error loading file from db:", err)

		return err
	}

	result := removeTaskFromArray(tasks, id)

	err = writeTasksToFile(result)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Printf("Task with ID:%s has been deleted from tasks.json successfully.", id)

	return nil

}

// Save tasks to a file
func saveTasksToFile(title, durationStr string) error {

	duration, err := parseDuration(durationStr)

	if err != nil {
		fmt.Println("Error parsing duration: ", err)
		return err
	}

	task := Task{
		Id:        generateGoogleUUID(),
		Title:     title,
		Duration:  duration,
		Completed: false,
	}

	file, err := os.OpenFile("tasks.json", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}

	file.Close()

	existingData, err := os.ReadFile("tasks.json")
	if err != nil && !os.IsNotExist(err) {
		// Handle error, unless the file doesn't exist yet
		fmt.Println("Error reading file:", err)
		return err
	}

	var tasks []Task
	if len(existingData) > 0 {
		if err := json.Unmarshal(existingData, &tasks); err != nil {
			fmt.Println("Error unmarshaling existing JSON data:", err)
			return err
		}
	}

	tasks = append(tasks, task)

	err = writeTasksToFile(tasks)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	fmt.Println("Task added to tasks.json successfully.")

	return nil
}

func parseDuration(input string) (time.Time, error) {
	duration, err := time.ParseDuration(input)

	if err != nil {
		fmt.Println("Error parsing duration: ", err)
		return time.Time{}, err
	}
	referenceTime := time.Now()
	resultTime := referenceTime.Add(duration)

	return resultTime, nil
}

func removeTaskFromArray(tasks []Task, id string) []Task {
	var result []Task

	for _, task := range tasks {
		if task.Id != id {
			result = append(result, task)
		}
	}

	return result

}

func writeTasksToFile(tasks []Task) error {
	updatedData, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}

	// Write the updated JSON data back to the file
	err = os.WriteFile("tasks.json", updatedData, 0644)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil
}
