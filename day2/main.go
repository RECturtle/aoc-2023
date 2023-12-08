package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var maxBlockMap = map[string]int{
	"blue":  14,
	"green": 13,
	"red":   12,
}

func main() {
	var sum int
	var minSum int
	var digitRegex = regexp.MustCompile(`\d+`)

	data, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		var minBlockCount = make(map[string]int)
		countGame := true
		power := 1

		gameData := strings.Split(line, ": ")
		gameIndex, _ := strconv.Atoi(digitRegex.FindString(gameData[0]))
		gameScoreList := gameData[1:][0]
		gameCounts := strings.Split(gameScoreList, "; ")

		for _, gameSet := range gameCounts {
			cubeCounts := strings.Split(gameSet, ", ")

			for _, cubeCount := range cubeCounts {
				countSet := strings.Split(cubeCount, " ")
				color := countSet[1]
				count, _ := strconv.Atoi(countSet[0])

				// Set minBlockCount for the current color
				if m, ok := minBlockCount[color]; ok {
					minBlockCount[color] = int(math.Max(float64(m), float64(count)))
				} else {
					minBlockCount[color] = count
				}

				// If too many blocks pulled, don't count the index towards total
				if count > maxBlockMap[color] {
					countGame = false
				}
			}
		}

		// Calculate values for answers
		for _, v := range minBlockCount {
			power *= v
		}
		minSum += power

		if countGame {
			sum += gameIndex
		}
	}

	fmt.Println("Part One:", sum)
	fmt.Println("Part Two:", minSum)
}
