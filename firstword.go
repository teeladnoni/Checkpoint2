package main

func FirstWord(s string) string {
	if s == "" {
		return ""
	}
	start := false
	index := 0
	for i, b := range s {
		if b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' {
			start = true
		}

		if start {
			if b == ' ' {
				index += i
				break
			}
		}
	}
	return s[:index] + "\n"
}

func main() {
	print(FirstWord("Hello Who"))
}
