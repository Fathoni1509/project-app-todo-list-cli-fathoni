package utils

import (
	"fmt"
	"os"
	"project-app-todo-list-cli-fathoni/model"
	"text/tabwriter"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func colorStatus(status string) string {
	switch status {
	case "new":
		return Green + status + Reset
	case "pending":
		return Blue + status + Reset
	case "progress":
		return Yellow + status + Reset
	case "completed":
		return Blue + status + Reset
	default:
		return status
	}
}

func PrintTabel(todos []model.Todo) {
	if len(todos) == 0 {
		fmt.Printf("%s!!! No Tasks !!!!%s\n", Red, Reset)
		return
	}

	table := tabwriter.NewWriter(os.Stdout, 21, 0, 1, ' ', 0)

	fmt.Println("===================================== List Tasks ====================================")
	fmt.Fprintln(table, "+--------------------+--------------------+--------------------+--------------------+")
	fmt.Fprintf(table, "| ID\t| Task\t| Status\t| Priority\t|\n")
	fmt.Fprintln(table, "+--------------------+--------------------+--------------------+--------------------+")

	for _, t := range todos {
		colorStatus := colorStatus(t.Status)
		fmt.Fprintf(table, "| %d \t| %s\t| %-27s\t| %s\t|\n", t.ID, t.Task, colorStatus, t.Priority)
		fmt.Fprintln(table, "+--------------------+--------------------+--------------------+--------------------+")
	}

	table.Flush()
}
