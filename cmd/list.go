/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"slices"
	"text/tabwriter"

	"github.com/opfocus/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todos",
	Long:  `List will read todo file, and list todos`,
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		panic(err)
	}
	slices.SortFunc(items, todo.SortItems)
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	colNames := func() string {
		tmp := ""
		for _, name := range todo.ColumnName {
			tmp += name + "\t"
		}
		return tmp
	}()
	// print all column names in first line
	fmt.Fprintln(w, colNames)
	defer w.Flush()
	//filte list
	for _, item := range items {
		shouldPrint := false
		dateLayout := "2006-01-02"

		output := item.Lable() + "\t" +
			item.PrettyDone() + "\t" +
			item.PrettyP() + "\t" +
			item.Text + "\t" +
			item.CreateAt.Format(dateLayout) + "\t" +
			item.Tag + "\t" +
			item.Category + "\t"

		if !item.DoneAt.IsZero() {
			output += item.DoneAt.Format(dateLayout) + "\t"
		} else {
			output += "\t"
		}
		switch {
		case allOpt:
			shouldPrint = true
		case priorityOpt != 0:
			if item.Priority == priorityOpt {
				shouldPrint = true
			}
		case doneOpt && item.Done == doneOpt:
			if item.Done == doneOpt {
				shouldPrint = true
			}
		default:
			if item.Done == doneOpt {
				shouldPrint = true
			}
		}

		if shouldPrint {
			fmt.Fprintln(w, output)
		}
	}
}

var (
	doneOpt     bool
	allOpt      bool
	priorityOpt int
)

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	listCmd.Flags().BoolVar(&doneOpt, "done", false, "show 'Done' Todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "show all Todos")
	listCmd.Flags().IntVarP(&priorityOpt, "priority", "p", 0, "show special priority todos")
}
