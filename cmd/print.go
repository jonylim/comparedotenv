/*
Copyright © 2022 Jony Lim

*/
package cmd

import (
	"github.com/jonylim/comparedotenv/cmd/run"

	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:     "print FILE",
	Short:   "Print content of an .env file",
	Long:    `Print content of an .env file.`,
	Example: "  comparedotenv print .env",
	Args:    cobra.ExactArgs(1),
	Run:     run.Print,
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
