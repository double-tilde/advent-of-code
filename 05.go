package main

import "fmt"

func FifthProblem() {
	input := utils.GetDelimSeperated
	// 47|53
	// 97|13
	// 97|61
	// 97|47
	// 75|29
	// 61|13
	// 75|53
	// 29|13
	// 97|29
	// 53|29
	// 61|53
	// 97|53
	// 61|29
	// 47|13
	// 75|47
	// 97|75
	// 47|61
	// 75|61
	// 47|29
	// 75|13
	// 53|13

	// 75,47,61,53,29 -- correct
	// 97,61,53,29,13 -- correct
	// 75,29,13       -- correct
	// 75,97,47,61,53 -- incorrect
	// 61,13,29       -- incorrect
	// 97,13,75,29,47 -- incorrect

	// 1. For the rules, create a [][]int with the first num as [i][0] and the second as [i][1]
	//    The pipe can be used as the deliminator
	// 2. For the page numbers, also create a [][]int, with each list as [i][i]
	// 3. Starting with the first set of page numbers, start with [0][0] and [0][1] and put them in a []int
	//    Check the positions. If they are in the right order against the rules, they should appear in the
	//    same positions that they are in currently in their comparrison slice
	// 4. If they are not matching, return early as the page numbers are in the wrong order, if they are,
	//    then move to compare [0][0] and [0][2], and so on
	// 5. If the slices are correct, append them to a [][]int of correct slices
	// 6. For each of the correct slices, find the int in the middle position and add them up

	fmt.Println("Fifth Problem:")
}
