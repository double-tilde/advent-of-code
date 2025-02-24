package main

import (
	"aoc-2024/utils"
	"fmt"
)

// 2. start at the 0, 0 position, check left
// 3. if the next letter is what we expect, continue that direction
// 4. if not, move to left down, repeat, if not, moce to down, and so on
// 5. when a full word is found, increase the count of valid words found

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

func searchWord(matrix [][]rune, startRow, startCol int, dir []int, word string) bool {
	row, col := startRow, startCol

	for i := 0; i < len(word); i++ {
		if !isInBounds(matrix, row, col) || matrix[row][col] != rune(word[i]) {
			return false
		}
		row += dir[0]
		col += dir[1]
	}
	return true
}

func searchMatrix(matrix [][]rune) int {
	// t, tr, r, br, b, bl, l, tl
	dirs := [][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	word := "XMAS"
	count := 0

	for row := range matrix {
		for col := range matrix[row] {
			for _, dir := range dirs {
				if searchWord(matrix, row, col, dir, word) {
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
	res := searchMatrix(matrix)

	fmt.Println("Fourth Problem:", res)
}
