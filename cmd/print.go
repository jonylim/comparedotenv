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
	Use:   "print FILE",
	Short: "Print content of an .env file",
	Long:  `Print content of an .env file.`,
	Example: `  comparedotenv print .env
  comparedotenv print .env -f DATABASE`,
	Args: cobra.ExactArgs(1),
	Run:  run.Print,
}

func init() {
	rootCmd.AddCommand(printCmd)

	printCmd.PersistentFlags().StringVarP(&run.PrintFlags.FilterKey, "filter-key", "f", "", "filter key")
}
