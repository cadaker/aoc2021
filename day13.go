package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

const (
	X = 0
	Y = 1
)

type fold struct {
	axis int
	coord int
}

type pointSet map[point]bool

func parseInput() (pointSet, []fold, error) {
	points := pointSet{}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "" {
			break
		}
		coords := strings.Split(input.Text(), ",")
		x, errx := strconv.Atoi(coords[0])
		y, erry := strconv.Atoi(coords[1])
		if errx != nil {
			return nil, nil, errx
		} else if erry != nil {
			return nil, nil, erry
		} else {
			points[point{x, y}] = true
		}
	}

	var folds []fold

	for input.Scan() {
		fields := strings.Fields(input.Text())
		if len(fields) != 3 {
			return nil, nil, errors.New("bad fold line: " + input.Text())
		}
		parts := strings.Split(fields[2], "=")
		if len(parts) != 2 {
			return nil, nil, errors.New("bad fold spec: " + parts[2])
		}
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}
		var axis int
		if parts[0] == "x" {
			axis = X
		} else {
			axis = Y
		}
		folds = append(folds, fold{axis, n})
	}
	return points, folds, nil
}

func foldX(points pointSet, coord int) pointSet {
	ret := pointSet{}
	for p, _ := range points {
		if p.x <= coord {
			ret[p] = true
		} else {
			ret[point{coord - (p.x - coord), p.y}] = true
		}
	}
	return ret
}

func foldY(points pointSet, coord int) pointSet {
	ret := pointSet{}
	for p, _ := range points {
		if p.y <= coord {
			ret[p] = true
		} else {
			ret[point{p.x, coord - (p.y - coord)}] = true
		}
	}
	return ret
}

func displayPoints(points pointSet) {
	maxX := math.MinInt
	maxY := math.MinInt
	for p, _ := range points {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if points[point{x,y}] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	points, folds, err := parseInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	if folds[0].axis == X {
		fmt.Println(len(foldX(points, folds[0].coord)))
	} else {
		fmt.Println(len(foldY(points, folds[0].coord)))
	}

	for _, f := range folds {
		if f.axis == X {
			points = foldX(points, f.coord)
		} else {
			points = foldY(points, f.coord)
		}
	}
	displayPoints(points)
}
