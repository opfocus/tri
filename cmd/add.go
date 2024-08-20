/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"

	"github.com/opfocus/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list`,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		panic(err)
	}
	for _, v := range args {
		item := todo.Item{Text: v, CreateAt: time.Now(), Tag: tag, Category: category}
		item.SetPriority(priority)
		items = append(items, item)
	}
	err = todo.SaveItems(viper.GetString("datafile"), items)
	if err != nil {
		panic(err)
	}
}

var (
	priority int
	tag      string
	category string
)

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: High=1, Middle=2, Low=3")
	addCmd.Flags().StringVarP(&tag, "tag", "t", "none", "Tag: task in which tag")
	addCmd.Flags().StringVarP(&category, "category", "c", "None", "task in which category")
}
