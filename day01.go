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

	windowIncreases := 0
	window0 := -1
	window1 := -1
	window2 := -1

	for input.Scan() {
		num, err := strconv.Atoi(input.Text())
		if err == nil {
			if num > lastNumber {
				increases++
			}
			lastNumber = num

			if window0 != -1 && window1 != -1 && window2 != -1 {
				if window1 + window2 + num > window0 + window1 + window2 {
					windowIncreases++
				}
			}
			window0 = window1
			window1 = window2
			window2 = num
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println(increases)
	fmt.Println(windowIncreases)
}
