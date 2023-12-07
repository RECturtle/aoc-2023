package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		fmt.Println(line)
	}
}
