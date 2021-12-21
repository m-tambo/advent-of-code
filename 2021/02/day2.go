package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// add "aim" to args for second part of the puzzle
	var useAim bool
	if len(os.Args) > 2 && os.Args[2] == "aim" {
		useAim = true
	}

	filePath := os.Args[1]
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	var hPos int
	var vPos int
	var aim int
	var dir string
	var dist int

	for {
		_, err := fmt.Fscanf(f, "%s %d\n", &dir, &dist)
		if err != nil {
			if err == io.EOF {
				break // done reading file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		if useAim {

			hPos, vPos, aim = navigateWithAim(dir, dist, hPos, vPos, aim)
		} else {

			hPos, vPos = navigate(dir, dist, hPos, vPos)
		}
	}

	fmt.Println("Total: ", hPos * vPos)
}

func navigate(direction string, distance int, horizontalPos int, verticalPos int) (int, int) {
	switch direction {
	case "forward":
		horizontalPos += distance
	case "down":
		verticalPos += distance
	case "up":
		verticalPos -= distance
	}

	return horizontalPos, verticalPos
}

func navigateWithAim(direction string, distance int, horizontalPos int, verticalPos int, aim int) (int, int, int) {
	switch direction {
	case "forward":
		horizontalPos += distance
		verticalPos = verticalPos + distance * aim
	case "down":
		aim += distance
	case "up":
		aim -= distance
	}

	return horizontalPos, verticalPos, aim
}