package main

import "github.com/01-edu/z01"

func main() {
	// z01.PrintRune('1')
	s := "Hello World!"
	for _, a := range s {
		z01.PrintRune(a)
	}
	z01.PrintRune('\n')
}
