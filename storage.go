package main

import (
	"encoding/json"
	"os"
)

const TASK_DATA_FILE = "taskdata.json"

func saveTaskData() {
    file, _ := json.MarshalIndent(taskData, "", " ")
    _ = os.WriteFile(TASK_DATA_FILE, file, 0644)
}

func loadTaskData() {
    file, err := os.ReadFile(TASK_DATA_FILE)
    if err != nil {
        // If the file doesn't exist, initialize with empty data
        taskData = TaskData{Tasks: []Task{}}
        return
    }
    _ = json.Unmarshal(file, &taskData)
}
