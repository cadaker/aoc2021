package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseInput() ([]int, error) {
	var crabs []int
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fields := strings.Split(input.Text(), ",")
		for _, f := range fields {
			n, err := strconv.Atoi(f)
			if err == nil {
				crabs = append(crabs, n)
			} else {
				return crabs, err
			}
		}
	}
	return crabs, nil
}

func imin(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}


func iabs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func minElement(xs []int) int {
	min := math.MaxInt
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return min
}

func maxElement(xs []int) int {
	max := math.MinInt
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}

func alignmentFuel(crabs []int, pos int) int {
	fuel := 0
	for _, c := range crabs {
		fuel += iabs(c - pos)
	}
	return fuel
}

func main() {
	crabs, err := parseInput()
	if err != nil {
		fmt.Println("Failed to read input:", err)
		return
	}

	min := minElement(crabs)
	max := maxElement(crabs)

	minFuel := math.MaxInt
	for pos := min; pos <= max; pos++ {
		minFuel = imin(minFuel, alignmentFuel(crabs, pos))
	}
	fmt.Println(minFuel)
}
