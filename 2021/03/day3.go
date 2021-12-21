package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fullData, bitCounters := readAndCountData()

	gamma := convertIntsToBinaryStr(bitCounters)
	epsilon := generateInvertedBinaryString(gamma)

	moreCommon := make([]string, len(fullData))
	_ = copy(moreCommon, fullData)
	lesCommon := make([]string, len(fullData))
	_ = copy(lesCommon, fullData)

	for i := 0; i < len(gamma); i++ {
		if len(moreCommon) == 1 {
			break
		}
		b := determineMoreCommonBitAtIdx(moreCommon, i)
		moreCommon = filterOutByIdxMatch(moreCommon, i, b)
	}

	for i := 0; i < len(epsilon); i++ {
		if len(lesCommon) == 1 {
			break
		}
		b := determineLessCommonBitAtIdx(lesCommon, i)
		lesCommon = filterOutByIdxMatch(lesCommon, i, b)
	}

	fmt.Println("Part 1 (power consumption): ", convertBinaryStrToInt(gamma) * convertBinaryStrToInt(epsilon))
	fmt.Println("Part 2 (rating): ", convertBinaryStrToInt(moreCommon[0]) * convertBinaryStrToInt(lesCommon[0]))

	// tests
	//fmt.Println("10010 should be ", convertIntsToBinaryStr([]int{7,-5,-4,1,-2}))
	//fmt.Println("10010 should be ", generateInvertedBinaryString("01101"))
	//fmt.Println("13 should be ", convertBinaryStrToInt("01101"))
	//fmt.Println("[11101] should be ", filterOutByIdxMatch([]string{"10010","10010","11101"}, 4,"1"))
	//fmt.Println("1 should be ", determineMoreCommonBitAtIdx([]string{"10010","10010","11101","11110"}, 3))
	//fmt.Println("1 should be ", determineMoreCommonBitAtIdx([]string{"10010","10010","11101","11110"}, 1))
	//fmt.Println("1 should be ", determineLessCommonBitAtIdx([]string{"10010","10010","11101","11110"}, 4))
	//fmt.Println("0 should be ", determineLessCommonBitAtIdx([]string{"10010","10010","11101","11110"}, 1))
}

func convertIntsToBinaryStr(counters []int) string {
	binaryStr := ""

	for _, p := range counters {
		if p > 0 {
			binaryStr += "1"
		} else
		if p < 0 {
			binaryStr += "0"
		} else {
			println("what happened, this shouldn't be zero")
		}
	}

	return binaryStr
}

func convertBinaryStrToInt(binaryStr string) int64 {
	n, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return n
}

func generateInvertedBinaryString(input string) string {
	output := ""
	for _, c := range input {
		if string(c) == "1" {
			output += "0"
		}
		if string(c) == "0" {
			output += "1"
		}
	}
	return output
}

func filterOutByIdxMatch(data []string, idx int, comparator string) []string {
	var filtered []string
	for _, el := range data {
		if string(el[idx]) == comparator {
			filtered = append(filtered, el)
		}
	}
	return filtered
}

// if equal, return "1"
func determineMoreCommonBitAtIdx(data []string, idx int) string {
	counter := 0
	for _, el := range data {
		if string(el[idx]) == "1" {
			counter++
		}
		if string(el[idx]) == "0" {
			counter--
		}
	}
	if counter >= 0 {
		return "1"
	} else {
		return "0"
	}
}

// if equal, return "0"
func determineLessCommonBitAtIdx(data []string, idx int) string {
	counter := 0
	for _, el := range data {
		if string(el[idx]) == "1" {
			counter++
		}
		if string(el[idx]) == "0" {
			counter--
		}
	}
	if counter < 0 {
		return "1"
	} else {
		return "0"
	}
}

// read data from file and write to slice of strings while also keeping count of bits at each position
func readAndCountData() ([]string, []int) {
	filePath := os.Args[1]
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var fullData []string
	var bitCounters []int

	for {
		var line string

		_, err := fmt.Fscanf(f, "%s\n", &line)
		if err != nil {
			if err == io.EOF {
				break // done reading file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		// initialize counter slice first time through
		if len(bitCounters) == 0 {
			for len(bitCounters) < len(line) {
				bitCounters = append(bitCounters, 0)
			}
		}

		// tally the bit (1 or 0) at each position
		for i, d := range line {
			if string(d) == "1" {
				bitCounters[i]++
			}
			if string(d) == "0" {
				bitCounters[i]--
			}
		}

		// also build full data set
		fullData = append(fullData, line)
	}

	return fullData, bitCounters
}