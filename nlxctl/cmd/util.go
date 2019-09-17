package cmd

import "strings"

func splitConfigString(configString string) []string {
	return strings.Split(configString, "\n---\n")
}
