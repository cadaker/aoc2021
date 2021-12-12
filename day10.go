package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func parseInput() []string {
	var lines []string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		lines = append(lines, input.Text())
	}
	return lines
}

// Returns a pair of the first invalid character (or 0), followed by the
// paren stack at that point, or at the end
func checkLine(line string) (uint8, []uint8) {
	closer := map[uint8]uint8{'(':')', '{':'}', '[':']', '<':'>'}
	if len(line) == 0 {
		return 0, []uint8{}
	}
	var stack []uint8
	for i := 0; i < len(line); i++ {
		isOpener := closer[line[i]] != 0
		if isOpener {
			stack = append(stack, line[i])
		} else {
			opener := stack[len(stack)-1]
			if len(stack) == 0 || line[i] != closer[opener] {
				return line[i], stack
			} else {
				stack = stack[0:len(stack)-1]
			}
		}
	}
	return 0, stack
}

func computeCompletionScore(stack []uint8) int {
	scores := map[uint8]int{'(': 1, '[':2, '{':3, '<':4}
	score := 0
	for i := len(stack) - 1; i >= 0; i-- {
		score = score * 5 + scores[stack[i]]
	}
	return score
}

func getMiddleScore(scores []int) int {
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func main() {
	lines := parseInput()

	scoring := map[uint8]int{')':3, ']':57, '}':1197, '>':25137}
	totalScore := 0
	var completionScores []int
	for _, line := range lines {
		check, stack := checkLine(line)
		if check == 0 {
			completionScores = append(completionScores, computeCompletionScore(stack))
		} else if scoring[check] > 0 {
			totalScore += scoring[check]
		}
	}
	fmt.Println(totalScore)
	fmt.Println(getMiddleScore(completionScores))
}
