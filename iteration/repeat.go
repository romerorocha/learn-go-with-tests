package iteration

// Similar to strings.Repeat
func Repeat(character string, nTimes int) string {
	var charSequence string
	for i := 0; i < nTimes; i++ {
		charSequence += character
	}
	return charSequence
}
