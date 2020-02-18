package iteration

import "strings"

// Similar to strings.Repeat
func Repeat(character string, nTimes int) string {
	var sb strings.Builder
	for i := 0; i < nTimes; i++ {
		sb.WriteString(character)
	}
	return sb.String()
}
