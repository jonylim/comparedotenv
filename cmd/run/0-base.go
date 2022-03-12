package run

import (
	"log"
	"sort"

	"github.com/joho/godotenv"
)

func readEnvFileOrExit(filename string) map[string]string {
	envMap, err := godotenv.Read(filename)
	if err != nil {
		log.Fatal("Reading '", filename, "' failed: ", err)
	}
	return envMap
}

func getKeysAndMaxKeyLen(envMap map[string]string) (keys []string, maxLen int) {
	keys = make([]string, 0, len(envMap))
	for k := range envMap {
		keys = append(keys, k)
		if len(k) > maxLen {
			maxLen = len(k)
		}
	}
	sort.Strings(keys)
	return
}
