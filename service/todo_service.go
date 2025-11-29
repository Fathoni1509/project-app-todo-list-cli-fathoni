package service

import (
	"errors"
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

	// validasi nama task, priority
	// if strings.TrimSpace(taskName) == "" {
	// 	return errors.New("judul tidak boleh kosong")
	// }

	// if strings.TrimSpace(priority) == "" {
	// 	return errors.New("priority tidak boleh kosong")
	// }

	todos, err := utils.ReadTasksFromFile()
	if err != nil {
		return err
	}

	// cek duplikat
	for _, t := range todos {
		if strings.EqualFold(t.Task, taskName) {
			return fmt.Errorf("tugas %s sudah ada", taskName)
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

	fmt.Println("List Task")
	if len(todos) == 0 {
		fmt.Println("tidak ada tugas")
		return nil
	}

	for _, t := range todos {
		fmt.Printf("ID: %d, Task: %s, Status: %s, Priority: %s\n", t.ID, t.Task, t.Status, t.Priority)
	}
	// argumen command blum dicek
	return nil
}

// method search task
func (todoservice *TodoService) SearchTask(search string) error {
	if strings.TrimSpace(search) == "" {
		return errors.New("keyword tidak boleh kosong")
	}

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

	fmt.Println("Hasil Pencarian")
	if len(found) == 0 {
		fmt.Println("tidak ada tugas yang sesuai")
		return nil
	}

	for _, t := range found {
		fmt.Printf("ID: %d, Task: %s, Status: %s, Priority: %s\n", t.ID, t.Task, t.Status, t.Priority)
	}

	return nil
}

// method update
func (todoservice *TodoService) UpdateTask(id int, taskName, status, priority string) error {
	todos, err := utils.ReadTasksFromFile()

	if err != nil {
		return err
	}

	found := false
	var success bool

	for i, t := range todos {
		if id == t.ID {
			found = true
			success = true

			if taskName != "" {
				todos[i].Task = taskName
			}

			if status != "" {
				status = strings.ToLower(status)
				switch status{
				case "pending", "progress", "completed":
					todos[i].Status = status
				default:
					fmt.Println("status not allowed, enter status: pending, progress, or completed")
					success = false
				}
			}

			if priority != "" {
				priority = strings.ToLower(priority)
				switch priority{
				case "low", "medium", "high":
					todos[i].Priority = priority
				default:
					fmt.Println("priority not allowed, enter priority: low, medium, or high")
					success = false
				}
			}

			break
		}
	}

	if !found {
		fmt.Println("ID task not found")
		return nil
	}
	
	if !success {
		fmt.Printf("Update id task:%d failed", id)
		return nil
	}

	fmt.Printf("Update id task:%d success", id)
	return utils.WriteTasksToFile(todos)
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
		fmt.Printf("Delete failed: ID task %v not found\n", id)
		return nil
	}

	fmt.Printf("Delete task with ID %v success\n", id)
	return utils.WriteTasksToFile(newTodos)
}