package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var times int
	fmt.Fscan(in, &times)

	for i := 0; i < times; i++ {
		var word string
		fmt.Fscan(in, &word)

		chars := []rune(word)

		if isHex(chars) == true {
			num, _ := strconv.ParseInt(word, 16, 64)

			if isPrime(num) == true {
				fmt.Print("1/1\n")
			} else {
				fmt.Print("0/1\n")
			}
		} else if isBinary(chars) == true {
			numBi, _ := strconv.ParseInt(word, 2, 64)
			numDec, _ := strconv.ParseInt(word, 10, 64)
			numHex, _ := strconv.ParseInt(word, 16, 64)
			numOct, _ := strconv.ParseInt(word, 8, 64)

			stat := 0

			if isPrime(numBi) {
				stat++
			} else if isPrime(numDec) {
				stat++
			} else if isPrime(numHex) {
				stat++
			} else if isPrime(numOct) {
				stat++
			}

			switch stat {
			case 0:
				fmt.Print("0/1\n")
			case 1:
				fmt.Print("1/4\n")
			case 2:
				fmt.Print("1/2\n")
			case 3:
				fmt.Print("3/4\n")
			case 4:
				fmt.Print("1/1\n")
			}

		} else if isOctal(chars) == true {
			numOct, _ := strconv.ParseInt(word, 8, 64)
			numDec, _ := strconv.ParseInt(word, 10, 64)
			numHex, _ := strconv.ParseInt(word, 16, 64)

			stat := 0

			if isPrime(numDec) {
				stat++
			} else if isPrime(numHex) {
				stat++
			} else if isPrime(numOct) {
				stat++
			}

			switch stat {
			case 0:
				fmt.Print("0/1\n")
			case 1:
				fmt.Print("1/3\n")
			case 2:
				fmt.Print("2/3\n")
			case 3:
				fmt.Print("1/1\n")
			}

		} else {
			numDec, _ := strconv.ParseInt(word, 10, 64)
			numHex, _ := strconv.ParseInt(word, 16, 64)

			stat := 0

			if isPrime(numDec) {
				stat++
			} else if isPrime(numHex) {
				stat++
			}

			switch stat {
			case 0:
				fmt.Print("0/1\n")
			case 1:
				fmt.Print("1/2\n")
			case 2:
				fmt.Print("1/1\n")
			}
		}
	}
}

func isHex(list []rune) bool {
	var hex string = "ABCDEF"
	compare := []rune(hex)

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(compare); j++ {
			if compare[j] == list[i] {
				return true
			}
		}
	}
	return false
}

func isPrime(num int64) bool {
	if num == 1 {
		return false
	} else if num == 2 {
		return true
	} else if num%2 == 0 {
		return false
	}

	for i := int64(3); i*i <= num; i = i + 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func isBinary(list []rune) bool {
	var bi string = "10"
	compare := []rune(bi)

	for i := 0; i < len(list); i++ {
		if list[i] == compare[0] || list[i] == compare[1] {

		} else {
			return false
		}
	}
	return true
}

func isOctal(list []rune) bool {
	var oct string = "89"
	compare := []rune(oct)

	for i := 0; i < len(list); i++ {
		if list[i] == compare[0] || list[i] == compare[1] {
			return false
		}
	}

	return true
}
