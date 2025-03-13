package main

import (
	"aoc-2024/get"
	"fmt"
)

func getSlicePairs(s []int) [][]int {
	var sm [][]int
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if j == i {
				continue
			}

			lhs, rhs := s[i], s[j]
			if j < i {
				lhs, rhs = rhs, lhs
			}

			sm = append(sm, []int{lhs, rhs})
		}
	}
	return sm
}

func compareSlices(rulesSet [][]int, pages []int) bool {
	for i := range rulesSet {
		if rulesSet[i][0] == pages[1] && rulesSet[i][1] == pages[0] {
			return false
		}
	}
	return true
}

func FifthProblem() {
	rulesSet := get.IntMatrixPipeDelim("./assets/05-file.txt")
	pagesSet := get.IntMatrixCommaDelim("./assets/05-file.txt")

	var selectedPages [][]int
	var goodPages [][]int
	for _, pages := range pagesSet {
		ok := true
		selectedPages = getSlicePairs(pages)

		for i := 0; i < len(selectedPages); i++ {
			ok = compareSlices(rulesSet, selectedPages[i])
			if !ok {
				break
			}
		}

		if ok {
			goodPages = append(goodPages, pages)
		}
	}

	var res int
	for _, pages := range goodPages {
		mp := 0 + len(pages)/2
		res += pages[mp]
	}

	fmt.Println("Fifth Problem:", res)
}
