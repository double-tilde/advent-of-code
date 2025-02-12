package main

import (
	"aoc-2024/utils"
	"fmt"
)

// sort the slices into ascending order
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

// put one of the slices into a map
func createMap(list []int) map[int]int {
	m := make(map[int]int)

	for i, v := range list {
		m[i] = v
	}

	return m
}

// compare the other slice to the map
func compare(m map[int]int, s []int) int {
	sum := 0
	for k, v := range m {
		diff := 0

		if s[k] == v {
			diff = 0
		}

		if s[k] > v {
			diff = s[k] - v
		}

		if s[k] < v {
			diff = v - s[k]
		}

		sum = sum + diff
	}

	return sum
}

func firstProblem() {
	// Get the lists
	l1 := utils.GetIntRecords("./assets/01-list1.csv")
	l2 := utils.GetIntRecords("./assets/01-list2.csv")

	sortList(l1)
	sortList(l2)

	m := createMap(l1)

	res := compare(m, l2)

	fmt.Println(res)
}
