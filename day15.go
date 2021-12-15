package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseInput() ([][]int, error) {
	input := bufio.NewScanner(os.Stdin)
	var grid [][]int
	for input.Scan() {
		var line []int
		digits := strings.Split(input.Text(), "")
		for _, d := range digits {
			n, err := strconv.Atoi(d)
			if err != nil {
				return nil, err
			}
			line = append(line, n)
		}
		grid = append(grid, line)
	}
	return grid, nil
}

type node struct {
	x int
	y int
	weight int
}

type heap struct {
	elems []node
}

func (h *heap) push(x int, y int, weight int) {
	h.elems = append(h.elems, node{x, y, weight})
	i := len(h.elems) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if h.elems[parent].weight > h.elems[i].weight {
			h.elems[parent], h.elems[i] = h.elems[i], h.elems[parent]
			i = parent
		} else {
			break
		}
	}
}

func (h *heap) empty() bool {
	return len(h.elems) == 0
}

func (h *heap) pop() node {
	if len(h.elems) == 0 {
		return node{}
	}
	ret := h.elems[0]
	h.elems[0] = h.elems[len(h.elems)-1]
	h.elems = h.elems[0:len(h.elems)-1]

	i := 0
	for true {
		child0 := i * 2 + 1
		child1 := i * 2 + 2
		if child0 < len(h.elems) && h.elems[i].weight > h.elems[child0].weight {
			h.elems[i], h.elems[child0] = h.elems[child0], h.elems[i]
			i = child0
		} else if child1 < len(h.elems) && h.elems[i].weight > h.elems[child1].weight {
			h.elems[i], h.elems[child1] = h.elems[child1], h.elems[i]
			i = child1
		} else {
			break
		}
	}
	return ret
}

type point struct {
	x int
	y int
}

func isBetter(bestPaths map[point]int, x int, y int, weight int) bool {
	b := bestPaths[point{x, y}]
	return b == 0 || weight < b
}

func gridPaths(grid [][]int, startx int, starty int) map[point]int {
	gridHeight := len(grid)
	gridWidth := len(grid[0])
	pqueue := heap{}
	pqueue.push(0, 0, 0)
	bestPaths := map[point]int{point{0,0}: math.MaxInt}
	for !pqueue.empty() {
		node := pqueue.pop()
		x := node.x
		y := node.y
		if isBetter(bestPaths, x, y, node.weight) {
			bestPaths[point{x, y}] = node.weight
			if x > 0 && isBetter(bestPaths, x-1, y, node.weight + grid[y][x-1]) {
				pqueue.push(x-1, y, node.weight + grid[y][x-1])
			}
			if x < gridWidth - 1 && isBetter(bestPaths, x+1, y, node.weight + grid[y][x+1]) {
				pqueue.push(x+1, y, node.weight + grid[y][x+1])
			}
			if y > 0 && isBetter(bestPaths, x, y-1, node.weight + grid[y-1][x]) {
				pqueue.push(x, y-1, node.weight + grid[y-1][x])
			}
			if y < gridHeight - 1 && isBetter(bestPaths, x, y+1, node.weight + grid[y+1][x]) {
				pqueue.push(x, y+1, node.weight + grid[y+1][x])
			}
		}
	}
	return bestPaths
}

func wrap(n int) int {
	if n > 9 {
		return n - 9
	} else {
		return n
	}
}

func extendGrid(grid [][]int) [][]int {
	height := len(grid)
	width := len(grid[0])
	var biggrid [][]int
	for rn := 0; rn < 5; rn++ {
		for r := 0; r < height; r++ {
			var line []int
			for cn := 0; cn < 5; cn++ {
				for c := 0; c < width; c++ {
					line = append(line, wrap(grid[r][c] + cn + rn))
				}
			}
			biggrid = append(biggrid, line)
		}
	}
	return biggrid
}

func main() {
	grid, err := parseInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	bestPaths := gridPaths(grid, 0, 0)
	height := len(grid)
	width := len(grid[0])
	fmt.Println(bestPaths[point{height-1, width-1}])
	biggrid := extendGrid(grid)
	bestPathsBig := gridPaths(biggrid, 0, 0)
	fmt.Println(bestPathsBig[point{height*5-1, width*5-1}])

}
