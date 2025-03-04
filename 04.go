package main

import (
	"aoc-2024/ui"
	"aoc-2024/utils"
	"fmt"
	"os"
	"time"
)

type SearchInput struct {
	Word       string
	Directions [][]int
	Linear     bool
}

type WordPosition struct {
	Char string
	Row  int
	Col  int
}

func createMatrix(strs []string) [][]string {
	matrix := [][]string{}

	for _, str := range strs {
		row := []string{}
		for _, char := range str {
			row = append(row, string(char))
		}
		matrix = append(matrix, row)
	}

	return matrix
}

func isInBounds(matrix [][]string, row, col int) bool {
	return row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0])
}

func createUI(matrix [][]string, searchLen int, curWord []WordPosition) string {
	var uiMatrix string

	highlightMap := make(map[[2]int]string)

	for _, char := range curWord {
		if len(curWord) == searchLen {
			highlightMap[[2]int{char.Row, char.Col}] = ui.GreenBgBlackText
		} else {
			highlightMap[[2]int{char.Row, char.Col}] = ui.YellowBgBlackText
		}
	}

	for j := range matrix {
		for k := range matrix[j] {
			if color, exists := highlightMap[[2]int{j, k}]; exists {
				uiMatrix += ui.StringColor(matrix[j][k], color)
			} else {
				uiMatrix += matrix[j][k]
			}
		}
		uiMatrix += "\n"
	}

	return uiMatrix
}

func linearWordSearch(
	matrix [][]string,
	row, col int,
	si SearchInput,
	dir []int,
	sigChan chan os.Signal,
	ticker *time.Ticker,
) bool {
	curWord := []WordPosition{}

	for i := 0; i < len(si.Word); i++ {

		if !isInBounds(matrix, row, col) || matrix[row][col] != string(si.Word[i]) {
			return false
		}

		if matrix[row][col] == string(si.Word[i]) {
			curWord = append(curWord, WordPosition{Char: matrix[row][col], Row: row, Col: col})
		}

		uiMatrix := createUI(matrix, len(si.Word), curWord)
		ui.Create(uiMatrix, sigChan, ticker)

		row += dir[0]
		col += dir[1]
	}

	return true
}

func shapedWordSearch(
	matrix [][]string,
	row, col int,
	si SearchInput,
	sigChan chan os.Signal,
	ticker *time.Ticker,
) bool {
	dirs := si.Directions
	curWord := []WordPosition{}
	var cRow, cCol, oppRow, oppCol int

	for _, dir := range dirs {
		dRow := dir[0]
		dCol := dir[1]

		if !isInBounds(matrix, row, col) {
			return false
		}

		exists := false
		for _, char := range curWord {
			if char.Char == "A" {
				exists = true
			}
		}

		if dRow == 0 && dCol == 0 && matrix[row][col] == "A" {
			cRow = row
			cCol = col
			curWord = append(curWord, WordPosition{Char: matrix[cRow][cCol], Row: cRow, Col: cCol})
		}

		if !isInBounds(matrix, cRow+dRow, cCol+dCol) {
			return false
		}

		if exists && dRow == 1 && dCol == 1 && (matrix[cRow+dRow][cCol+dCol] == "M" ||
			matrix[cRow+dRow][cCol+dCol] == "S") {
			oppRow = cRow + dRow
			oppCol = cCol + dCol
			curWord = append(
				curWord,
				WordPosition{Char: matrix[oppRow][oppCol], Row: oppRow, Col: oppCol},
			)
		}

		if exists && dRow == -1 && dCol == -1 &&
			matrix[cRow+dRow][cCol+dCol] != matrix[oppRow][oppCol] &&
			(matrix[cRow+dRow][cCol+dCol] == "M" ||
				matrix[cRow+dRow][cCol+dCol] == "S") {
			curWord = append(
				curWord,
				WordPosition{
					Char: matrix[cRow+dRow][cCol+dCol],
					Row:  cRow + dRow,
					Col:  cCol + dCol,
				},
			)
		}

		if exists && dRow == -1 && dCol == 1 && (matrix[cRow+dRow][cCol+dCol] == "M" ||
			matrix[cRow+dRow][cCol+dCol] == "S") {
			oppRow = cRow + dRow
			oppCol = cCol + dCol
			curWord = append(
				curWord,
				WordPosition{Char: matrix[oppRow][oppCol], Row: oppRow, Col: oppCol},
			)
		}

		if exists && dRow == 1 && dCol == -1 &&
			matrix[cRow+dRow][cCol+dCol] != matrix[oppRow][oppCol] &&
			(matrix[cRow+dRow][cCol+dCol] == "M" ||
				matrix[cRow+dRow][cCol+dCol] == "S") {
			curWord = append(
				curWord,
				WordPosition{
					Char: matrix[cRow+dRow][cCol+dCol],
					Row:  cRow + dRow,
					Col:  cCol + dCol,
				},
			)
		}

		uiMatrix := createUI(matrix, len(dirs), curWord)
		ui.Create(uiMatrix, sigChan, ticker)
	}

	return len(curWord) == 5
}

func getCount(matrix [][]string, si SearchInput, sigChan chan os.Signal, ticker *time.Ticker) int {
	count := 0

	for row := range matrix {
		for col := range matrix[row] {
			if si.Linear {
				for _, dir := range si.Directions {
					if linearWordSearch(matrix, row, col, si, dir, sigChan, ticker) {
						count++
					}
				}
			} else {
				if shapedWordSearch(matrix, row, col, si, sigChan, ticker) {
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

	// si1 := SearchInput{
	// 	Word:       "XMAS",
	// 	Directions: [][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}},
	// 	Linear:     true,
	// }

	si2 := SearchInput{
		Word:       "MAS",
		Directions: [][]int{{0, 0}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}},
		Linear:     false,
	}

	sigChan, ticker := ui.Setup(10)
	defer ticker.Stop()

	res := getCount(matrix, si2, sigChan, ticker)

	fmt.Println("Fourth Problem:", res)
	fmt.Println(ui.ShowCursor)
}
