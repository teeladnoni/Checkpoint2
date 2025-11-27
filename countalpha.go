package main

func CountAlpha(s string) int {
	count := 0
	for _, d := range s {
		if d >= 'a' && d <= 'z' || d >= 'A' && d <= 'Z' {
			count++
		}
	}
	return count
}

func main() {
	print(CountAlpha("abcDgfhyrt1256l"))

}
