package main

import (
	"aoc-2024/ui"
	"aoc-2024/utils"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	reset  = "\033[0m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

func createMatrix(input []string) [][]rune {
	var matrix [][]rune

	for i := range input {
		var row []rune
		for j := range input[i] {
			row = append(row, rune(input[i][j]))
		}
		matrix = append(matrix, row)
	}

	return matrix
}

func isInBounds(matrix [][]rune, row, col int) bool {
	return row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0])
}

// TODO: Put the matched yellow chars in a string and when this is equal to word,
// make the ui show the green word
// Maybe they need to go in a map? something like that

func searchWord(
	matrix [][]rune,
	row, col int,
	dir []int,
	word string,
	sigChan chan os.Signal,
	ticker *time.Ticker,
) bool {
	for i := 0; i < len(word); i++ {
		var uiMatrix string
		curWord := make(map[string][]int)

		for j := range matrix {
			for k := range matrix[j] {
				if row == j && col == k {
					curWord[string(matrix[j][k])] = []int{j, k}
					for key := range curWord {
						uiMatrix += yellow + key + reset
					}
				} else {
					uiMatrix += string(matrix[j][k])
				}
			}
			uiMatrix += "\n"
		}

		ui.CreateUI(uiMatrix, sigChan, ticker)

		if !isInBounds(matrix, row, col) || matrix[row][col] != rune(word[i]) {
			return false
		}

		row += dir[0]
		col += dir[1]
	}
	return true
}

func getCount(matrix [][]rune) int {
	dirs := [][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	word := "XMAS"
	count := 0

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	ticker := time.NewTicker(time.Second / 3)
	defer ticker.Stop()
	fmt.Print(ui.HideCursor)

	for row := range matrix {
		for col := range matrix[row] {
			for _, dir := range dirs {
				if searchWord(matrix, row, col, dir, word, sigChan, ticker) {
					count++
				}
			}
		}
	}

	return count
}

func fourthProblem() {
	input := utils.GetLineSeperatedRecords("./assets/04-file.txt")

	matrix := createMatrix(input)
	res := getCount(matrix)

	fmt.Println("Fourth Problem:", res)
}
