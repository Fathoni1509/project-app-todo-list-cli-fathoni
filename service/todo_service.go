package service

import (
	"fmt"
	"project-app-todo-list-cli-fathoni/model"
	"project-app-todo-list-cli-fathoni/utils"
	"strings"

)

type TodoService struct{}

func NewTodoService() TodoService {
	return TodoService{}
}

// method add task
func (todoservice *TodoService) AddTodo(taskName string, priority string) error {

	todos, err := utils.ReadTasksFromFile()
	if err != nil {
		return err
	}

	// cek duplikat
	for _, t := range todos {
		if strings.EqualFold(t.Task, taskName) {
			return fmt.Errorf("%serror: the task title %s already exists%s", utils.Red, taskName, utils.Reset)
		}
	}

	// generate ID
	newID := 1
	if len(todos) > 0 {
		newID = todos[len(todos)-1].ID + 1
	}

	// susun format
	newTodo := model.Todo {
		ID: newID,
		Task: strings.ToLower(taskName),
		Status: "new",
		Priority: priority,
	}

	todos = append(todos, newTodo)
	return utils.WriteTasksToFile(todos)
}

// method list
func (todoservice *TodoService) ListTask() error {
	todos, err := utils.ReadTasksFromFile()
	
	if err != nil {
		return err
	}

	utils.PrintTabel(todos)
	// argumen command blum dicek
	return nil
}

// method search task
func (todoservice *TodoService) SearchTask(search string) error {

	todos, err := utils.ReadTasksFromFile()
	
	if err != nil {
		return err
	}

	var found []model.Todo
	for _, t := range todos {
		if strings.Contains(strings.ToLower(t.Task), strings.ToLower(search)){
			found = append(found, t)
		}
	}

	utils.PrintTabel(found)

	return nil
}

// method update
func (todoservice *TodoService) UpdateTask(id int, taskName, status, priority string) error {
	todos, err := utils.ReadTasksFromFile()

	if err != nil {
		return err
	}

	foundIndex := -1
	for i, t := range todos {
		if t.ID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return  fmt.Errorf("%serror: task id %d not found%s", utils.Red, id, utils.Reset)
	}

	if taskName != "" {
		todos[foundIndex].Task = taskName
	}

	if status != "" {
		status = strings.ToLower(status)
		switch status{
		case "new", "pending", "progress", "completed":
			todos[foundIndex].Status = status
		default:
			return fmt.Errorf("%s error: status not allowed, enter status: new, pending, progress, or completed%s", utils.Red, utils.Reset)
		}
	}

	if priority != "" {
		priority = strings.ToLower(priority)
		switch priority{
		case "low", "medium", "high":
			todos[foundIndex].Priority = priority
		default:
			return fmt.Errorf("%serror: priority not allowed, enter priority: low, medium, or high%s", utils.Red, utils.Reset)
		}
	}

	err = utils.WriteTasksToFile(todos)
	if err != nil {
		return  err
	}

	fmt.Printf("%sTask updated successfully.%s", utils.Green, utils.Reset)
	return nil
}

// method delete
func (todoservice *TodoService) DeleteTask(id int) error {
	todos, err := utils.ReadTasksFromFile()

	if err != nil {
		return err
	}

	var newTodos []model.Todo

	found := false

	for _, t := range todos {
		if id == t.ID {
			found = true
			continue
		}
		newTodos = append(newTodos, t)
	}

	if !found {
		return fmt.Errorf("%serror: delete failed, task %d not found%s", utils.Red, id, utils.Reset)
	}

	err = utils.WriteTasksToFile(newTodos)
	if err != nil {
		return err
	}
	fmt.Printf("%sTask delete successfully%s\n", utils.Green, utils.Reset)
	return nil
}