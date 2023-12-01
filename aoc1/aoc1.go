package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var digitsRE = regexp.MustCompile(`\d|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)`)

func digits(line string) []int {
	var is = []int{}
	for _, s := range digitsRE.FindAllString(line, -1) {
		var i int
		switch s {
		case "1", "one":
			i = 1
		case "2", "two":
			i = 2
		case "3", "three":
			i = 3
		case "4", "four":
			i = 4
		case "5", "five":
			i = 5
		case "6", "six":
			i = 6
		case "7", "seven":
			i = 7
		case "8", "eight":
			i = 8
		case "9", "nine":
			i = 9
		}
		is = append(is, i)
	}
	return is
}

func Calibrate(values []string) int {
	result := 0
	for _, s := range values {
		ds := digits(s)
		result += 10*(ds[0]) + ds[len(ds)-1]
	}
	return result
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println(Calibrate(lines))
}
