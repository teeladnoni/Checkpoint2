package piscine

func ThirdTimeIsACharm(str string) string {
	// 1. Check if string is too short (optional, but good for speed)
	// Note: If you cannot use len(), you can remove this block.
	if len(str) < 3 {
		return "\n"
	}

	result := ""

	// 2. Use a manual counter
	index := 0

	for _, r := range str {
		// If index is 2 (3rd char), 5 (6th char), etc.
		// We use (index + 1) % 3 == 0 to find every 3rd character.
		if (index+1)%3 == 0 {
			result += string(r)
		}
		index++
	}

	return result + "\n"
}
