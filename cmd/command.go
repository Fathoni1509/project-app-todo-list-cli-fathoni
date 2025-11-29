package cmd

import (
	"fmt"
	"project-app-todo-list-cli-fathoni/service"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// inisialisasi service
var svc = service.NewTodoService()

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "command add task",
	Run: func(cmd *cobra.Command, args []string) {
		// id, _ := strconv.Atoi(args[0])
		taskName, _ := cmd.Flags().GetString("task")

		if strings.TrimSpace(taskName) == "" {
			fmt.Println("Gagal menambahkan: judul task tidak boleh kosong")
			return
		}

		priority, _ := cmd.Flags().GetString("priority")

		if strings.TrimSpace(priority) == "" {
			fmt.Println("Gagal menambahkan: prioritas task tidak boleh kosong")
			return
		}

		// cek penulisan prioritas (belum)
		priority = strings.ToLower(priority)
		if priority != "low" && priority != "medium" && priority != "high" {
			fmt.Println("Gagal menambahkan: status not allowed, enter status: low, medium, or high")
			return
		}

		// err := svc.AddTodo(taskName, priority)
		// if err != nil {
		// 	fmt.Printf("gagal menambahkan task: %v\n", err)
		// 	return
		// }

		svc.AddTodo(taskName, priority)

		fmt.Println("task:", taskName)
		fmt.Println("priority:", priority)
		fmt.Printf("sukses menambahkan task: %v\n", taskName)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "command to view all task",
	Run: func(cmd *cobra.Command, args []string) {
		err := svc.ListTask()
		if err != nil {
			fmt.Printf("Gagal memuat data: %v\n", err)
		}
	},
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "command to search title task",
	Run: func(cmd *cobra.Command, args []string) {
		search, _ := cmd.Flags().GetString("task")
		// if search == "" {
		// 	fmt.Println("keyword tidak boleh kosong")
		// }

		err := svc.SearchTask(search)
		if err != nil {
			fmt.Printf("Gagal menampilkan: %v\n", err)
			return
		}
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "command to update task",
	Run: func(cmd *cobra.Command, args []string) {
		idStr, _ := cmd.Flags().GetString("id")
		taskName, _ := cmd.Flags().GetString("task")
		status, _ := cmd.Flags().GetString("status")
		priority, _ := cmd.Flags().GetString("priority")

		if idStr == "" {
			fmt.Println("ID tidak boleh kosong")
			return
		}

		if taskName == "" && status == "" && priority == "" {
			fmt.Println("Error: masukkan minimal dua flag")
			return
		}

		id, _ := strconv.Atoi(idStr)

		svc.UpdateTask(id, taskName, status, priority)
	},
}

var deleteCmd = &cobra.Command{
	Use: "delete",
	Short: "command to delete task",
	Run: func(cmd *cobra.Command, args []string) {
		idStr, _ := cmd.Flags().GetString("id")

		if idStr == "" {
			fmt.Println("ID tidak boleh kosong")
			return
		}

		id, _ := strconv.Atoi(idStr)

		svc.DeleteTask(id)
	},
}

// var updatecmd = &cobra.Command{
// 	Use: "update [id_product]",
// 	Short: "command update product",
// 	Run: func (cmd *cobra.Command, args []string){
// 		id, _ := strconv.Atoi(args[0])
// 		name, _ := cmd.Flags().GetString("name")
// 		fmt.Println("id:", id)
// 		fmt.Println("name_product:", name)
// 	},
// }

func init() {
	// registrasikan
	// updatecmd.Flags().StringP("name", "n", "", "name of the item update")
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)

	addCmd.Flags().StringP("task", "t", "", "name of the task")
	addCmd.Flags().StringP("priority", "p", "", "priority of the task")

	searchCmd.Flags().StringP("task", "t", "", "search task based on title")

	updateCmd.Flags().StringP("id", "i", "", "id task")
	updateCmd.Flags().StringP("task", "t", "", "name of the task")
	updateCmd.Flags().StringP("status", "s", "", "status of the task")
	updateCmd.Flags().StringP("priority", "p", "", "priority of the task")

	deleteCmd.Flags().StringP("id", "i", "", "delete task based on ID task")

	// listCmd.Flags().StringP("list", "l", "", "list all task")
	// // rootCmd.AddCommand(updatecmd)
	// if err := rootCmd.Execute(); err != nil {
	// 	fmt.Println(err)
	// }
}
