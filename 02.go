package main

import (
	"aoc-2024/utils"
	"fmt"
)

func checkIncrementing(nums []int) (bool, []int) {
	incrementing := true
	var badLevels []int

	if len(nums) <= 1 {
		return incrementing, badLevels
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] || nums[i]-nums[i-1] > 3 {
			incrementing = false
			badLevels = append(badLevels, nums[i-1])
		}
	}

	return incrementing, badLevels
}

func checkDecrementing(nums []int) (bool, []int) {
	decrementing := true
	var badLevels []int

	if len(nums) <= 1 {
		return decrementing, badLevels
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] <= nums[i] || nums[i-1]-nums[i] > 3 {
			decrementing = false
			badLevels = append(badLevels, nums[i-1])
		}
	}

	return decrementing, badLevels
}

func tryReomvals(original, badLevels []int) (bool, int) {
	for i := range badLevels {
		newLevels := []int{}

		for j := range original {
			if original[j] != badLevels[i] {
				newLevels = append(newLevels, original[j])
			}
		}

		safe, _ := checkIncrementing(newLevels)

		if safe {
			return true, badLevels[i]
		}
	}

	return false, 0
}

func tryReomvalsDec(original, badLevels []int) (bool, int) {
	for i := range badLevels {
		newLevels := []int{}

		for j := range original {
			// TODO: This is removing both 4s in 8 6 4 4 1
			if original[j] != badLevels[i] {
				newLevels = append(newLevels, original[j])
			}
		}

		fmt.Println("new levels dec", newLevels)
		safe, _ := checkDecrementing(newLevels)

		if safe {
			return true, badLevels[i]
		}
	}

	return false, 0
}

func loop(list [][]int) int {
	safeStrings := 0

	for _, nums := range list {
		original := []int{}
		original = append(original, nums...)
		safe := false

		fmt.Println("checking levels for", nums)

		incrementing, badLevels := checkIncrementing(nums)

		if incrementing {
			safeStrings++
			fmt.Println("original order safe inc", nums)
		}

		r := 0

		if !incrementing && len(badLevels) > 0 {
			safe, r = tryReomvals(original, badLevels)
		}

		if safe {
			safeStrings++
			fmt.Println("one reomval safe inc", nums, r)
		}

		decrementing, badLevels := checkDecrementing(nums)

		if decrementing {
			safeStrings++
			fmt.Println("original order safe dec", nums)
		}

		if !decrementing && len(badLevels) > 0 {
			safe, r = tryReomvalsDec(original, badLevels)
		}

		if safe {
			safeStrings++
			fmt.Println("one reomval safe dec", nums, r)
		}

	}

	return safeStrings
}

func secondProblem() {
	listOfNums := utils.GetSpaceSeperatedNums("./assets/02-file.txt")

	safe := loop(listOfNums)
	fmt.Println("Problem 2:", safe)
}
