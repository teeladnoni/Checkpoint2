package piscine

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	// Validation check
	if !isValidCamelCase(s) {
		return s
	}

	result := ""

	for i, char := range s {
		// If it is an uppercase letter and it is NOT the first character
		// append an underscore first.
		if char >= 'A' && char <= 'Z' {
			if i != 0 {
				result += "_"
			}
		}
		result += string(char)
	}

	return result
}

func isValidCamelCase(s string) bool {
	// We use a variable to store the previous character safely
	// instead of accessing s[i-1] which can be dangerous.
	var prevChar rune

	for i, char := range s {
		// 1. Check for numbers
		if char >= '0' && char <= '9' {
			return false
		}

		// 2. Check for symbols (must be letters only)
		if (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') {
			return false
		}

		// 3. Check Uppercase specific rules
		if char >= 'A' && char <= 'Z' {
			// Rule: Cannot be the last character
			if i == len(s)-1 {
				return false
			}

			// Rule: The character before cannot be Uppercase
			// We use prevChar here instead of s[i-1]
			if i > 0 {
				if prevChar >= 'A' && prevChar <= 'Z' {
					return false
				}
			}
		}
		
		// Save current char as 'previous' for the next loop iteration
		prevChar = char
	}

	return true
}
