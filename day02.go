package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	depth := 0
	distance := 0

	aim := 0
	aimDistance := 0
	aimDepth := 0

	for input.Scan() {
		parts := strings.Split(input.Text(), " ")
		if len(parts) == 2 {
			command := parts[0]
			amount, err := strconv.Atoi(parts[1])
			if err == nil {
				if command == "forward" {
					distance += amount
					aimDistance += amount
					aimDepth += aim * amount
					continue
				} else if command == "down" {
					depth += amount
					aim += amount
					continue
				} else if command == "up" {
					depth -= amount
					aim -= amount
					continue
				}
			}
		}
		fmt.Println("Failed to process line:", input.Text())
	}
	fmt.Println(depth*distance)
	fmt.Println(aimDepth*aimDistance)
}
