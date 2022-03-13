/*
Copyright © 2022 Jony Lim

*/
package cmd

import (
	"github.com/jonylim/comparedotenv/cmd/run"
	"github.com/spf13/cobra"
)

// compareCmd represents the compare command
var compareCmd = &cobra.Command{
	Use:   "compare FILE-1 FILE-2 [FILE-N]...",
	Short: "Compare values from 2 or more env files",
	Long:  `Compare values from 2 or more env files.`,
	Example: `  comparedotenv compare .env.example .env
  comparedotenv compare .env.example .env -f DATABASE`,
	Args: cobra.MinimumNArgs(2),
	Run:  run.Compare,
}

func init() {
	rootCmd.AddCommand(compareCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	compareCmd.PersistentFlags().StringVarP(&run.CompareFlags.FilterKey, "filter-key", "f", "", "filter key")
}
