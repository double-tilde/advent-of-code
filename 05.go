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

func checkPages(rulesSet [][]int, pages []int) ([]int, bool) {
	var ok bool
	selectedPages := getSlicePairs(pages)

	// TODO: fix the copying problem, night need pointers or make sure model.pair is correct

	for i := 0; i < len(selectedPages); i++ {
		pos, vals, ok := compareSlices(
			rulesSet,
			selectedPages[i].Positions,
			selectedPages[i].Values,
		)

		if !ok {
			pages[pos[0]] = vals[1]
			pages[pos[1]] = vals[0]
		}
	}

	return pages, ok
}

func FifthProblem() {
	rulesSet := get.IntMatrixPipeDelim("./assets/05-file.txt")
	pagesSet := get.IntMatrixCommaDelim("./assets/05-file.txt")

	var goodPages [][]int
	var incorrectPages [][]int

	for _, pages := range pagesSet {
		ok, incorrect := true, false
		selectedPages := getSlicePairs(pages)

		for i := 0; i < len(selectedPages); i++ {
			_, _, ok := compareSlices(
				rulesSet,
				selectedPages[i].Positions,
				selectedPages[i].Values,
			)
			if !ok {
				incorrect = true
				break
			}
		}

		if ok && !incorrect {
			goodPages = append(goodPages, pages)
		}

		if ok && incorrect {
			incorrectPages = append(incorrectPages, pages)
		}
	}

	fmt.Println(goodPages)
	fmt.Println(incorrectPages)

	var correctedPages [][]int

	for _, pages := range incorrectPages {
		p, ok := checkPages(rulesSet, pages)

		if !ok {
			checkPages(rulesSet, p)
		}
		correctedPages = append(correctedPages, p)
	}

	fmt.Println(correctedPages)

	var res int
	for _, pages := range goodPages {
		mp := len(pages) / 2
		res += pages[mp]
	}

	var res2 int
	for _, pages := range correctedPages {
		mp := len(pages) / 2
		res2 += pages[mp]
	}

	fmt.Println("Fifth Problem:", res, res2)
}
