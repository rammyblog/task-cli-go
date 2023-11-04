package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Duration time.Time `json:"deadline"`
}

func generateGoogleUUID() string {
	uuid := uuid.New()
	return uuid.String()
}

// Read tasks from a file
func loadTasksFromFile() error {

	var tasks []Task

	data, err := os.ReadFile("tasks.json")

	if err != nil && !os.IsNotExist(err) {
		// Handle error, unless the file doesn't exist yet
		fmt.Println("Error reading file:", err)
		return nil
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		fmt.Println("Error unmarshaling JSON data:", err)
		return err
	}

	for _, task := range tasks {
		fmt.Printf("ID: %s \nTitle: %s \nDeadline: %s\n\n", task.Id, task.Title, task.Duration)
	}

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
		Id:       generateGoogleUUID(),
		Title:    title,
		Duration: duration,
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
