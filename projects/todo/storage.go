package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const taskFile = "tasks.json"

// SaveTasks writes the current tasks slice to a JSON file
func SaveTasks() error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(taskFile, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// LoadTasks reads tasks from the JSON file into the tasks slice
func LoadTasks() error {
	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		// No file, start fresh
		tasks = []Task{}
		nextID = 1
		return nil
	}

	data, err := ioutil.ReadFile(taskFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return err
	}

	// Set nextID to max existing ID + 1
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	nextID = maxID + 1

	return nil
}
