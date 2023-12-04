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

/*
--- Day 1: Trebuchet?! ---
Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?
*/

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

	fLog.Println("--- Day 1: Trebuchet?! ---")
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
