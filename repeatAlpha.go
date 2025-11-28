package main

func RepeatAlpha(s string) string {
	words := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			rep := int(c-'a') + 1
			for i := 0; i < rep; i++ {
				words += string(c)
			}
		} else if c >= 'A' && c <= 'Z' {
			rep := int(c-'A') + 1
			for i := 0; i < rep; i++ {
				words += string(c)
			}
		} else {
			words += string(c)
		}
	}
	return words
}

func main() {
	print(RepeatAlpha("Efo"))
}
