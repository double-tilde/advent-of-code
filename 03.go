package main

import (
	"aoc-2024/get"
	"cmp"
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

var (
	mulRegex  = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRegex   = regexp.MustCompile(`do\(\)`)
	dontRegex = regexp.MustCompile(`don't\(\)`)
)

type match struct {
	text  string
	index int
}

func findValidFuncs(input string) []string {
	regexes := []*regexp.Regexp{mulRegex, doRegex, dontRegex}
	var matches []match

	for _, re := range regexes {
		for _, loc := range re.FindAllStringIndex(input, -1) {
			matches = append(matches, match{text: input[loc[0]:loc[1]], index: loc[0]})
		}
	}

	slices.SortFunc(matches, func(a, b match) int {
		return cmp.Compare(a.index, b.index)
	})

	var validFuncs []string
	do := true

	for _, match := range matches {
		if match.text == "do()" {
			do = true
		}
		if match.text == "don't()" {
			do = false
		}

		if do {
			validFuncs = append(validFuncs, match.text)
		}

	}

	return validFuncs
}

func getNumberPairs(strs []string) [][]int {
	var numPairs [][]int

	for _, str := range strs {

		matches := mulRegex.FindStringSubmatch(str)

		if len(matches) == 3 {
			num1, _ := strconv.Atoi(matches[1])
			num2, _ := strconv.Atoi(matches[2])

			numPairs = append(numPairs, []int{num1, num2})
		}

	}

	return numPairs
}

func calculateProduct(numPairs [][]int) int {
	res := 0

	for _, pair := range numPairs {
		res += pair[0] * pair[1]
	}
	return res
}

func ThirdProblem() {
	input := get.StringFromFile("./assets/03-file.txt")

	validFuncs := findValidFuncs(input)

	numberPairs := getNumberPairs(validFuncs)
	res := calculateProduct(numberPairs)

	fmt.Println("Problem 3:", res)
}
