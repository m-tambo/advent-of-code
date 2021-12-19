package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	nums := readDataFromFile()
	data := sumEveryGroupOfN(nums, 3)

	fmt.Println("Part 1 number of increases: ", numberOfIncreases(nums))
	fmt.Println("Part 2 number of increases: ", numberOfIncreases(data))
}

func numberOfIncreases(dataPts []int) int {
	inc := 0
	for i := 1; i < len(dataPts); i++ {
		if dataPts[i] > dataPts[i-1] {
			inc++
		}
	}
	return inc
}


func sumEveryGroupOfN(dataSet []int, n int) []int {
	var newSet []int
	for i := 0; i <= len(dataSet)-n; i++ {
		newSum := 0
		for j := i; j < n+i; j++ {
			newSum += dataSet[j]
		}
		newSet = append(newSet, newSum)
	}
	return newSet
}

func readDataFromFile() []int {
	filePath := os.Args[1]
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	var line int
	var nums []int

	for {
		_, err := fmt.Fscanf(f, "%d\n", &line)
		if err != nil {
			if err == io.EOF {
				break // done reading file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, line)
	}

	return nums
}