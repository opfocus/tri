/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/opfocus/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// undoCmd represents the undo command
var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Change completed to incomplete",
	Long:  `if mistake marked done, use this command to change it`,
	Run:   undoRun,
}

func undoRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		panic(err)
	}
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not valid lable", err)
	}
	if i > 0 && i <= len(items) {
		var zeroTime time.Time
		items[i-1].Done, items[i-1].DoneAt = false, zeroTime
		fmt.Printf("%q %v\n", items[i-1].Text, "marked undo")
		todo.SaveItems(viper.GetString("datafile"), items)
	} else {
		log.Println(i, " doesn't match any items")
	}

}

func init() {
	rootCmd.AddCommand(undoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// undoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// undoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
