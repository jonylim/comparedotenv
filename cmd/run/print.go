package run

import (
	"fmt"

	"github.com/jonylim/comparedotenv/internal/helper"

	"github.com/spf13/cobra"
)

var PrintFlags = struct {
	FilterKey string
}{}

func Print(cmd *cobra.Command, args []string) {
	envMap := readEnvFileOrExit(args[0])
	keys, maxLen := getKeysAndMaxKeyLen(envMap, PrintFlags.FilterKey)
	for _, k := range keys {
		fmt.Println(helper.PadRight(k, " ", maxLen), "=", envMap[k])
	}
}
