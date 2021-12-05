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

func iabs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func isign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}

func intersects(l0 Line, l1 Line) []Point {
	len0 := imax(iabs(l0.P1.X - l0.P0.X), iabs(l0.P1.Y - l0.P0.Y))
	dx0 := isign(l0.P1.X - l0.P0.X)
	dy0 := isign(l0.P1.Y - l0.P0.Y)
	len1 := imax(iabs(l1.P1.X - l1.P0.X), iabs(l1.P1.Y - l1.P0.Y))
	dx1 := isign(l1.P1.X - l1.P0.X)
	dy1 := isign(l1.P1.Y - l1.P0.Y)

	points := []Point{}

	for k0 := 0; k0 <= len0; k0++ {
		x0 := l0.P0.X + k0*dx0
		y0 := l0.P0.Y + k0*dy0
		for k1 := 0; k1 <= len1; k1++ {
			x1 := l1.P0.X + k1*dx1
			y1 := l1.P0.Y + k1*dy1
			if x0 == x1 && y0 == y1 {
				points = append(points, Point{x0, y0})
			}
		}
	}
	return points
}

func main() {
	input, err := parseInput()
	if err != nil {
		fmt.Println("Failed to parse input:", err)
		return
	}

	allIntersections := map[Point]bool{}
	allAxisAlignedIntersections := map[Point]bool{}
	for i, line0 := range input {
		for _, line1 := range input[i+1:] {
			intersections := intersects(line0, line1)
			for _, p := range intersections {
				allIntersections[p] = true
				if isAxisAligned(line0) && isAxisAligned(line1) {
					allAxisAlignedIntersections[p] = true
				}
			}
		}
	}
	fmt.Println(len(allAxisAlignedIntersections))
	fmt.Println(len(allIntersections))
}
