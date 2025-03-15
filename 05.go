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
			if j == i {
				continue
			}

			lhs, rhs := s[i], s[j]
			posL, posR := i, j
			if j < i {
				lhs, rhs = rhs, lhs
				posL, posR = j, i
			}

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

	var selectedPages []model.Pair
	var goodPages [][]int
	var badPages [][]int

	for _, pages := range pagesSet {
		ok := true
		selectedPages = getSlicePairs(pages)

		for i := 0; i < len(selectedPages); i++ {
			pos, vals, ok := compareSlices(
				rulesSet,
				selectedPages[i].Positions,
				selectedPages[i].Values,
			)
			if !ok {
				badPages = append(badPages, pages)
				fmt.Println(pos, vals)
				break
			}
		}

		if ok {
			goodPages = append(goodPages, pages)
		}
	}

	var res int
	for _, pages := range goodPages {
		mp := len(pages) / 2
		res += pages[mp]
	}

	var res2 int
	for _, pages := range badPages {
		mp := len(pages) / 2
		res2 += pages[mp]
	}

	fmt.Println("Fifth Problem:", res, res2)
}
