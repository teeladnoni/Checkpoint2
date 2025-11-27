package main

func RetainFirstHalf(str string) string {
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str
	}

	newlen := len(str) / 2
	return str[:newlen]
}

func main() {
	print(RetainFirstHalf("hellor"))
}
