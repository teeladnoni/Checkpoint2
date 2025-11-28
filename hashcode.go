package main

func HashCode(dec string) string {
	if len(dec) == 0 {
		return ""
	}

	lent := len(dec)
	var result string

	for _, f := range dec {
		cal := (int(f) + lent) % 127
		if cal < 33 {
			cal += 33
		}
		result += string(rune(cal))
	}
	return result
}
