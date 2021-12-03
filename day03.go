package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	totalLines := 0
	zeroCounts := make(map[int]int)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		totalLines++
		line := input.Text()
		for i := 0; i < len(line); i++ {
			if line[i] == '0' {
				zeroCounts[i]++
			}
		}
	}

	wordSize := len(zeroCounts)
	gamma := 0
	epsilon := 0
	for i := 0; i < wordSize; i++ {
		gamma *= 2
		epsilon *= 2
		if zeroCounts[i] > totalLines/2 {
			epsilon += 1
		} else {
			gamma += 1
		}
	}

	fmt.Println(epsilon * gamma)
}
