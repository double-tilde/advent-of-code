package main

import (
	"aoc-2024/get"
	"aoc-2024/model"
	"aoc-2024/ui"
	"fmt"
)

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

func validPairs(positionMap map[string][][]model.WordPosition) map[string][][]model.WordPosition {
	for k, v := range positionMap {
		if len(v) <= 1 {
			delete(positionMap, k)
		}
	}
	return positionMap
}

func wordSearch(
	matrix [][]string,
	row, col int,
	si model.SearchInput,
	dir []int,
) (bool, []model.WordPosition) {
	var validWord []model.WordPosition
	for i := 0; i < len(si.Word); i++ {

		if !isInBounds(matrix, row, col) || matrix[row][col] != string(si.Word[i]) {
			return false, nil
		}

		if matrix[row][col] == string(si.Word[i]) {
			validWord = append(
				validWord,
				model.WordPosition{Char: matrix[row][col], Row: row, Col: col},
			)
		}

		row += dir[0]
		col += dir[1]
	}

	return true, validWord
}

func mappedWords(validWords [][]model.WordPosition) map[string][][]model.WordPosition {
	positionMap := make(map[string][][]model.WordPosition)

	for _, word := range validWords {
		for _, char := range word {
			if char.Char == "A" {
				key := fmt.Sprintf("%d-%d", char.Row, char.Col)
				positionMap[key] = append(positionMap[key], word)
				break
			}
		}
	}

	return positionMap
}

func getWords(
	matrix [][]string,
	si model.SearchInput,
) [][]model.WordPosition {
	var validWords [][]model.WordPosition

	for row := range matrix {
		for col := range matrix[row] {
			for _, dir := range si.Directions {
				ok, validWord := wordSearch(matrix, row, col, si, dir)
				if ok {
					validWords = append(validWords, validWord)
				}
			}
		}
	}

	return validWords
}

func FourthProblem() {
	input := get.StringSliceFromFile("./assets/04-file.txt")

	matrix := createMatrix(input)

	si1 := model.SearchInput{
		Word:       "XMAS",
		Directions: [][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}},
	}

	sigChan, ticker := ui.Setup(10)
	defer ticker.Stop()

	validWords := getWords(matrix, si1)
	count := 0

	mw := mappedWords(validWords)

	fmt.Println(mw)
	fmt.Println("")

	for _, words := range mw {
		for _, word := range words {
			ui.Matrix(matrix, len(si1.Word), word, sigChan, ticker)
			count++
		}
	}

	si2 := model.SearchInput{
		Word:       "MAS",
		Directions: [][]int{{1, 1}, {-1, 1}, {-1, -1}, {1, -1}},
	}

	validWords2 := getWords(matrix, si2)
	count2 := 0

	mw2 := mappedWords(validWords2)
	mw2 = validPairs(mw2)
	if len(validWords2) > 0 {
		count2 = len(mw2)
	}

	for _, words := range mw2 {
		var allWords []model.WordPosition
		for _, word := range words {
			allWords = append(allWords, word...)
			if len(allWords) == len(si2.Word)*2 {
				ui.Matrix(matrix, (len(si2.Word) * 2), allWords, sigChan, ticker)
			}
		}
	}

	fmt.Println("Fourth Problem:", count, count2)
	fmt.Println(ui.ShowCursor)
}
