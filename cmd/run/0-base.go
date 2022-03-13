package run

import (
	"log"
	"sort"
	"strings"

	"github.com/joho/godotenv"
)

func readEnvFileOrExit(filename string) map[string]string {
	envMap, err := godotenv.Read(filename)
	if err != nil {
		log.Fatal("Reading '", filename, "' failed: ", err)
	}
	return envMap
}

func getKeysAndMaxKeyLen(envMap map[string]string, filterKey string) (keys []string, maxLen int) {
	keys = make([]string, 0, len(envMap))
	if filterKey == "" {
		for k := range envMap {
			keys = append(keys, k)
			if len(k) > maxLen {
				maxLen = len(k)
			}
		}
	} else {
		filterKey = strings.ToUpper(filterKey)
		for k := range envMap {
			if strings.Contains(strings.ToUpper(k), filterKey) {
				keys = append(keys, k)
				if len(k) > maxLen {
					maxLen = len(k)
				}
			}
		}
	}
	sort.Strings(keys)
	return
}

func mergeKeys(list ...[]string) []string {
	allKeys := make([]string, 0)
	keyExists := make(map[string]bool)
	for _, keys := range list {
		for i := range keys {
			if !keyExists[keys[i]] {
				allKeys = append(allKeys, keys[i])
				keyExists[keys[i]] = true
			}
		}
	}
	sort.Strings(allKeys)
	return allKeys
}
