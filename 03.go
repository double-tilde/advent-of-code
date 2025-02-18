package main

import (
	"aoc-2024/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func compileRegex(p string) *regexp.Regexp {
	r, err := regexp.Compile(p)
	if err != nil {
		log.Fatal("Invalid regex", err)
	}

	return r
}

func findMulFuncs(input string) []string {
	var mulFuncs []string

	r := compileRegex(`mul\((\d{1,3}),(\d{1,3})\)`)

	mulFuncs = r.FindAllString(input[:], -1)

	return mulFuncs
}

func getNumberPairs(strs []string) [][]int {
	var validNumPairs []int
	var numPairs [][]int

	for _, str := range strs {
		r := compileRegex(`(\d{1,3}),(\d{1,3})`)

		// Extract matched number pairs
		matches := r.FindAllStringSubmatch(str, -1)

		for _, match := range matches {
			if len(match) == 3 {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])

				validNumPairs = []int{num1, num2}
				numPairs = append(numPairs, validNumPairs)
			}
		}

	}

	return numPairs
}

func calculateProduct(numPairs [][]int) int {
	res := 0

	for _, pair := range numPairs {
		for i := range pair {
			if i == 0 {
				res += pair[i] * pair[i+1]
			}
		}
	}
	return res
}

func thirdProblem() {
	input := utils.GetOneString("./assets/03-file.txt")

	mulFuncs := findMulFuncs(input)
	numberPairs := getNumberPairs(mulFuncs)
	res := calculateProduct(numberPairs)

	fmt.Println("Problem 3:", res)
}
