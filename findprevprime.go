package main

func CheckPrime(nb int) bool {
	if nb == 2 {
		return true
	}

	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}

	return true
}

func FindPrevPrime(nb int) int {
	if nb <= 2 {
		return 0
	}

	for i := nb; i > 1; i-- {
		if CheckPrime(i) == true {
			return i
		}
	}
	return 0
}

func main() {
	print(FindPrevPrime(3))

}
