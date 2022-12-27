package main

import "strings"

func ToUpper2(str string) string {
	var builder strings.Builder // strings 패키지의 Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + c - 'a')
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}
