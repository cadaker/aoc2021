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

	for input.Scan() {
		parts := strings.Split(input.Text(), " ")
		if len(parts) == 2 {
			command := parts[0]
			amount, err := strconv.Atoi(parts[1])
			if err == nil {
				if command == "forward" {
					distance += amount
					continue
				} else if command == "down" {
					depth += amount
					continue
				} else if command == "up" {
					depth -= amount
					continue
				}
			}
		}
		fmt.Println("Failed to process line:", input.Text())
	}
	fmt.Println(depth*distance)
}
