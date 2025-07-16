package base

import (
	"log/slog"
	"strconv"
	"strings"
)

// VersionCompare
// 比较版本号  v1 和 v2,
// v1 > v2 返回 1,
// v1 = v2 返回 0,
// v1 < v2 返回 -1
func VersionCompare(v1, v2 string) int {
	version1 := strings.Split(v1, ".")
	version2 := strings.Split(v2, ".")

	maxLen := max(len(version1), len(version2))

	if len(version1) < maxLen {
		for _, _ = range version2[len(version1):] {
			version1 = append(version1, "0")
		}
	} else if len(version2) < maxLen {
		for _, _ = range version1[len(version2):] {
			version2 = append(version2, "0")
		}
	}

	for index, kv := range version1 {
		v1k, err := strconv.Atoi(kv)
		if err != nil {
			slog.Error("version compare error", "data", version1, "index", index, "error", err.Error())
		}
		v2k, err := strconv.Atoi(version2[index])
		if err != nil {
			slog.Error("version compare error", "data", version2, "index", index, "error", err.Error())
		}
		if v1k > v2k {
			return 1
		} else if v1k < v2k {
			return -1
		}
	}

	return 0

}
