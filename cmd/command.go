package cmd

import (
	"fmt"
	"project-app-todo-list-cli-fathoni/service"
	"project-app-todo-list-cli-fathoni/utils"
	"strings"

	"github.com/spf13/cobra"
)

// inisialisasi service
var svc = service.NewTodoService()

// command untuk menambahkan task
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "command add task",
	Run: func(cmd *cobra.Command, args []string) {
		taskName, _ := cmd.Flags().GetString("task")
		priority, _ := cmd.Flags().GetString("priority")

		if strings.TrimSpace(taskName) == "" {
			fmt.Printf("%sError: task title must not be empty%s\n", utils.Red, utils.Reset)
			return
		}

		if strings.TrimSpace(priority) == "" {
			priority = "low"
		}

		err := svc.AddTodo(taskName, priority)
		if err != nil {
			fmt.Printf("%sError: %s%s", utils.Red, err, utils.Reset)
			return
		}

		fmt.Printf("%sTask add successfully.%s\n", utils.Green, utils.Reset)
	},
}

// command untuk melihat daftar task
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "command to view all task",
	Run: func(cmd *cobra.Command, args []string) {
		err := svc.ListTask()
		if err != nil {
			fmt.Printf("%sError: %v%s\n", utils.Red, err, utils.Reset)
		}
	},
}

// command untuk mencari task
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "command to search title task",
	Run: func(cmd *cobra.Command, args []string) {
		search, _ := cmd.Flags().GetString("task")

		if strings.TrimSpace(search) == "" {
			fmt.Printf("%sError: keyword must not be empty%s", utils.Red, utils.Reset)
			return
		}

		err := svc.SearchTask(search)
		if err != nil {
			fmt.Printf("%sError: %v%s\n", utils.Red, err, utils.Reset)
		}
	},
}

// command untuk update task
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "command to update task",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		taskName, _ := cmd.Flags().GetString("task")
		status, _ := cmd.Flags().GetString("status")
		priority, _ := cmd.Flags().GetString("priority")

		if id == 0 {
			fmt.Printf("%sID must not be 0%s\n", utils.Red, utils.Reset)
			return
		}

		if taskName == "" && status == "" && priority == "" {
			fmt.Printf("%sError: enter at least one change (task/status/priority)%s", utils.Red, utils.Reset)
			return
		}

		err := svc.UpdateTask(id, taskName, status, priority)
		if err != nil {
			fmt.Printf("%sError: %v%s\n", utils.Red, err, utils.Reset)
		}
	},
}

// command untuk hapus task
var deleteCmd = &cobra.Command{
	Use: "delete",
	Short: "command to delete task",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")

		err := svc.DeleteTask(id)
		if err != nil {
			fmt.Printf("%sError: %v%s\n", utils.Red, err, utils.Reset)
		}
	},
}

func init() {
	// registrasikan
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)

	// tambahkan task
	addCmd.Flags().StringP("task", "t", "", "name of the task")
	addCmd.Flags().StringP("priority", "p", "", "priority of the task")

	// cari task
	searchCmd.Flags().StringP("task", "t", "", "search task based on title")

	// update task
	updateCmd.Flags().IntP("id", "i", 0, "id task")
	updateCmd.Flags().StringP("task", "t", "", "name of the task")
	updateCmd.Flags().StringP("status", "s", "", "status of the task")
	updateCmd.Flags().StringP("priority", "p", "", "priority of the task")

	// hapus task
	deleteCmd.Flags().IntP("id", "i", 0, "delete task based on ID task")
}
