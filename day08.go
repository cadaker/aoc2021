package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Signal = string
type Digit = string

type Input = struct {
	inputs []Signal
	outputs []Digit
}

func sortChars(s string) string {
	cs := strings.Split(s, "")
	sort.Strings(cs)
	return strings.Join(cs, "")
}

func sortCharsList(ss []string) []string {
	var ret []string
	for _, s := range ss {
		ret = append(ret, sortChars(s))
	}
	return ret
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
			sortCharsList(strings.Split(parts[0], " ")),
			sortCharsList(strings.Split(parts[1], " "))})
	}
	return ret, nil
}

func common(s, s2 string) string {
	var found []string
	for _, c := range strings.Split(s, "") {
		if strings.Index(s2, c) != -1 {
			found = append(found, c)
		}
	}
	return strings.Join(found, "")
}

func solve(input Input) map[string]int {
	mapping := map[string]int{}
	rmapping := map[int]string{}
	var fives []string
	var sixes []string
	for _, s := range input.inputs {
		if len(s) == 2 {
			mapping[s] = 1
			rmapping[1] = s
		} else if len(s) == 3 {
			mapping[s] = 7
			rmapping[7] = s
		} else if len(s) == 4 {
			mapping[s] = 4
			rmapping[4] = s
		} else if len(s) == 7 {
			mapping[s] = 8
			rmapping[8] = s
		} else if len(s) == 5 {
			fives = append(fives, s)
		} else if len(s) == 6 {
			sixes = append(sixes, s)
		} else {
			panic("Strange input length")
		}
	}
	if len(sixes) != 3 || len(fives) != 3 {
		panic("Something wrong with the input")
	}

	// Algorithm for finding the tricky 5- and 6-length digits
	// 6 & 1 != 1  identifies 6
	// 9 & 4 == 4  identifies 9
	// 3 & 1 == 1  identifies 3
	// 6 & 5 == 5  identifies 5

	for _, s := range sixes {
		if common(s, rmapping[1]) != rmapping[1] {
			mapping[s] = 6
			rmapping[6] = s
		} else if common(s, rmapping[4]) == rmapping[4] {
			mapping[s] = 9
			rmapping[9] = s
		} else {
			mapping[s] = 0
			rmapping[0] = s
		}
	}
	if len(rmapping[6]) == 0 {
		panic("Could not find digit 6")
	} else if len(rmapping[9]) == 0 {
		panic("Could not find digit 9")
	} else if len(rmapping[0]) == 0 {
		panic("Could not find digit 0")
	}
	for _, s := range fives {
		if common(s, rmapping[1]) == rmapping[1] {
			mapping[s] = 3
			rmapping[3] = s
		} else if common(s, rmapping[6]) == s {
			mapping[s] = 5
			rmapping[5] = s
		} else if s != rmapping[3] {
			mapping[s] = 2
			rmapping[2] = s
		}
	}
	if len(rmapping[3]) == 0 {
		panic("Could not find digit 3")
	} else if len(rmapping[5]) == 0 {
		panic("Could not find digit 5")
	} else if len(rmapping[2]) == 0 {
		panic("Could not find digit 2")
	}
	if len(mapping) != 10 || len(rmapping) != 10 {
		panic("Something did not get mapped")
	}
	return mapping
}

func solveNumber(input Input) int {
	mapping := solve(input)
	ret := 0
	for _, s := range input.outputs {
		ret = ret * 10 + mapping[s]
	}
	return ret
}

func main() {
	input, err := parseInput()
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	digits1478 := 0
	sum := 0
	for _, line := range input {
		for _, o := range line.outputs {
			if len(o) == 2 || len(o) == 4 || len(o) == 3 || len(o) == 7 {
				digits1478++
			}
		}
		sum += solveNumber(line)
	}
	fmt.Println(digits1478)
	fmt.Print(sum)
}
