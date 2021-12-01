package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	increases := 0
	lastNumber := math.MaxInt
	for input.Scan() {
		num, err := strconv.Atoi(input.Text())
		if err == nil {
			if num > lastNumber {
				increases++
			}
			lastNumber = num
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println(increases)
}
