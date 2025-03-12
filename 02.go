package main

import (
	"aoc-2024/get"
	"fmt"
)

func isIncrementing(level []int) (bool, map[int]int) {
	isInc := true
	bl := make(map[int]int)

	if len(level) <= 1 {
		return isInc, bl
	}

	for i := 1; i < len(level); i++ {
		if level[i-1] >= level[i] || level[i]-level[i-1] > 3 {
			isInc = false
			bl[i-1] = level[i-1]
			bl[i] = level[i]
		}
	}

	return isInc, bl
}

func isDecrementing(level []int) (bool, map[int]int) {
	isDec := true
	bl := make(map[int]int)

	if len(level) <= 1 {
		return isDec, bl
	}

	for i := 1; i < len(level); i++ {
		if level[i-1] <= level[i] || level[i-1]-level[i] > 3 {
			isDec = false
			bl[i-1] = level[i-1]
			bl[i] = level[i]
		}
	}

	return isDec, bl
}

func tryRemovals(orig []int, bl map[int]int, try func([]int) (bool, map[int]int)) bool {
	for pos := range bl {
		newLevels := []int{}

		for p, v := range orig {
			// Create a new level missing one of the bad levels
			if p != pos {
				newLevels = append(newLevels, v)
			}
		}

		ok, _ := try(newLevels)

		if ok {
			return true
		}
	}

	return false
}

func loop(sliceOfLevels [][]int) int {
	safeLevels := 0

	for _, level := range sliceOfLevels {
		original := []int{}
		original = append(original, level...)
		ok := false

		// Check level is incrementing
		incrementing, badLevels := isIncrementing(level)
		if incrementing {
			safeLevels++
		}

		// If not, try removing one bad level
		if !incrementing && len(badLevels) > 0 {
			ok = tryRemovals(original, badLevels, isIncrementing)
		}
		if ok {
			safeLevels++
		}

		// Check level is decrementing
		decrementing, badLevels := isDecrementing(level)
		if decrementing {
			safeLevels++
		}

		// If not, try removing one bad level
		if !decrementing && len(badLevels) > 0 {
			ok = tryRemovals(original, badLevels, isDecrementing)
		}
		if ok {
			safeLevels++
		}

	}

	return safeLevels
}

func SecondProblem() {
	sliceOfLevels := get.IntMatrixFromFile("./assets/02-file.txt")

	safe := loop(sliceOfLevels)
	fmt.Println("Problem 2:", safe)
}
