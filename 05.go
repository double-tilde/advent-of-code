package main

import (
	"aoc-2024/get"
	"aoc-2024/model"
	"fmt"
)

func getSlicePairs(s []int) []model.Pair {
	var pairs []model.Pair

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if j <= i {
				continue
			}

			lhs, rhs := s[i], s[j]
			posL, posR := i, j

			pairs = append(pairs, model.Pair{Positions: []int{posL, posR}, Values: []int{lhs, rhs}})
		}
	}
	return pairs
}

func compareSlices(rulesSet [][]int, pos, vals []int) ([]int, []int, bool) {
	for i := range rulesSet {
		if rulesSet[i][0] == vals[1] && rulesSet[i][1] == vals[0] {
			return pos, vals, false
		}
	}
	return nil, nil, true
}

func FifthProblem() {
	rulesSet := get.IntMatrixPipeDelim("./assets/05-file.txt")
	pagesSet := get.IntMatrixCommaDelim("./assets/05-file.txt")

	var correctPages [][]int
	var incorrectPages [][]int

	for _, pages := range pagesSet {

		selectedPages := getSlicePairs(pages)

		correct := true
		for i := 0; i < len(selectedPages); i++ {
			_, _, correct = compareSlices(
				rulesSet,
				selectedPages[i].Positions,
				selectedPages[i].Values,
			)
			if !correct {
				break
			}
		}

		if correct {
			correctPages = append(correctPages, pages)
		}

		if !correct {
			incorrectPages = append(incorrectPages, pages)
		}

	}

	fmt.Println(correctPages)
	fmt.Println(incorrectPages)

	var correctedPages [][]int
	// TODO: Makes this part recursive

	for _, pages := range incorrectPages {
		selectedPages := getSlicePairs(pages)

		correct := true
		vals := []int{}
		for i := 0; i < len(selectedPages); i++ {
			_, vals, correct = compareSlices(
				rulesSet,
				selectedPages[i].Positions,
				selectedPages[i].Values,
			)

			if !correct {
				for j := range pages {
					if pages[j] == vals[0] {
						pages[j] = vals[1]
					}

					if pages[j] == vals[1] {
						pages[j] = vals[0]
					}
				}
				break
			}
		}

	}

	var res int
	for _, pages := range correctPages {
		mp := len(pages) / 2
		res += pages[mp]
	}

	fmt.Println("Fifth Problem:", res)
}
