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

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one task from tasks",
	Long:  `This command uses to remove one task with special No. `,
	Run:   rmRun,
}

func rmRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		panic(err)
	}
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not valid lable", err)
	}
	if i > 0 && i <= len(items) {
		items = append(items[:i-1], items[i:]...)
		fmt.Printf("remove done")
		todo.SaveItems(viper.GetString("datafile"), items)
	} else {
		log.Println(i, " doesn't match any items")
	}

}
func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
