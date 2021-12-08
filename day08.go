package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Signal = string
type Digit = string

type Input = struct {
	inputs []Signal
	outputs []Digit
}

func parseInput() ([]Input, error) {
	var ret []Input
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		parts := strings.Split(input.Text(), " | ")
		if len(parts) != 2 {
			return nil, errors.New("Bad format in input line: " + input.Text())
		}
		ret = append(ret, Input{
			strings.Split(parts[0], " "),
			strings.Split(parts[1], " ")})
	}
	return ret, nil
}

func main() {
	input, err := parseInput()
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	digits1478 := 0
	for _, line := range input {
		for _, o := range line.outputs {
			if len(o) == 2 || len(o) == 4 || len(o) == 3 || len(o) == 7 {
				digits1478++
			}
		}
	}
	fmt.Println(digits1478)
}
