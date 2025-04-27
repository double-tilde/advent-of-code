package main

import (
	"aoc-2024/get"
	"aoc-2024/model"
	"fmt"
)

// TODO: Make faster, see note

func getSlicePairs(s []int) []model.Pair {
	var pairs []model.Pair

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			lhs, rhs := s[i], s[j]
			posL, posR := i, j

			pairs = append(pairs, model.Pair{Positions: []int{posL, posR}, Values: []int{lhs, rhs}})
		}
	}
	return pairs
}

func compareSlices(rulesSet [][]int, vals []int) bool {
	for i := range rulesSet {
		if rulesSet[i][0] == vals[1] && rulesSet[i][1] == vals[0] {
			return false
		}
	}
	return true
}

// NOTE: use go routines and channels, see my sandbox for how to do it
func checkPages(rulesSet, startingPages [][]int, swap bool) ([][]int, [][]int) {
	correctPages, incorrectPages := [][]int{}, [][]int{}

	for _, pages := range startingPages {

		selectedPairs := getSlicePairs(pages)

		correct := true
		for _, pairs := range selectedPairs {
			if !compareSlices(rulesSet, pairs.Values) {
				if swap {
					pages[pairs.Positions[0]], pages[pairs.Positions[1]] = pages[pairs.Positions[1]], pages[pairs.Positions[0]]
				}
				correct = false
				break
			}
		}

		if !correct {
			incorrectPages = append(incorrectPages, pages)
		}

		if correct {
			correctPages = append(correctPages, pages)
		}

	}

	return correctPages, incorrectPages
}

func addUpMedian(intMatrix [][]int) int {
	var res int
	for _, ints := range intMatrix {
		mp := len(ints) / 2
		res += ints[mp]
	}

	return res
}

func FifthProblem() {
	rulesSet := get.IntMatrixPipeDelim("./assets/05-file.txt")
	pagesSet := get.IntMatrixCommaDelim("./assets/05-file.txt")

	correctPages, incorrectPages := checkPages(rulesSet, pagesSet, false)

	var correctedPages [][]int
	for {
		tempCorrectedPages, stillIncorrectPages := checkPages(rulesSet, incorrectPages, true)

		correctedPages = append(correctedPages, tempCorrectedPages...)

		if len(stillIncorrectPages) == 0 {
			break
		}

		incorrectPages = stillIncorrectPages
	}

	res := addUpMedian(correctPages)
	res2 := addUpMedian(correctedPages)

	fmt.Println("Fifth Problem:", res, res2)
}
