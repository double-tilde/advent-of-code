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
	isLinear   bool
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

// TODO: Make ui nice for new search method
func shapedWordSearch(
	matrix [][]string,
	row, col int,
	si SearchInput,
	dir []int,
	sigChan chan os.Signal,
	ticker *time.Ticker,
) (bool, []WordPosition) {
	var validWord []WordPosition
	for i := 0; i < len(si.Word); i++ {

		if !isInBounds(matrix, row, col) || matrix[row][col] != string(si.Word[i]) {
			return false, nil
		}

		if matrix[row][col] == string(si.Word[i]) {
			validWord = append(
				validWord,
				WordPosition{Char: matrix[row][col], Row: row, Col: col},
			)
		}

		// uiMatrix := createUI(matrix, len(si.Word), curWord)
		// fmt.Println(uiMatrix)
		// ui.Create(uiMatrix, sigChan, ticker)

		row += dir[0]
		col += dir[1]
	}

	return true, validWord
}

func mappedWords(validWords [][]WordPosition) map[string][][]WordPosition {
	positionMap := make(map[string][][]WordPosition)

	for _, word := range validWords {
		for _, char := range word {
			if char.Char == "A" {
				key := fmt.Sprintf("%d-%d", char.Row, char.Col)
				positionMap[key] = append(positionMap[key], word)
				break
			}
		}
	}

	for k, v := range positionMap {
		if len(v) <= 1 {
			delete(positionMap, k)
		}
	}

	return positionMap
}

func getCount(matrix [][]string, si SearchInput, sigChan chan os.Signal, ticker *time.Ticker) int {
	count := 0
	var validWords [][]WordPosition

	for row := range matrix {
		for col := range matrix[row] {
			for _, dir := range si.Directions {
				if si.isLinear {
					if linearWordSearch(matrix, row, col, si, dir, sigChan, ticker) {
						count++
					}
				} else {
					ok, validWord := shapedWordSearch(matrix, row, col, si, dir, sigChan, ticker)
					if ok {
						validWords = append(validWords, validWord)
					}
				}
			}
		}
	}

	mappedWords := mappedWords(validWords)
	if len(validWords) > 0 {
		count = len(mappedWords)
	}

	return count
}

func fourthProblem() {
	input := utils.GetLineSeperatedRecords("./assets/04-file.txt")

	matrix := createMatrix(input)

	// si1 := SearchInput{
	// 	Word:       "XMAS",
	// 	Directions: [][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}},
	// 	isLinear:   true,
	// }

	si2 := SearchInput{
		Word:       "MAS",
		Directions: [][]int{{1, 1}, {-1, 1}, {-1, -1}, {1, -1}},
		isLinear:   false,
	}

	sigChan, ticker := ui.Setup(20)
	defer ticker.Stop()

	// res1 := getCount(matrix, si1, sigChan, ticker)

	res2 := getCount(matrix, si2, sigChan, ticker)
	fmt.Println("Fourth Problem:", res2)

	fmt.Println(ui.ShowCursor)
}
