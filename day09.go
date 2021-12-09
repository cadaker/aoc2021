package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grid = [][]int

func parseInput() (Grid, error) {
	var ret Grid
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		digits := strings.Split(input.Text(), "")
		var line []int
		for _, c := range digits {
			n, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			line = append(line, n)
		}
		ret = append(ret, line)
	}
	return ret, nil
}

func isMin(grid *Grid, row int, col int) bool {
	height := len(*grid)
	width := len((*grid)[0])
	n := (*grid)[row][col]
	return (row == 0 || n < (*grid)[row-1][col]) &&
			(row == height-1 || n < (*grid)[row+1][col]) &&
			(col == 0 || n < (*grid)[row][col-1]) &&
			(col == width - 1 || n < (*grid)[row][col+1])
}

func main() {
	grid, err := parseInput()
	if err != nil {
		fmt.Println("Failed to parse input:", err)
		return
	}

	totalRisk := 0
	height := len(grid)
	width := len(grid[0])
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if isMin(&grid, r, c) {
				totalRisk += grid[r][c] + 1
			}
		}
	}
	fmt.Println(totalRisk)
}
