package main

import (
	"aoc-2024/utils"
	"fmt"
)

func checkIncrementing(nums []int) (bool, map[int]int) {
	inc := true
	bl := make(map[int]int)

	if len(nums) <= 1 {
		return inc, bl
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] || nums[i]-nums[i-1] > 3 {
			inc = false
			bl[i-1] = nums[i-1]
			bl[i] = nums[i]
		}
	}

	return inc, bl
}

func checkDecrementing(nums []int) (bool, map[int]int) {
	dec := true
	bl := make(map[int]int)

	if len(nums) <= 1 {
		return dec, bl
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] <= nums[i] || nums[i-1]-nums[i] > 3 {
			dec = false
			bl[i-1] = nums[i-1]
			bl[i] = nums[i]
		}
	}

	return dec, bl
}

func tryReomvals(orig []int, bl map[int]int, try func([]int) (bool, map[int]int)) bool {
	for pos := range bl {
		nl := []int{}

		for p, v := range orig {
			// Create a new levels missing one of the bad levels
			if p != pos {
				nl = append(nl, v)
			}
		}

		safe, _ := try(nl)

		if safe {
			return true
		}
	}

	return false
}

func loop(list [][]int) int {
	safeStrings := 0

	for _, nums := range list {
		original := []int{}
		original = append(original, nums...)
		safe := false

		// See if levels are incrementing
		incrementing, badLevels := checkIncrementing(nums)
		if incrementing {
			safeStrings++
		}

		// See if removing one level makes levels increment
		if !incrementing && len(badLevels) > 0 {
			safe = tryReomvals(original, badLevels, checkIncrementing)
		}
		if safe {
			safeStrings++
		}

		// See if levels are decrementing
		decrementing, badLevels := checkDecrementing(nums)
		if decrementing {
			safeStrings++
		}

		// See if removing one level makes levels decrement
		if !decrementing && len(badLevels) > 0 {
			safe = tryReomvals(original, badLevels, checkDecrementing)
		}
		if safe {
			safeStrings++
		}

	}

	return safeStrings
}

func secondProblem() {
	listOfNums := utils.GetSpaceSeperatedNums("./assets/02-file.txt")

	safe := loop(listOfNums)
	fmt.Println("Problem 2:", safe)
}
