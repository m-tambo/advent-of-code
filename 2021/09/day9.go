package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type point struct {
	x int
	y int
}

func main() {
	filePath := os.Args[1]
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	var str string
	var heightMap []string
	var lowPointSum int
	var lowPoints []point
	var basinSizes []int

	for {
		_, err := fmt.Fscanf(f, "%s", &str)
		if err != nil {
			if err == io.EOF {
				break // done reading file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		// tack on padding to each end
		str = "9" + str + "9"
		heightMap = append(heightMap, str)
	}

	// construct a max border and tack onto top and bottom of map
	var border []string
	border = append(border, "")
	for i := 0; i < len(heightMap[0]); i++ {
		border[0] += "9"
	}
	heightMap = append(border, heightMap...)
	heightMap = append(heightMap, border...)

	// start at 1 and end at len-1 to ignore the borders
	for i := 1; i < len(heightMap)-1; i++ {

		for j := 1; j < len(heightMap[i])-1; j++ {

			if isLowPoint(heightMap, i, j) {

				lowPointSum += strToInt(string(heightMap[i][j])) + 1
				lowPoints = append(lowPoints, point{i,j})
			}
		}
	}

	for _, p := range lowPoints {
		basinSizes = append(basinSizes, measureBasinFromPoint(heightMap, p))
	}

	sort.Ints(basinSizes)
	i := len(basinSizes)
	topThreeMultiplied := basinSizes[i-1] * basinSizes[i-2] * basinSizes[i-3]

	fmt.Println("Total height of low points: ", lowPointSum)
	//fmt.Println("List of low points: ", lowPoints)
	fmt.Println("Three largest basins multiplied: ", topThreeMultiplied)
}

func isLowPoint(data []string, idx1, idx2 int) bool {
	num := strToInt(string(data[idx1][idx2]))

	belowVal := strToInt(string(data[idx1+1][idx2]))
	aboveVal := strToInt(string(data[idx1-1][idx2]))
	rightVal := strToInt(string(data[idx1][idx2+1]))
	leftVal := strToInt(string(data[idx1][idx2-1]))

	if num < aboveVal && num < leftVal && num < rightVal && num < belowVal {
		return true
	} else {
		return false
	}
}

func measureBasinFromPoint(data []string, center point) int {
	pts := []point{center}
	pp := &pts

	addAllNeighbors(data, pp, center)
	return len(*pp)
}

func addAllNeighbors(data []string, pointSet *[]point, p point) {
	neighbors := getNeighboringPoints(p)

	// for each neighbor: if not 9 and not already in list, add to pts and check their neighbors
	for _, n := range neighbors {
		if string(data[n.x][n.y]) != "9" && !containsPoint(*pointSet, n) {

			*pointSet = append(*pointSet, n)
			addAllNeighbors(data, pointSet, n)
		}
	}
}

func getNeighboringPoints(p point) []point {
	var n []point
	n = append(n, point{p.x - 1,p.y})
	n = append(n, point{p.x, p.y-1})
	n = append(n, point{p.x,p.y + 1})
	n = append(n, point{p.x + 1,p.y})
	return n
}

func containsPoint(s []point, p point) bool {
	for _, a := range s {
		if a.x == p.x && a.y == p.y{
			return true
		}
	}
	return false
}

func strToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return num
}
