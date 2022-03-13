package run

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/jonylim/comparedotenv/internal/helper"
	"github.com/spf13/cobra"
)

var CompareFlags = struct {
	FilterKey string
}{}

func Compare(cmd *cobra.Command, args []string) {
	envMaps := make([]map[string]string, len(args))
	envKeys := make([][]string, len(args))
	var maxFilenameLen, maxKeyLen int
	var hasRelativeFilename bool

	for i := range args {
		m := readEnvFileOrExit(args[i])
		keys, maxLen := getKeysAndMaxKeyLen(m, CompareFlags.FilterKey)

		envMaps[i] = m
		envKeys[i] = keys
		if len(args[i]) > maxFilenameLen {
			maxFilenameLen = len(args[i])
		}
		if maxLen > maxKeyLen {
			maxKeyLen = maxLen
		}
		if strings.HasPrefix(args[i], "..") {
			hasRelativeFilename = true
		}
	}

	for i, filename := range args {
		if hasRelativeFilename {
			if absFilepath, err := filepath.Abs(filename); absFilepath != "" && err == nil {
				filename = absFilepath
			}
		}
		fmt.Printf("[file-%d] %s\n", i+1, filename)
	}
	fmt.Println()

	allKeys := mergeKeys(envKeys...)
	emptyKey := helper.PadRight("        ", " ", maxKeyLen)
	for _, key := range allKeys {
		fmt.Println(key)
		for i := range args {
			fmt.Printf("%s [file-%d] %s\n", emptyKey, i+1, envMaps[i][key])
		}
	}
}
