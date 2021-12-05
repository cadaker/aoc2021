package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	P0 Point
	P1 Point
}

func normalizeLine(line Line) Line {
	if line.P0.X > line.P1.X {
		return Line{line.P1, line.P0}
	} else if line.P0.X == line.P1.X && line.P0.Y > line.P1.Y {
		return Line{line.P1, line.P0}
	} else {
		return line
	}
}

func parsePoint(s string) (Point, error) {
	xy := strings.Split(s, ",")
	if len(xy) != 2 {
		return Point{}, errors.New("could not split coordinates: " + s)
	}
	x, errx := strconv.Atoi(xy[0])
	y, erry := strconv.Atoi(xy[1])
	if errx != nil {
		return Point{}, errx
	} else if erry != nil  {
		return Point{}, erry
	} else {
		return Point{x, y}, nil
	}
}

func parseInput() ([]Line, error) {
	var lines []Line
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		parts := strings.Split(input.Text(), " -> ")
		if len(parts) != 2 {
			return lines, errors.New("could not split line: " + input.Text())
		}
		p0, err0 := parsePoint(parts[0])
		p1, err1 := parsePoint(parts[1])
		if err0 != nil {
			return lines, err0
		} else if err1 != nil {
			return lines, err1
		} else {
			lines = append(lines, normalizeLine(Line{p0, p1}))
		}
	}
	return lines, nil
}

func isHorizontal(line Line) bool {
	return line.P0.Y == line.P1.Y
}

func isVertical(line Line) bool {
	return line.P0.X == line.P1.X
}

func isAxisAligned(line Line) bool {
	return isVertical(line) || isHorizontal(line)
}

func imin(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func imax(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func xrange(y int, x0 int, x1 int) []Point {
	var points []Point
	for x := x0; x <= x1; x++ {
		points = append(points, Point{x, y})
	}
	return points
}

func yrange(x int, y0 int, y1 int) []Point {
	var points []Point
	for y := y0; y <= y1; y++ {
		points = append(points, Point{x, y})
	}
	return points
}

func simpleIntersects(l0 Line, l1 Line) []Point {
	if isHorizontal(l0) && isHorizontal(l1) && l0.P0.Y == l1.P0.Y {
		if l0.P0.X > l1.P0.X {
			return simpleIntersects(l1, l0)
		} else {
			x0 := imax(l0.P0.X, l1.P0.X)
			x1 := imin(l0.P1.X, l1.P1.X)
			return xrange(l0.P0.Y, x0, x1)
		}
	} else if isVertical(l0) && isVertical(l1) && l0.P0.X == l1.P0.X {
		if l0.P0.Y > l1.P0.Y {
			return simpleIntersects(l1, l0)
		} else {
			y0 := imax(l0.P0.Y, l1.P0.Y)
			y1 := imin(l0.P1.Y, l1.P1.Y)
			return yrange(l0.P0.X, y0, y1)
		}
	} else if isHorizontal(l0) && isVertical(l1) {
		return simpleIntersects(l1, l0)
	} else if isVertical(l0) && isHorizontal(l1) {
		x0 := l0.P0.X
		y1 := l1.P0.Y
		if l1.P0.X <= x0 && x0 <= l1.P1.X && l0.P0.Y <= y1 && y1 <= l0.P1.Y {
			return []Point{Point{x0, y1}}
		} else {
			return []Point{}
		}
	} else {
		return []Point{}
	}
}

func main() {
	input, err := parseInput()
	if err != nil {
		fmt.Println("Failed to parse input:", err)
		return
	}

	allIntersections := map[Point]bool{}
	for i, line0 := range input {
		for _, line1 := range input[i+1:] {
			if isAxisAligned(line0) && isAxisAligned(line1) {
				intersections := simpleIntersects(line0, line1)
				for _, p := range intersections {
					allIntersections[p] = true
				}
			}
		}
	}
	fmt.Println(len(allIntersections))
}
