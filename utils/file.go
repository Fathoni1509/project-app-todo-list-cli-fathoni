package utils

import (
	"encoding/json"
	"errors"
	"os"
	"project-app-todo-list-cli-fathoni/model"
)

// path file json untuk simpan todo
const TaskFilePath = "data/todos.json"

// cek direktori dan file
func CheckFile() error {
	_, err := os.Stat(TaskFilePath)
	if errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir("data", 0755); err != nil {
			return err
		}
		return os.WriteFile(TaskFilePath, []byte("[]"), 0644)
	}
	return nil
}

// read file
func ReadTasksFromFile() ([]model.Todo, error) {
	if err := CheckFile(); err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(TaskFilePath)
	if err != nil {
		return nil, err
	}

	var todos []model.Todo
	if err := json.Unmarshal(bytes, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

// write file
func WriteTasksToFile(todos []model.Todo) error {
	bytes, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return  err
	}
	return os.WriteFile(TaskFilePath, bytes, 0644)
}