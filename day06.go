package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fish timer -> # of fishes

type Fish = map[int]int

func parseInput() (Fish, error) {
	fish := Fish{}
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fields := strings.Split(input.Text(), ",")
		for _, f := range fields {
			n, err := strconv.Atoi(f)
			if err == nil {
				fish[n]++
			} else {
				return fish, err
			}
		}
	}
	return fish, nil
}

func iterate(fish Fish) Fish {
	newFish := Fish{}
	for i := 1; i <= 8; i++ {
		newFish[i-1] = fish[i]
	}
	newFish[6] += fish[0]
	newFish[8] = fish[0]
	return newFish
}

func totalFish(fish Fish) int {
	total := 0
	for _, count := range fish {
		total += count
	}
	return total
}

func main() {
	fish, err := parseInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 80; i++ {
		fish = iterate(fish)
	}
	fmt.Println(totalFish(fish))
}
