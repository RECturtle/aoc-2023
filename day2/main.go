package main

import (
	"fmt"
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
	var digitRegex = regexp.MustCompile(`\d+`)

	data, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error reading file: ", err)
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		countGame := true

		gameData := strings.Split(line, ": ")

		gameIndex, _ := strconv.Atoi(digitRegex.FindString(gameData[0]))
		gameScoreList := gameData[1:][0]

		gameCounts := strings.Split(gameScoreList, "; ")

		// Loop through the set of cube pulls to see if any break the max value for each color
	countCheck:
		for _, gameSet := range gameCounts {
			cubeCounts := strings.Split(gameSet, ", ")

			for _, cubeCount := range cubeCounts {
				countSet := strings.Split(cubeCount, " ")
				count, _ := strconv.Atoi(countSet[0])

				// Compare the number of cubes pulled for the color
				// against the max value in the map
				if count > maxBlockMap[countSet[1]] {
					countGame = false
					break countCheck
				}
			}
		}

		if countGame {
			sum += gameIndex
		}
	}
	fmt.Println(sum)
}
