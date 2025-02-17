package main

import (
	"aoc-2024/utils"
	"fmt"
)

func sortList(list []int) []int {
	// selection sort
	for i := 0; i < len(list)-1; i++ {
		mIdx := i
		for j := i + 1; j <= len(list)-1; j++ {
			if list[j] < list[mIdx] {
				mIdx = j
			}
		}

		list[i], list[mIdx] = list[mIdx], list[i]
	}
	return list
}

func createMap(list []int) map[int]int {
	m := make(map[int]int)

	for i, v := range list {
		m[i] = v
	}

	return m
}

// See how many times a value in the left list appears in the right list
// Multiply that value by amount of appearances then add up all of those values
func compare(m map[int]int, s []int) int {
	sum := 0
	for i := range s {
		appears := 0
		for _, v := range m {
			if s[i] == v {
				appears++
			}
		}
		appears = s[i] * appears
		sum += appears
	}

	return sum
}

func firstProblem() {
	l1 := utils.GetIntRecords("./assets/01-list1.csv")
	l2 := utils.GetIntRecords("./assets/01-list2.csv")

	sortList(l1)
	sortList(l2)

	// For part 2 of the question, list 2 needs to be in the map
	m := createMap(l2)

	res := compare(m, l1)

	fmt.Println("Problem 1:", res)
}
