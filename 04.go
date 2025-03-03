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
	uiMatrix string,
	si SearchInput,
	curWord map[string][]int,
) string {
	for j := range matrix {
		for k := range matrix[j] {
			curChar := string(matrix[j][k])

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
	return uiMatrix
}

func shapedWordSearch(
	matrix [][]string,
	uiMatrix string,
	si SearchInput,
	curWord map[string][]int,
) string {
	return uiMatrix
}

func searchWord(
	matrix [][]string,
	row, col int,
	si SearchInput,
	dir []int,
	sigChan chan os.Signal,
	ticker *time.Ticker,
) bool {
	curWord := make(map[string][]int)

	for i := 0; i < len(si.word); i++ {
		if !isInBounds(matrix, row, col) || matrix[row][col] != string(si.word[i]) {
			return false
		}

		var uiMatrix string

		if matrix[row][col] == string(si.word[i]) {
			curWord[string(matrix[row][col])] = []int{row, col}
		}

		if si.linear {
			uiMatrix = linearWordSearch(matrix, uiMatrix, si, curWord)
		} else {
			uiMatrix = shapedWordSearch(matrix, uiMatrix, si, curWord)
		}

		ui.Create(uiMatrix, sigChan, ticker)

		row += dir[0]
		col += dir[1]
	}

	return true
}

func getCount(matrix [][]string, si SearchInput, sigChan chan os.Signal, ticker *time.Ticker) int {
	count := 0

	for row := range matrix {
		for col := range matrix[row] {
			for _, dir := range si.directions {
				if searchWord(matrix, row, col, si, dir, sigChan, ticker) {
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

	si1 := createSearchInput(
		"XMAS",
		[][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}},
		true,
	)

	sigChan, ticker := ui.Setup(10)
	defer ticker.Stop()

	res := getCount(matrix, si1, sigChan, ticker)

	fmt.Println("Fourth Problem:", res)
	fmt.Println(ui.ShowCursor)
}
