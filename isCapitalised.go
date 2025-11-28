package main

func IsCapitalized(s string) bool {
	start := true
	for _, r := range s {
		if start {
			if r >= 'a' && r <= 'z' {
				return false
			}
			start = false
		}

		if r == ' ' {
			start = true
		}
	}
	return true
}

func main() {
	print(IsCapitalized("   Ade Ola"))
}
