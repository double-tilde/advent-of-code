package main

import (
	"aoc-2024/utils"
	"fmt"
)

func checkIncrementing(nums, original []int) bool {
	order := true
	var badLevels []int

	if len(nums) <= 1 {
		return order
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] || nums[i]-nums[i-1] > 3 {
			order = false
			badLevels = append(badLevels, i-1)
		}
	}

	if len(badLevels) > 0 {
		for i := range badLevels {
			newLevels := []int{}

			for j := range original {
				if original[j] != badLevels[i] {
					newLevels = append(newLevels, original[j])
				}
			}

			fmt.Println("Original:", original)
			fmt.Println("New Levels:", newLevels)

			if checkIncrementing(newLevels, original) {
				return true
			}
		}
	}

	return order
}

func checkDecrementing(nums, original []int) bool {
	order := true
	var badLevels []int

	if len(nums) <= 1 {
		return order
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] <= nums[i] || nums[i-1]-nums[i] > 3 {
			order = false
			badLevels = append(badLevels, i-1)
		}
	}

	if len(badLevels) > 0 {
		for i := range badLevels {
			newLevels := []int{}

			for j := range original {
				if original[j] != badLevels[i] {
					newLevels = append(newLevels, original[j])
				}
			}

			fmt.Println("Original:", original)
			fmt.Println("New Levels:", newLevels)

			if checkDecrementing(newLevels, original) {
				return true
			}
		}
	}

	return order
}

func loop(list [][]int) int {
	safeStrings := 0

	for _, nums := range list {
		original := []int{}
		original = append(original, nums...)

		if checkIncrementing(nums, original) || checkDecrementing(nums, original) {
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
