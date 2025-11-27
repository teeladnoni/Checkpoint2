package main

func LastWord(s string) string {
	if len(s) == 0 {
		return ""
	}
	runes := []rune(s)
	end := len(runes) - 1

	for end >= 0 && runes[end] == ' ' {
		end--
	}

	if end < 0 {
		return ""
	}
	start := end
	for start >= 0 && runes[start] != ' ' {
		start--
	}

	words := string(runes[start+1 : end+1])
	words += "\n"
	return words
}

func main() {
	print(LastWord("Helloban Worldhjgdad  "))
}
