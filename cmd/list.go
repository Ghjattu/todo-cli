/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

var doneOpt bool
var allOpt bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all todo",
	Long:  `The command 'list' will display all todo items in order`,
	Run: func(cmd *cobra.Command, args []string) {
		items, err := todo.ReadItems(dataPath)
		if err != nil {
			log.Printf("read items error: %v\n", err)
		}
		sort.Sort(todo.ByPri(items))
		w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
		for i, item := range items {
			if allOpt || item.Done == doneOpt {
				fmt.Fprintln(w, item.Label(i+1)+"\t"+item.PrettyDone()+"\t"+item.PrettyP()+"\t"+item.Timestamp+"\t"+item.Text+"\t")
			}
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.
	listCmd.Flags().BoolVarP(&doneOpt, "done", "d", false, "show 'done' todos.")
	listCmd.Flags().BoolVarP(&allOpt, "all", "a", false, "show all todos.")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
