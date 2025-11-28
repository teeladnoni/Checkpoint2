package main

func PrintIf(str string) {
	if len(str) >= 3 || len(str) == 0 {
		print("G\n")
	}
	print("invalid input")
}
