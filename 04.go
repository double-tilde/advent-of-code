package main

import (
	"aoc-2024/ui"
	"aoc-2024/utils"
	"fmt"
	"os"
	"time"
)

type SearchInput struct {
	word       string
	directions [][]int
	linear     bool
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

func createSearchInput(w string, d [][]int, l bool) SearchInput {
	si := SearchInput{
		word:       w,
		directions: d,
		linear:     l,
	}

	return si
}

func isInBounds(matrix [][]string, row, col int) bool {
	return row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0])
}

func linearWordSearch(
	matrix [][]string,
	row, col int,
	si SearchInput,
	dir []int,
	sigChan chan os.Signal,
	ticker *time.Ticker,
) bool {
	curWord := make(map[string][]int)

	for i := 0; i < len(si.word); i++ {

		var uiMatrix string

		if !isInBounds(matrix, row, col) || matrix[row][col] != string(si.word[i]) {
			return false
		}

		if matrix[row][col] == string(si.word[i]) {
			curWord[matrix[row][col]] = []int{row, col}
		}

		for j := range matrix {
			for k := range matrix[j] {
				curChar := matrix[j][k]

				val, exists := curWord[curChar]
				if exists && len(curWord) == len(si.word) && val[0] == j &&
					val[1] == k {
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

func shapedWordSearch(
	matrix [][]string,
	row, col int,
	si SearchInput,
	sigChan chan os.Signal,
	ticker *time.Ticker,
) bool {
	dirs := si.directions

	curWord := make(map[string][]int)
	var cRow, cCol, oppRow, oppCol int

	for _, dir := range dirs {
		dRow := dir[0]
		dCol := dir[1]

		if !isInBounds(matrix, row, col) {
			return false
		}

		_, exists := curWord["A"]

		if dRow == 0 && dCol == 0 && matrix[row][col] == "A" {
			cRow = row
			cCol = col
			curWord[matrix[cRow][cCol]] = []int{cRow, cCol}
		}

		if !isInBounds(matrix, cRow+dRow, cCol+dCol) {
			return false
		}

		if exists && dRow == 1 && dCol == 1 && (matrix[cRow+dRow][cCol+dCol] == "M" ||
			matrix[cRow+dRow][cCol+dCol] == "S") {
			oppRow = cRow + dRow
			oppCol = cCol + dCol
			curWord[matrix[oppRow][oppCol]+"11"] = []int{oppRow, oppCol}
		}

		if exists && dRow == -1 && dCol == -1 &&
			matrix[cRow+dRow][cCol+dCol] != matrix[oppRow][oppCol] &&
			(matrix[cRow+dRow][cCol+dCol] == "M" ||
				matrix[cRow+dRow][cCol+dCol] == "S") {
			curWord[matrix[cRow+dRow][cCol+dCol]+"-1-1"] = []int{cRow + dRow, cCol + dCol}
		}

		if exists && dRow == -1 && dCol == 1 && (matrix[cRow+dRow][cCol+dCol] == "M" ||
			matrix[cRow+dRow][cCol+dCol] == "S") {
			oppRow = cRow + dRow
			oppCol = cCol + dCol
			curWord[matrix[oppRow][oppCol]+"-11"] = []int{oppRow, oppCol}
		}

		if exists && dRow == 1 && dCol == -1 &&
			matrix[cRow+dRow][cCol+dCol] != matrix[oppRow][oppCol] &&
			(matrix[cRow+dRow][cCol+dCol] == "M" ||
				matrix[cRow+dRow][cCol+dCol] == "S") {
			curWord[matrix[cRow+dRow][cCol+dCol]+"1-1"] = []int{cRow + dRow, cCol + dCol}
		}

	}

	if len(curWord) != 5 {
		return false
	}

	return true
}

func getCount(matrix [][]string, si SearchInput, sigChan chan os.Signal, ticker *time.Ticker) int {
	count := 0

	for row := range matrix {
		for col := range matrix[row] {
			if si.linear {
				for _, dir := range si.directions {
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

	// si1 := createSearchInput(
	// 	"XMAS",
	// 	[][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}},
	// 	true,
	// )

	si2 := createSearchInput("MAS", [][]int{{0, 0}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}, false)
	// si2 := createSearchInput("MAS", [][]int{{0, 0}, {1, 1}, {-1, -1}}, false)

	sigChan, ticker := ui.Setup(10)
	defer ticker.Stop()

	res := getCount(matrix, si2, sigChan, ticker)

	fmt.Println("Fourth Problem:", res)
	fmt.Println(ui.ShowCursor)
}
