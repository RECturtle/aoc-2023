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
	var partOneSum int
	var partTwoSum int
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
		allGameSets := strings.Split(gameScoreList, "; ")

		for _, gameSet := range allGameSets {
			cubeCounts := strings.Split(gameSet, ", ")

			for _, cubeCount := range cubeCounts {
				countSet := strings.Split(cubeCount, " ")
				color := countSet[1]
				count, _ := strconv.Atoi(countSet[0])

				partOne(color, count, &countGame)
				partTwo(color, count, minBlockCount)
			}
		}

		// Calculate values for answers
		for _, v := range minBlockCount {
			power *= v
		}

		partTwoSum += power

		if countGame {
			partOneSum += gameIndex
		}
	}

	fmt.Println("Part One:", partOneSum)
	fmt.Println("Part Two:", partTwoSum)
}

func partOne(color string, count int, countGame *bool) {
	if count > maxBlockMap[color] {
		*countGame = false
	}
}

func partTwo(color string, count int, minBlockCount map[string]int) {
	if m, ok := minBlockCount[color]; ok {
		minBlockCount[color] = int(math.Max(float64(m), float64(count)))
	} else {
		minBlockCount[color] = count
	}
}
