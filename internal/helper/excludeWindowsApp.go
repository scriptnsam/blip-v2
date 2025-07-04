package helper

import "strings"

func ShouldExcludeWindowsApp(name string) bool {
	name = strings.ToLower(name)
	excludedPatterns := []string{
		"microsoft visual c++",
		"microsoft .net",
		"redistributable",
		"security update",
		"hotfix",
		"nvidia driver",
		"intel",
		"update for",
		"windows sdk",
		"kb", // KB updates
		"driver",
		"support assistant",
		"framework",
	}

	for _, pattern := range excludedPatterns {
		if strings.Contains(name, pattern) {
			return true
		}
	}

	return false
}
