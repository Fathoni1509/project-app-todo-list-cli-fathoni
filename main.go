package main

import (
	// "fmt"
	"project-app-todo-list-cli-fathoni/cmd"
	// "strconv"

	// "github.com/spf13/cobra"
)

// var rootCmd = &cobra.Command{
// 	Use: "todo",
// 	Short: "aplikasi Todo list sederhana dengan golang CLI",
// }

// var addcmd = &cobra.Command{
// 	Use: "add [id_product] [product]",
// 	Short: "command add price",
// 	Run: func (cmd *cobra.Command, args []string){
// 		id, _ := strconv.Atoi(args[0])
// 		product := args[1]
// 		fmt.Println("id:", id)
// 		fmt.Println("product:", product)
// 	},
// }

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

func main() {
	// registrasikan
	// updatecmd.Flags().StringP("name", "n", "", "name of the item update")
	// rootCmd.AddCommand(addcmd)
	// rootCmd.AddCommand(updatecmd)
	// if err := rootCmd.Execute(); err != nil {
	// 	fmt.Println(err)
	// }


	cmd.Execute()
}