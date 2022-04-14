package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filePath := os.Args[1]
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	var str string
	var crabs []int

	for {
		_, err := fmt.Fscanf(f, "%s", &str)
		if err != nil {
			if err == io.EOF {
				break // done reading file
			}
			fmt.Println(err)
			os.Exit(1)
		}
	}

	strSlice := strings.Split(str, ",")
	for _, v := range strSlice {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		crabs = append(crabs, num)
	}

	fmt.Println("Total dist to mean: ", totalDistanceMoved(crabs, findMean(crabs), false))
	fmt.Println("Total dist to avg: ", totalDistanceMoved(crabs, findAvg(crabs), true))
}

func findMean(s []int) int {
	var mean int

	sort.Ints(s)

	if len(s) % 2 != 0 {
		idx := int(math.Ceil(float64(len(s) / 2)))
		mean = s[idx]
	} else {
		idx := len(s) / 2
		mean = int(math.Round(float64(s[idx] + s[idx-1]) / 2))
	}

	return mean
}

func findAvg(s []int) int {
	var sum int

	for _, v := range s {
		sum += v
	}

	return int(math.Floor(float64(sum) / float64(len(s))))
}

func totalDistanceMoved(points []int, dest int, multiplier bool) int {
	var dist int

	for _, p := range points {
		if !multiplier {
			dist += int(math.Abs(float64(dest - p)))
		} else {
			x := int(math.Abs(float64(dest) - float64(p)))

			for i := 0; i <= x; i++ {
				dist += i
			}
		}
	}

	return dist
}
