package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// GetIntRecords opens the csv file, checks the validity and adds each number
// to the num slice
func GetIntRecords(file string) []int {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading file", err)
		return nil
	}

	var nums []int

	for _, row := range records {
		for _, v := range row {
			if v == "" {
				continue
			}
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Error converting to int", err)
				continue
			}
			nums = append(nums, num)
		}
	}

	return nums
}

func GetSpaceSeperatedNums(file string) [][]int {
	var sliceOfNums [][]int

	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var lineNums []int
		fields := strings.Fields(scanner.Text())
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				log.Fatal("String is not an int", err)
			}
			lineNums = append(lineNums, num)
		}
		sliceOfNums = append(sliceOfNums, lineNums)
	}

	return sliceOfNums
}
