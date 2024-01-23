package iteration

import "strings"

func Repeat(repeatCount int, char string) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += char
	}

	return repeated
}

func RepeatStringBuilder(repeatCount int, char string) string {
	var sb strings.Builder

	for i := 0; i < repeatCount; i++ {
		sb.WriteString(char)
	}

	return sb.String()
}
