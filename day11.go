package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() [][]int {
	var grid [][]int
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		digits := strings.Split(input.Text(), "")
		var numbers []int
		for _, digit := range digits {
			n, err := strconv.Atoi(digit)
			if err != nil {
				panic("Failed to parse number: " + digit)
			}
			numbers = append(numbers, n)
		}
		grid = append(grid, numbers)
	}
	return grid
}

const FLASHED = -1

func flash(grid [][]int, row int, col int) {
	height := len(grid)
	width := len(grid[0])
	grid[row][col] = FLASHED
	for r := row-1; r <= row+1; r++ {
		for c := col-1; c <= col+1; c++ {
			if r >= 0 && c >= 0 && r < height && c < width && grid[r][c] != FLASHED {
				grid[r][c]++
			}
		}
	}
}

func step(grid [][]int) int {
	height := len(grid)
	width := len(grid[0])

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			grid[r][c]++
		}
	}

	flashes := 0
	boardChanged := true
	for boardChanged {
		boardChanged = false
		for r := 0; r < height; r++ {
			for c := 0; c < width; c++ {
				if grid[r][c] > 9 {
					flash(grid, r, c)
					flashes++
					boardChanged = true
				}
			}
		}
	}
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if grid[r][c] < 0 {
				grid[r][c] = 0
			}
		}
	}
	return flashes
}

func main() {
	grid := parseInput()

	totalFlashes := 0
	for i := 0; i < 100; i++ {
		totalFlashes += step(grid)
	}
	fmt.Println(totalFlashes)
	for iter := 101; true; iter++ {
		if step(grid) == 100 {
			fmt.Println(iter)
			break
		}
	}
}
