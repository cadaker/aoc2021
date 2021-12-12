package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type graph map[string][]string

func parseInput() graph {
	var ret = graph{}
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		parts := strings.Split(input.Text(), "-")
		if len(parts) != 2 {
			panic("Invalid input line: " + input.Text())
		}
		ret[parts[0]] = append(ret[parts[0]], parts[1])
		ret[parts[1]] = append(ret[parts[1]], parts[0])
	}
	return ret
}

func isBigCave(s string) bool {
	return 'A' <= s[0] && s[0] <= 'Z'
}

func contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

func countPaths(g graph) int {
	count := 0
	queue := [][]string{{"start"}}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		currentPos := path[len(path)-1]
		if currentPos == "end" {
			count++
		} else {
			for _, next := range g[currentPos] {
				if isBigCave(next) || !contains(path, next) {
					newpath := append([]string{}, path...)
					queue = append(queue, append(newpath, next))
				}
			}
		}
	}
	return count
}

func main() {
	g := parseInput()

	fmt.Println(countPaths(g))
}
