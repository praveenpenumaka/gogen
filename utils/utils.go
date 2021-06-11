package utils

import "strings"

func ToCap(arg0 string, args ...string) string {
	return strings.Replace(arg0, string(arg0[0]), strings.ToUpper(string(arg0[0])), 1)
}
