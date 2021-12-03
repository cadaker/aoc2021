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

func countZerosInPosition(lines []string, index uint) int {
	zeroCount := 0
	for _, line := range lines {
		if line[index] == '0' {
			zeroCount++
		}
	}
	return zeroCount
}

func filter(lines []string, index uint, match uint8) []string {
	ret := make([]string, 0)
	for _, line := range lines {
		if line[index] == match {
			ret = append(ret, line)
		}
	}
	return ret
}

func binaryToInt(line string) int {
	num := 0
	for _, ch := range line {
		num *= 2
		if ch == '1' {
			num += 1
		}
	}
	return num
}

func lifeSupportRating(lines []string) int {
	oxygenLines := lines
	oxygenIndex := uint(0)
	for len(oxygenLines) > 1 {
		zeros := countZerosInPosition(oxygenLines, oxygenIndex)
		if zeros > len(oxygenLines) / 2 {
			oxygenLines = filter(oxygenLines, oxygenIndex, '0')
		} else {
			oxygenLines = filter(oxygenLines, oxygenIndex, '1')
		}
		oxygenIndex++
	}

	scrubberLines := lines
	scrubberIndex := uint(0)
	for len(scrubberLines) > 1 {
		zeros := countZerosInPosition(scrubberLines, scrubberIndex)
		if zeros > len(scrubberLines) / 2 {
			scrubberLines = filter(scrubberLines, scrubberIndex, '1')
		} else {
			scrubberLines = filter(scrubberLines, scrubberIndex, '0')
		}
		scrubberIndex++
	}
	return binaryToInt(oxygenLines[0]) * binaryToInt(scrubberLines[0])
}

func main() {
	lines := make([]string, 0)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		lines = append(lines, input.Text())
	}

	zeroCounts := countZeros(lines)

	fmt.Println(powerConsumption(zeroCounts, len(lines)))
	fmt.Println(lifeSupportRating(lines))
}
