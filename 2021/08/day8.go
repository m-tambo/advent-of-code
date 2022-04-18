package main

import (
	"fmt"
	"io"
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

	codeMap := map[string]int{"abcefg":0, "cf":1, "acdeg":2, "acdfg":3, "bcdf":4, "abdfg":5, "abdefg":6, "acf":7, "abcdefg":8, "abcdfg":9}
	var str string
	var strSlice []string
	toggle := false
	lastFour := 0
	count := 0

	for {
		_, err := fmt.Fscanf(f, "%s", &str)
		if err != nil {
			if err == io.EOF {
				break // done reading file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		// don't add the pipe to the strSlice
		if str == "|" {
			toggle = true
		} else {
			strSlice = append(strSlice, str)
		}

		// start counting to four at the pipe, ignoring the pipe itself
		if toggle && str != "|" {
			lastFour++
		}

		// evaluate once the end of the line is reached
		if lastFour == 4 {
			stringToAdd := ""
			decoderMap := createReverseMap(getDecoder(strSlice))

			// add values from last four of the strSlice
			for i := len(strSlice)-1; i > len(strSlice)-5; i-- {
				decodedStr := ""
				for _, c := range strSlice[i] {
					decodedStr += decoderMap[string(c)]
				}
				str := sortStringByCharacter(decodedStr)
				println(str, "->", codeMap[str])
				stringToAdd = strconv.Itoa(codeMap[str]) + stringToAdd
			}

			num, err := strconv.Atoi(stringToAdd)
			if err != nil {
				fmt.Println(err)
			}
			count += num

			// reset for the next line
			toggle = false
			lastFour = 0
			strSlice = []string{}
		}
	}

	fmt.Println(count)

	// test
	//fmt.Println("3 should be ", numberOfCharsInCommon("darcet", "clytes"))
	//fmt.Println("d should be ", subtractBFromA("darcet", "create"))
}

func getDecoder(strz []string) map[string]string {
	m := map[string]string{"a":"", "b":"", "c":"", "d":"", "e":"", "f":"", "g":""}
	var one, four, seven, eight, five, three, two string

	// establish known numbers
	for _, str := range strz {
		if len(str) == 2 && one == "" {
			one = sortStringByCharacter(str)
		}
		if len(str) == 3 && seven == "" {
			seven = sortStringByCharacter(str)
		}
		if len(str) == 4 && four == "" {
			four = sortStringByCharacter(str)
		}
		if len(str) == 7 && eight == "" {
			eight = sortStringByCharacter(str)
		}
	}

	// seven has position A, one does not
	if len(subtractBFromA(seven, one)) == 1 {
		m["a"] = subtractBFromA(seven, one)
	}

	// on a digital clock, 2, 3, 5 have len 5
	for _, str := range strz {
		if len(str) == 5 {
			if numberOfCharsInCommon(str, four) == 3 && numberOfCharsInCommon(str, seven) == 2  && five == "" {
				five = sortStringByCharacter(str)
			}
			if strings.Contains(str, one) && three == "" {
				three = sortStringByCharacter(str)
			}
			if numberOfCharsInCommon(str, four) == 2 && two == "" {
				two = sortStringByCharacter(str)
			}
		}
	}

	if len(five) > 0 {
		m["g"] = subtractBFromA(five, four+seven)
		m["c"] = subtractBFromA(one, five)
		m["f"] = subtractBFromA(seven, m["c"] + m["a"])
		m["e"] = subtractBFromA(eight, four +  m["a"] + m["g"])

		if len(two) > 0 {
			m["d"] = subtractBFromA(two, m["a"] + m["c"] + m["g"] + m["e"])
		} else
		if len(three) > 0 {
			m["d"] = subtractBFromA(three, m["a"] + m["c"] + m["g"] + m["f"])
		}

		m["b"] = subtractBFromA(four, m["d"] + m["c"] + m["f"])

	} else {
		fmt.Println("shit, there is no FIVE in this list")
	}

	return m
}

func createReverseMap(m map[string]string) map[string]string {
	r := map[string]string{}
	for k, v := range m {
		r[v] = k
	}
	return r
}

func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func sortStringByCharacter(s string) string {
	r := stringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func numberOfCharsInCommon(a, b string) int {
	count := 0
	for _, char := range a {
		if strings.Contains(b, string(char)) {
			count++
		}
	}
	return count
}

func subtractBFromA(a, b string) string {
	newStr := ""
	for _, char := range a {
		if !strings.Contains(b, string(char)) {
			newStr += string(char)
		}
	}
	return newStr
}
