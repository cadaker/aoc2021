package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
)

func parseInput() (string, map[string]string, error) {
	input := bufio.NewScanner(os.Stdin)
	if !input.Scan() {
		return "", nil, errors.New("no first line")
	}
	template := input.Text()
	if !input.Scan() || input.Text() != "" {
		return "", nil, errors.New("expected an empty line")
	}

	rules := map[string]string{}
	for input.Scan() {
		parts := strings.Split(input.Text(), " -> ")
		if len(parts) != 2 || len(parts[0]) != 2 || len(parts[1]) != 1 {
			return "", nil, errors.New("bad split: " + input.Text())
		}
		rules[parts[0]] = parts[1]
	}

	return template, rules, nil
}

type counts map[string]int

type entry struct {
	pair string
	iters int
}

func pairExpansions(rules map[string]string, iters int) map[entry]counts {
	table := map[entry]counts{}
	for p, _ := range rules {
		table[entry{p, 0}] = counts{p[0:1]: 1, p[1:2]:1}
	}
	for iter := 1; iter <= iters; iter++ {
		for p, insert := range rules {
			newCounts := counts{}
			for char, n := range table[entry{p[0:1] + insert, iter-1}] {
				newCounts[char] += n
			}
			for char, n := range table[entry{insert + p[1:2], iter-1}] {
				newCounts[char] += n
			}
			newCounts[insert]--
			table[entry{p, iter}] = newCounts
		}
	}
	return table
}

func countExpanded(template string, table map[entry]counts, iters int) counts {
	expandedCounts := counts{}
	for i := 0; i+1 < len(template); i++ {
		for char, n := range table[entry{template[i:i+2], iters}] {
			expandedCounts[char] += n
		}
		if i > 0 {
			expandedCounts[template[i:i+1]]--
		}
	}
	return expandedCounts
}

func findResult(c counts) int {
	max := math.MinInt
	min := math.MaxInt
	for _, n := range c {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return max - min
}

func main() {
	template, rules, err := parseInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	table := pairExpansions(rules, 40)
	expandedCounts10 := countExpanded(template, table, 10)
	expandedCounts40 := countExpanded(template, table, 40)
	fmt.Println(findResult(expandedCounts10))
	fmt.Println(findResult(expandedCounts40))
}
