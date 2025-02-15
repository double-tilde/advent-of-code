package main

import (
	"aoc-2024/utils"
	"fmt"
)

// TODO: I need a way to make comparrisons to beyond just i-1, use a map
// perhaps you could do a sort and if nothing changes that means things are okay?
// or a different way perhaps

func checkIncrementing(nums []int) bool {
	order := true
	for i := range nums {
		if i > 0 && nums[i-1] >= nums[i] {
			return order == false
		}
		if i > 0 {
			diff := nums[i] - nums[i-1]
			if diff > 3 {
				return order == false
			}
		}
	}
	return order
}

func checkDecrementing(nums []int) bool {
	order := true
	for i := range nums {
		if i > 0 && nums[i-1] <= nums[i] {
			return order == false
		}
		if i > 0 {
			diff := nums[i-1] - nums[i]
			if diff > 3 {
				return order == false
			}
		}
	}
	return order
}

func loop(list [][]int) int {
	safeStrings := 0

	for i := range list[:] {
		if checkIncrementing(list[i]) {
			safeStrings++
		}
		if checkDecrementing(list[i]) {
			safeStrings++
		}
	}

	return safeStrings
}

func seconfProblem() {
	listOfNums := utils.GetSpaceSeperatedNums("./assets/02-file.txt")

	safe := loop(listOfNums)
	fmt.Println("Problem 2:", safe)
}
