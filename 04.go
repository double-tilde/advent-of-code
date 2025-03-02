package main

import (
	"aoc-2024/ui"
	"aoc-2024/utils"
	"fmt"
	"os"
	"time"
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

func searchWord(
	matrix [][]rune,
	row, col int,
	dir []int,
	word string,
	sigChan chan os.Signal,
	ticker *time.Ticker,
) bool {
	curWord := make(map[string][]int)

	for i := 0; i < len(word); i++ {
		if !isInBounds(matrix, row, col) || matrix[row][col] != rune(word[i]) {
			return false
		}

		var uiMatrix string

		if matrix[row][col] == rune(word[i]) {
			curWord[string(matrix[row][col])] = []int{row, col}
		}

		for j := range matrix {
			for k := range matrix[j] {
				curChar := string(matrix[j][k])

				val, exists := curWord[string(matrix[j][k])]
				if exists && len(curWord) == len(word) && val[0] == j && val[1] == k {
					uiMatrix += ui.StringColor(curChar, ui.GreenBgBlackText)
				} else if exists && val[0] == j && val[1] == k {
					uiMatrix += ui.StringColor(curChar, ui.YellowBgBlackText)
				} else {
					uiMatrix += curChar
				}
			}
			uiMatrix += "\n"
		}

		ui.Create(uiMatrix, sigChan, ticker)

		row += dir[0]
		col += dir[1]
	}

	return true
}

func getCount(matrix [][]rune) int {
	dirs := [][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	word := "XMAS"
	count := 0

	sigChan, ticker := ui.Setup(10)
	defer ticker.Stop()

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
