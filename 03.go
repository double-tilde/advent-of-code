package main

import (
	"aoc-2024/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func getValidFuncs(str string) []string {
	var validFuncs []string

	pattern := `mul\((\d{1,3}),(\d{1,3})\)`

	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal("Invalid regex", err)
	}

	validFuncs = r.FindAllString(str[:], -1)

	return validFuncs
}

func getValidNumsSet(strs []string) [][]int {
	var validNumsSet [][]int
	var validNums []int

	for _, str := range strs {
		pattern := `(\d{1,3}),(\d{1,3})`

		r, err := regexp.Compile(pattern)
		if err != nil {
			log.Fatal("Invalid regex", err)
		}

		// Extract matched number pairs
		matches := r.FindAllStringSubmatch(str, -1)

		for _, match := range matches {
			if len(match) == 3 {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])

				validNums = []int{num1, num2}
				validNumsSet = append(validNumsSet, validNums)
			}
		}

	}

	return validNumsSet
}

func calculate(numsSet [][]int) int {
	answer := 0

	for _, nums := range numsSet {
		for i := range nums {
			if i == 0 {
				answer += nums[i] * nums[i+1]
			}
		}
	}
	return answer
}

func thirdProblem() {
	str := utils.GetOneString("./assets/03-file.txt")

	validFuncs := getValidFuncs(str)
	validNumsSet := getValidNumsSet(validFuncs)
	answer := calculate(validNumsSet)

	fmt.Println("Problem 3:", answer)
}
