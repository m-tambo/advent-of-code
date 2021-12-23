package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	// use args to include diagonal lines
	includeDiagonals := false
	if len(os.Args) > 2 && os.Args[2] == "diagonal" {
		includeDiagonals = true
	}

	filePath := os.Args[1]
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	// record each point as a string formatted "x,y", with the int being its number of occurrences
	pointMap := map[string]int{}
	var x1, x2, y1, y2 int

	for {
		_, err := fmt.Fscanf(f, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			if err == io.EOF {
				break // done reading file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		// determine 45 degree lines using rise/run
		is45Degree := math.Abs(float64(x2) - float64(x1)) == math.Abs(float64(y2) - float64(y1))

		if x1 == x2 {
			if y1 > y2 {
				for i := 0; i <= y1-y2; i++ {
					s := fmt.Sprintf("%d,%d", x1, y1-i)
					insertOrIncrement(pointMap, s)
				}
			} else {
				for i := 0; i <= y2-y1; i++ {
					s := fmt.Sprintf("%d,%d", x1, y1+i)
					insertOrIncrement(pointMap, s)
				}
			}
		} else

		if y1 == y2 {
			if x1 > x2 {
				for i := 0; i <= x1-x2; i++ {
					s := fmt.Sprintf("%d,%d", x1-i, y1)
					insertOrIncrement(pointMap, s)
				}
			} else {
				for i := 0; i <= x2-x1; i++ {
					s := fmt.Sprintf("%d,%d", x1+i, y1)
					insertOrIncrement(pointMap, s)
				}
			}
		} else

		if includeDiagonals && is45Degree {
			if x1 > x2 && y1 > y2 {
				for i := 0; i <= x1-x2; i++ {
					s := fmt.Sprintf("%d,%d", x1-i, y1-i)
					insertOrIncrement(pointMap, s)
				}
			}

			if x1 > x2 && y2 > y1 {
				for i := 0; i <= x1-x2; i++ {
					s := fmt.Sprintf("%d,%d", x1-i, y1+i)
					insertOrIncrement(pointMap, s)
				}
			}

			if x2 > x1 && y1 > y2 {
				for i := 0; i <= x2-x1; i++ {
					s := fmt.Sprintf("%d,%d", x1+i, y1-i)
					insertOrIncrement(pointMap, s)
				}
			}

			if x2 > x1 && y2 > y1 {
				for i := 0; i <= x2-x1; i++ {
					s := fmt.Sprintf("%d,%d", x1+i, y1+i)
					insertOrIncrement(pointMap, s)
				}
			}
		}
	}

	fmt.Println("Total: ", numOfEntriesOccurringMoreThanNTimes(pointMap, 1))
}

func insertOrIncrement(m map[string]int, str string) map[string]int {
	if _, ok := m[str]; ok {
		m[str]++
	} else {
		m[str] = 1
	}

	return m
}

func numOfEntriesOccurringMoreThanNTimes(m map[string]int, n int) int {
	count := 0
	for _, v := range m {
		if v > n {
			count++
		}
	}
	return count
}