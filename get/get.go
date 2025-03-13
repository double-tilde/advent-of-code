package get

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func IntSliceFromCSV(file string) []int {
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

func IntMatrixFromFile(file string) [][]int {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var intMatrix [][]int

	for scanner.Scan() {
		var intSlice []int
		fields := strings.Fields(scanner.Text())

		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				log.Fatal("String is not an int", err)
			}
			intSlice = append(intSlice, num)
		}
		intMatrix = append(intMatrix, intSlice)
	}

	return intMatrix
}

func StringFromFile(file string) string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var s string
	for scanner.Scan() {
		strs := strings.Fields(scanner.Text())

		for _, str := range strs {
			s += str
		}
	}

	return s
}

func StringSliceFromFile(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	records := []string{}
	for scanner.Scan() {
		strs := scanner.Text()
		records = append(records, strs)
	}

	return records
}

func IntMatrixPipeDelim(file string) [][]int {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Error opeing file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var intMatrix [][]int
	for scanner.Scan() {
		var intSlice []int
		str := scanner.Text()
		if str == "" {
			break
		}
		strSlice := strings.Split(str, "|")

		for _, item := range strSlice {
			num, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal("String is not an int", err)
			}
			intSlice = append(intSlice, num)
		}
		intMatrix = append(intMatrix, intSlice)

	}

	return intMatrix
}

func IntMatrixCommaDelim(file string) [][]int {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Error opeing file", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var intMatrix [][]int
	for scanner.Scan() {
		var intSlice []int
		str := scanner.Text()
		if str == "" || strings.Contains(str, "|") {
			continue
		}

		strSlice := strings.Split(str, ",")

		for _, item := range strSlice {
			num, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal("String is not an int", err)
			}
			intSlice = append(intSlice, num)
		}
		intMatrix = append(intMatrix, intSlice)

	}

	return intMatrix
}
