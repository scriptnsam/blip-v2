package helper

import (
	"slices"
	"strings"
)

func ShouldExcludeBrewFormula(name string) bool {
	excludePrefixes := []string{
		"lib", "gcc", "gmp", "mpfr", "pkg-config", "python", "perl", "node",
		"rust", "openssl", "autoconf", "automake", "cmake", "make", "gettext", "coreutils",
	}

	for _, prefix := range excludePrefixes {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}
	return false
}

func ShouldExcludeCLITool(name string) bool {
	exclude := []string{
		"awk", "sed", "nvim", "jq", "curl", "wget", "g++", "clang", "python", "node", "npm", "go",
		"gcc", "cmake", "make", "autoconf", "automake", "ld", "lsof", "strace", "git",
	}
	return slices.Contains(exclude, name)
}

func ShouldExcludeMacPkg(id string) bool {
	return strings.HasPrefix(id, "com.apple.") ||
		strings.HasPrefix(id, "com.google.pkg.") ||
		strings.HasPrefix(id, "com.microsoft.pkg.") ||
		strings.HasPrefix(id, "org.postgresql.") ||
		strings.HasPrefix(id, "com.oracle.") ||
		strings.HasPrefix(id, "org.pqrs.") ||
		strings.Contains(id, "driver") ||
		strings.Contains(id, "support") ||
		strings.Contains(id, "framework")
}
