package piscine

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	// The lookup table for manual Hex conversion
	base := "0123456789abcdef"

	// --- PART 1: PRINT HEX VALUES ---
	for i := 0; i < 10; i++ {

		// 1. Math to get the two hex characters
		val1 := arr[i] / 16
		val2 := arr[i] % 16

		// 2. Print them
		z01.PrintRune(rune(base[val1]))
		z01.PrintRune(rune(base[val2]))

		// 3. Formatting Rules (The tricky part)
		// We need newlines after the 4th, 8th, and 10th byte.
		// Arrays are 0-indexed, so that is index 3, 7, and 9.
		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune('\n')
		} else {
			// Otherwise, just put a space between the hex numbers
			z01.PrintRune(' ')
		}
	}

	// --- PART 2: PRINT ASCII CHARACTERS ---
	for i := 0; i < 10; i++ {
		// If it is a printable character (Standard ASCII table range)
		if arr[i] >= 32 && arr[i] <= 126 {
			z01.PrintRune(rune(arr[i]))
		} else {
			// If it's a weird number/symbol, print a dot
			z01.PrintRune('.')
		}
	}

	// Final newline to be clean
	z01.PrintRune('\n')
}
