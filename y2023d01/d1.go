package y2023d01

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type d1ResultStruct struct {
	Part1_CalibrationValuesSum         int
	Part2_CalibrationValuesWithTextSum int
}

var fLog *log.Logger

func Run(inputPath string, logFilePath string) d1ResultStruct {

	logFile, err := os.Create(logFilePath)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	fLog = log.New(multiWriter, "", log.LstdFlags|log.Lshortfile)

	lines := getInput(inputPath)

	part1Values := getCalibrationValues(lines, false)
	// fmt.Printf("Calibration values: %v\n", values)
	part1Sum := 0
	for _, value := range part1Values {
		part1Sum += value
	}

	part2Values := getCalibrationValues(lines, true)
	// fmt.Printf("Calibration values: %v\n", values)
	part2Sum := 0
	for _, value := range part2Values {
		part2Sum += value
	}

	fLog.Println("--- 2023 Day 1: Trebuchet?! ---")
	fLog.Println("Part 1 sum:", part1Sum)
	fLog.Println("Part 2 sum:", part2Sum)

	return d1ResultStruct{
		Part1_CalibrationValuesSum:         part1Sum,
		Part2_CalibrationValuesWithTextSum: part2Sum,
	}
}

type searchValueFound struct {
	index int
	value string
}
type matches []searchValueFound

// Combine first and last digit of given lines to form a single two-digit number
func getCalibrationValues(lines []string, includeNumbersAsText bool) []int {
	var values []int
	for _, line := range lines {
		// use regex match
		// https://golang.org/pkg/regexp/syntax/

		// fmt.Println("--------------------------------------------------")
		// fmt.Println("Raw input:", line)

		var searched matches = make([]searchValueFound, 0)

		searchValues := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		if includeNumbersAsText {
			searchValues = append(searchValues, []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}...)
		}

		// get all instances of search values and their starting index using strings.Index
		for _, searchValue := range searchValues {
			curIndex := 0
			for strings.Contains(line[curIndex:], searchValue) {
				index := strings.Index(line[curIndex:], searchValue)
				// fmt.Println("Found:", searchValue, "at index:", index, "in:", line[curIndex:])
				searched = append(searched, searchValueFound{curIndex + index, numberTextToNumberAsString(searchValue)})
				curIndex += index + 1
			}
		}

		// fmt.Println("Searched:", searched)

		// sort by index
		sort.Slice(searched, func(i, j int) bool {
			return searched[i].index < searched[j].index
		})

		// fmt.Println("Sorted:", searched)

		firstDigit := searched[0].value
		lastDigit := searched[len(searched)-1].value

		// fmt.Println("First:", string(firstDigit), "Last:", string(lastDigit))

		parsed, _ := strconv.Atoi(fmt.Sprintf("%s%s", firstDigit, lastDigit))
		// fmt.Println("Parsed:", parsed)
		values = append(values, parsed)
	}

	return values
}

func numberTextToNumberAsString(numberText string) string {
	switch numberText {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return numberText
	}
}

func getInput(path string) []string {
	// read from file

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	// split by new line
	lines := strings.Split(string(data), "\n")

	return lines
}
