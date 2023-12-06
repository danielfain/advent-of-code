package day1

import (
	"strings"
	"unicode"
)

func sumPart1(calibrationDocument string) (sum int) {
	lines := strings.Split(calibrationDocument, "\n")

	for _, line := range lines {
		sum += computeSum(line)
	}

	return sum
}

func sumPart2(calibrationDocument string) (sum int) {
	lines := strings.Split(calibrationDocument, "\n")

	numbersMap := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	for _, line := range lines {
		line = replaceFirstNumberWord(line, numbersMap)
		line = replaceLastNumberWord(line, numbersMap)
		sum += computeSum(line)
	}

	return sum
}

func replaceFirstNumberWord(line string, numbersMap map[string]string) string {
	for i := 0; i < len(line); i++ {
		for j := i + 1; j <= len(line); j++ {
			substring := line[i:j]
			if val, contains := numbersMap[substring]; contains {
				return strings.Replace(line, substring, val, 1)
			}
		}
	}
	return line
}

func replaceLastNumberWord(line string, numbersMap map[string]string) string {
	for i := len(line); i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			substring := line[j:i]
			if val, contains := numbersMap[substring]; contains {
				return strings.Replace(line, substring, val, 1)
			}
		}
	}
	return line
}

func computeSum(line string) int {
	var first int
	var last int

	for _, char := range line {
		if first == 0 && unicode.IsNumber(char) {
			first = int(char - '0')
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		char := rune(line[i])
		if unicode.IsNumber(char) {
			last = int(char - '0')
			break
		}
	}

	return (first * 10) + last
}
