/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/opfocus/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit the task",
	Long:  `user can edit history tasks, when they need`,
	Run:   editRun,
}

func editRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		panic(err)
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not vaild label", err)
	}

	if i > 0 && i <= len(items) {
		if eidtPriority != 0 {
			items[i-1].SetPriority(eidtPriority)
		}
		if editText != "" {
			items[i-1].Text = editText
		}
		if editTag != "" {
			items[i-1].Tag = editTag
		}
		if editCategory != "" {
			items[i-1].Category = editCategory
		}

		todo.SaveItems(viper.GetString("datafile"), items)
		fmt.Printf("%v %v\n", items[i-1], "edit done")
	} else {
		log.Println(i, " doesn't match any items")
	}
}

var (
	eidtPriority int
	editText     string
	editTag      string
	editCategory string
)

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	editCmd.Flags().IntVarP(&eidtPriority, "edit_priority", "p", 0, "edit the task priority, H=1, M=2, L=3")
	editCmd.Flags().StringVarP(&editText, "edit_task", "t", "", "edit the task content")
	editCmd.Flags().StringVar(&editTag, "edit_tag", "", "edit the task tag")
	editCmd.Flags().StringVarP(&editCategory, "edit_category", "c", "", "edit the task category")

}
