package main

import (
	"aoc-2024/get"
	"fmt"
)

func getSlicePairs(s []int) [][]int {
	var sm [][]int
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			var sp []int
			if j == i {
				continue
			}
			sp = append(sp, s[i], s[j])
			sm = append(sm, sp)
		}
	}
	return sm
}

func compareSlices(rulesSet [][]int, pages []int) bool {
	for i := range rulesSet {
		fmt.Println("comparing", rulesSet[i][0], rulesSet[i][1], "with", pages[0], pages[1])
		if rulesSet[i][0] == pages[1] && rulesSet[i][1] == pages[0] {
			fmt.Println("wrong order", rulesSet[i][0], rulesSet[i][1], "see", pages[0], pages[1])
			return false
		}
		fmt.Println("right order", rulesSet[i][0], rulesSet[i][1], "see", pages[0], pages[1])
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
		fmt.Println(goodPages)
	}

	// 1. For the rules, create a [][]int with the first num as [i][0] and the second as [i][1]
	//    The pipe can be used as the deliminator
	// 2. For the page numbers, also create a [][]int, with each list as [i][i]
	// 3. Starting with the first set of page numbers, start with [0][0] and [0][1] and put them in a []int
	//    Check the positions. If they are in the right order against the rules, they should appear in the
	//    same positions that they are in currently in their comparrison slice
	// 4. If they are not matching, return early as the page numbers are in the wrong order, if they are,
	//    then move to compare [0][0] and [0][2], and so on
	// 5. If the slices are correct, append them to a [][]int of correct slices
	// 6. For each of the correct slices, find the int in the middle position and add them up

	fmt.Println("Fifth Problem:")
}
