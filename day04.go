package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board = [5][5]int

func parseBoard(input *bufio.Scanner) Board {
	var board Board
	for row := 0; row < 5; row++ {
		input.Scan()
		numbers := strings.Fields(input.Text())
		for col := 0; col < 5; col++ {
			n, _ := strconv.Atoi(numbers[col])
			board[row][col] = n
		}
	}
	return board
}

func mark(board *Board, call int) {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if board[row][col] == call {
				board[row][col] = -1
			}
		}
	}
}

func hasHorzBingo(board *Board, row int) bool {
	for col := 0; col < 5; col++ {
		if board[row][col] != -1 {
			return false
		}
	}
	return true
}

func hasVertBingo(board *Board, col int) bool {
	for row := 0; row < 5; row++ {
		if board[row][col] != -1 {
			return false
		}
	}
	return true
}

func hasHorzVertBingo(board *Board) bool {
	// Horizontal
	for row := 0; row < 5; row++ {
		if hasHorzBingo(board, row) {
			return true
		}
	}
	// Vertical
	for col := 0; col < 5; col++ {
		if hasVertBingo(board, col) {
			return true
		}
	}
	return false
}

func sumOfUnmarked(board Board) int {
	sum := 0
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if board[row][col] != -1 {
				sum += board[row][col]
			}
		}
	}
	return sum
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	if !input.Scan() {
		fmt.Println("Failed to read first line")
		return
	}
	var numbers []int
	for _, s := range strings.Split(input.Text(), ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Failed to parse number:", s)
		}
		numbers = append(numbers, n)
	}

	var boards []Board

	for input.Scan() {
		boards = append(boards, parseBoard(input))
	}

	for i, num := range numbers {
		for j := 0; j < len(boards); j++ {
			mark(&boards[j], num)
			if hasHorzVertBingo(&boards[j]) {
				fmt.Println("Bingo on board", j, "for number", i, "(", num, ")")
				fmt.Println(sumOfUnmarked(boards[j]) * num)
				return
			}
		}
	}
	fmt.Println("No bingo?!")
}
