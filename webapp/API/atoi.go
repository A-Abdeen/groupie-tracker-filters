package gt

func Atoi(s string) int {
	runes := []rune(s)
	var n int = 0
	length := len(runes)
	for i := 0; i < length; i++ {
		if s[i] > '9' || s[i] < '0' {
			continue
		}
		n = (n*10 + int(s[i]) - '0')

	}
	return n
}
