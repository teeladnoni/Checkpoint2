package main

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	negative := false
	if n < 0 {
		n = -n
		negative = true
	}

	if n < 10 {
		return string(rune(n) + '0')
	}
	words := ""
	for {
		dig := rune(n%10) + '0'
		words = string(dig) + words
		n = n / 10
		if n == 0 {
			break
		}
	}
	if negative {
		words = "-" + words
	}

	return words
}

func main() {
	print(Itoa(-564))
}
