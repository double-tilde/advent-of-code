package main

import (
	"aoc-2024/utils"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
)

func compileRegex(p string) *regexp.Regexp {
	r, err := regexp.Compile(p)
	if err != nil {
		log.Fatal("Invalid regex", err)
	}

	return r
}

type match struct {
	text  string
	index int
}

func findValidFuncs(input string) []string {
	regexes := []*regexp.Regexp{
		regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`),
		regexp.MustCompile(`do\(\)`),
		regexp.MustCompile(`don't\(\)`),
	}

	results := make(chan []match, len(regexes))

	for _, re := range regexes {
		go func(re *regexp.Regexp) {
			var matches []match
			for _, loc := range re.FindAllStringIndex(input, -1) {
				matches = append(matches, match{text: input[loc[0]:loc[1]], index: loc[0]})
			}
			results <- matches
		}(re)
	}

	var allMatches []match
	for i := 0; i < len(regexes); i++ {
		allMatches = append(allMatches, <-results...)
	}
	close(results)

	sort.Slice(allMatches, func(i, j int) bool {
		return allMatches[i].index < allMatches[j].index
	})

	var validFuncs []string
	do := true

	for _, match := range allMatches {
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

	validFuncs := findValidFuncs(input)

	numberPairs := getNumberPairs(validFuncs)
	res := calculateProduct(numberPairs)

	fmt.Println("Problem 3:", res)
}
