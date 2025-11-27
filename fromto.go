package main

import "strconv"

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid"
	}
	result := ""
	steps := 1
	if from > to {
		steps = -1
	}

	if from == to {
		if from < 10 {
			return "0" + strconv.Itoa(from)
		}
		return strconv.Itoa(from)
	}

	for i := from; ; i += steps {
		if i < 10 {
			result += "0" + strconv.Itoa(i)
		} else {
			result += strconv.Itoa(i)
		}
		if i == to {
			break
		}
		result += ", "
	}

	result += "\n"
	return result
}

func main() {
	print(FromTo(50, 30))
}
