package main

func FishAndChips(n int) {
	if n < 0 {
		print("number is negative")
		return
	}

	if n%2 == 0 && n%3 == 0 {
		print("fish and chips")
	} else if n%2 == 0 {
		print("fish")
	} else if n%3 == 0 {
		print("chips")
	}

	return "non divisible"
}
