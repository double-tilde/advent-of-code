package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// GetIntRecords opens the csv file, checks the validity and adds each number
// to the num slice
func GetIntRecords(file string) []int {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading csv", err)
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
