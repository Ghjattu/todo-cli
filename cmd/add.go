/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"time"
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `The command 'add' will create a new todo item to the list.`,
	Run: func(cmd *cobra.Command, args []string) {
		items, err := todo.ReadItems(dataPath)
		if err != nil {
			log.Printf("read items error: %v\n", err)
		}
		for _, x := range args {
			item := todo.Item{Text: x}
			item.SetPriority(priority)
			item.Done = false
			item.Timestamp = time.Now().Format("2006-01-02")
			items = append(items, item)
		}
		err = todo.SaveItems(dataPath, items)
		if err != nil {
			log.Printf("save items error: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1, 2, 3")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
