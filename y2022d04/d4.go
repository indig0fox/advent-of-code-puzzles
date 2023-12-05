package y2022d04

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type resultStruct struct {
	Part1_PairsWhereOneRangeFullyContainsTheOther     int
	Part2_PairsWhereOneRangePartiallyContainsTheOther int
}

var (
	fLog *log.Logger
)

type shipAssignmentPair struct {
	Assignment1      map[int]bool
	Assignment1Start int
	Assignment1End   int

	Assignment2      map[int]bool
	Assignment2Start int
	Assignment2End   int
	PairIndex        int
}

// func (s shipAssignmentPair) showAssignments() string {
// 	outString := ""

// 	for i := 1; i < 100; i++ {
// 		if s.Assignment1[i] {
// 			outString += "X"
// 		} else {
// 			outString += "_"
// 		}
// 	}
// 	outString += "\n"
// 	for i := 1; i < 100; i++ {
// 		if s.Assignment2[i] {
// 			outString += "X"
// 		} else {
// 			outString += "_"
// 		}
// 	}
// 	return outString
// }

func (s shipAssignmentPair) doesOneRangeFullyContainTheOther() bool {
	// assignment 1 fully contains assignment 2
	if s.Assignment1Start <= s.Assignment2Start && s.Assignment1End >= s.Assignment2End {
		return true
	}

	// assignment 2 fully contains assignment 1
	if s.Assignment2Start <= s.Assignment1Start && s.Assignment2End >= s.Assignment1End {
		return true
	}

	return false
}

func (s shipAssignmentPair) doesOneRangePartiallyContainTheOther() bool {
	for i := 1; i < 100; i++ {
		if s.Assignment1[i] && s.Assignment2[i] {
			return true
		}
	}
	return false
}

func Run(inputPath string, logFilePath string) resultStruct {

	logFile, err := os.Create(logFilePath)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	fLog = log.New(multiWriter, "", log.LstdFlags|log.Lshortfile)

	data := getInput(inputPath)
	assignmentMaps := convertInputToAssignmentMaps(data)

	// PART 1: Find assignment pairs where one fully contains the other
	totalPart1 := 0
	for _, pair := range assignmentMaps {
		// fLog.Println(pair)
		if pair.doesOneRangeFullyContainTheOther() {
			totalPart1++
		}
	}

	// PART 2: Find assignment pairs where one partially contains the other
	totalPart2 := 0
	for _, pair := range assignmentMaps {
		if pair.doesOneRangePartiallyContainTheOther() {
			totalPart2++
		}
	}

	fLog.Println("--- 2022 Day 4: Camp Cleanup ---")
	fLog.Println("Part 1: Number of pairs in which one range fully contains the other:", totalPart1)
	fLog.Println("Part 2: Number of pairs in which one range partially contains the other:", totalPart2)

	return resultStruct{
		Part1_PairsWhereOneRangeFullyContainsTheOther:     totalPart1,
		Part2_PairsWhereOneRangePartiallyContainsTheOther: totalPart2,
	}
}

func getInput(inputPath string) []string {
	fLog.Println("Getting game data from input...")
	bytes, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	inputData := strings.Split(string(bytes), "\n")
	return inputData
}

func convertInputToAssignmentMaps(input []string) []shipAssignmentPair {
	assignmentMaps := []shipAssignmentPair{}

	for pairIndex, line := range input {
		lineSplit := strings.Split(line, ",")
		for i := 0; i < len(lineSplit); i++ {
			lineSplit[i] = strings.TrimSpace(lineSplit[i])
		}
		assignmentMap := shipAssignmentPair{
			Assignment1: map[int]bool{},
			Assignment2: map[int]bool{},
			PairIndex:   pairIndex,
		}

		for i := 1; i < 100; i++ {
			assignmentMap.Assignment1[i] = false
			assignmentMap.Assignment2[i] = false
		}

		// assignment 1
		assignment1Split := strings.Split(lineSplit[0], "-")

		// get ints
		assignmentSplitIntStart := 0
		assignmentSplitIntEnd := 0

		var err error
		assignmentSplitIntStart, err = strconv.Atoi(assignment1Split[0])
		if err != nil {
			panic(err)
		}
		assignmentSplitIntEnd, err = strconv.Atoi(assignment1Split[1])
		if err != nil {
			panic(err)
		}

		assignmentMap.Assignment1Start = assignmentSplitIntStart
		assignmentMap.Assignment1End = assignmentSplitIntEnd

		for i := assignmentSplitIntStart; i <= assignmentSplitIntEnd; i++ {
			assignmentMap.Assignment1[i] = true
		}

		// assignment 2
		assignment2Split := strings.Split(lineSplit[1], "-")

		// get ints
		assignmentSplitIntStart = 0
		assignmentSplitIntEnd = 0

		assignmentSplitIntStart, err = strconv.Atoi(assignment2Split[0])
		if err != nil {
			panic(err)
		}
		assignmentSplitIntEnd, err = strconv.Atoi(assignment2Split[1])
		if err != nil {
			panic(err)
		}

		assignmentMap.Assignment2Start = assignmentSplitIntStart
		assignmentMap.Assignment2End = assignmentSplitIntEnd

		for i := assignmentSplitIntStart; i <= assignmentSplitIntEnd; i++ {
			assignmentMap.Assignment2[i] = true
		}

		assignmentMaps = append(assignmentMaps, assignmentMap)

	}
	return assignmentMaps
}
