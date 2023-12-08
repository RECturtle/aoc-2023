package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var NumLookup = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func readFiles(data []byte) []string {
	return strings.Split(string(data), "\n")
}

func partOne(lines []string) {
	var sum int
	for _, v := range lines {
		codes := regexp.MustCompile(`[^\d]+`).ReplaceAllString(v, "")

		strValue := string(codes[0]) + string(codes[len(codes)-1])
		calibrationValue, err := strconv.Atoi(strValue)

		if err != nil {
			fmt.Println("Unable to convert str to int:", err)
		}

		sum += calibrationValue
	}
	fmt.Printf("Day 1: %d\n", sum)
}

func partTwo(lines []string) {
	var sum int

	for _, line := range lines {
		var matches []string

		re := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)

		for i := range line {
			match := re.FindString(line[i:])

			if match == "" {
				continue
			}

			if m, ok := NumLookup[match]; ok {
				match = m
			}

			matches = append(matches, match)
		}

		strValue := string(matches[0]) + string(matches[len(matches)-1])
		calibrationValue, err := strconv.Atoi(strValue)

		if err != nil {
			fmt.Println("Unable to convert str to int:", err)
		}

		sum += calibrationValue
	}
	fmt.Printf("Day 2: %d\n", sum)
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error reading file: ", err)
	}

	codes := readFiles(data)

	partOne(codes)
	partTwo(codes)
}
