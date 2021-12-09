package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Point = struct {
	r int
	c int
}

func fillBasin(grid *Grid, fill *Grid, initialPoint Point, basinIndex int) {
	height := len(*grid)
	width := len((*grid)[0])
	queue := []Point{initialPoint}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if (*grid)[p.r][p.c] != 9 && (*fill)[p.r][p.c] == 0 {
			(*fill)[p.r][p.c] = basinIndex
			if p.r > 0 {
				queue = append(queue, Point{p.r-1, p.c})
			}
			if p.r < height-1 {
				queue = append(queue, Point{p.r+1, p.c})
			}
			if p.c > 0 {
				queue = append(queue, Point{p.r, p.c-1})
			}
			if p.c < width-1 {
				queue = append(queue, Point{p.r, p.c+1})
			}
		}
	}
}

func basinSizes(fill *Grid) map[int]int {
	sizes := map[int]int{}
	height := len(*fill)
	width := len((*fill)[0])
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			basinIndex := (*fill)[r][c]
			if basinIndex > 0 {
				sizes[basinIndex]++
			}
		}
	}
	return sizes
}

func answer(basinSizes map[int]int) int {
	var list []int
	for _, count := range basinSizes {
		list = append(list, count)
	}
	sort.Ints(list)
	length := len(list)
	return list[length-1] * list[length-2] * list[length-3]
}

func makeGrid(height int, width int) Grid {
	var g Grid
	for r := 0; r < height; r++ {
		g = append(g, make([]int, width))
	}
	return g
}

func main() {
	grid, err := parseInput()
	if err != nil {
		fmt.Println("Failed to parse input:", err)
		return
	}

	var lowPoints []Point
	totalRisk := 0
	height := len(grid)
	width := len(grid[0])
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if isMin(&grid, r, c) {
				totalRisk += grid[r][c] + 1
				lowPoints = append(lowPoints, Point{r, c})
			}
		}
	}
	fmt.Println(totalRisk)

	basinMap := makeGrid(height, width)
	for i, p := range lowPoints {
		fillBasin(&grid, &basinMap, p, i + 1)
	}
	sizes := basinSizes(&basinMap)
	fmt.Println(answer(sizes))
}
