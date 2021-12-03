package main

import (
	"bufio"
	"fmt"
	"os"
)

func countZeros(lines []string) map[int]int {
	ret := make(map[int]int)
	for _, line := range lines {
		for col, ch := range line {
			if ch == '0' {
				ret[col]++
			}
		}
	}
	return ret
}

func powerConsumption(zeroCounts map[int]int, wordCount int) int {
	wordSize := len(zeroCounts)
	gamma := 0
	epsilon := 0
	for i := 0; i < wordSize; i++ {
		gamma *= 2
		epsilon *= 2
		if zeroCounts[i] > wordCount/2 {
			epsilon += 1
		} else {
			gamma += 1
		}
	}
	return gamma * epsilon
}

func main() {
	lines := make([]string, 0)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		lines = append(lines, input.Text())
	}

	zeroCounts := countZeros(lines)

	fmt.Println(powerConsumption(zeroCounts, len(lines)))
}
