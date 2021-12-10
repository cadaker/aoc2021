package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseInput() []string {
	var lines []string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		lines = append(lines, input.Text())
	}
	return lines
}

// Returns the first invalid character, or 0 if the line is ok
func checkLine(line string) uint8 {
	closer := map[uint8]uint8{'(':')', '{':'}', '[':']', '<':'>'}
	if len(line) == 0 {
		return 0
	}
	var stack []uint8
	for i := 0; i < len(line); i++ {
		isOpener := closer[line[i]] != 0
		if isOpener {
			stack = append(stack, line[i])
		} else {
			opener := stack[len(stack)-1]
			if len(stack) == 0 || line[i] != closer[opener] {
				return line[i]
			} else {
				stack = stack[0:len(stack)-1]
			}
		}
	}
	return 0
}

func main() {
	lines := parseInput()

	scoring := map[uint8]int{')':3, ']':57, '}':1197, '>':25137}
	totalScore := 0
	for _, line := range lines {
		check := checkLine(line)
		if scoring[check] > 0 {
			totalScore += scoring[check]
		}
	}
	fmt.Println(totalScore)
}
