package y2023d03

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type d3ResultStruct struct {
	Part1_PartNumbersSum int
	Part2_GearRatiosSum  int
}

var fLog *log.Logger
var inputData []string
var foundNumbers []foundNumberStruct
var foundGears []foundGearStruct

func Run(inputPath string, logFilePath string) d3ResultStruct {

	logFile, err := os.Create(logFilePath)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	fLog = log.New(multiWriter, "", log.LstdFlags|log.Lshortfile)

	fLog.Println("Running d3.Run()")
	fLog.Println("inputPath:", inputPath)
	fLog.Println("logFilePath:", logFilePath)

	fLog.Println("Getting game data from input...")
	inputData = getGameDataFromInput(inputPath)

	fLog.Println("Processing input lines...")
	processInputLines(inputData)
	fLog.Printf("Found a total %v numbers\n", len(foundNumbers))

	fLog.Println("Determining which numbers qualify as part numbers...")
	partNumbers := []foundNumberStruct{}
	for _, number := range foundNumbers {
		if number.qualifiesAsPartNumber() {
			partNumbers = append(partNumbers, number)
		}
	}
	fLog.Printf("%v/%v numbers qualify as part numbers\n", len(partNumbers), len(foundNumbers))

	fLog.Println("Calculating the sum of all part numbers...")
	sum := 0
	for _, number := range partNumbers {
		sum = sum + number.number
	}

	fLog.Println("Sum of all part numbers:", sum)

	fLog.Println("Determining which gears have valid ratios...")
	gearsWithAdjacentPartNumbers := []foundGearStruct{}
	for _, gear := range foundGears {
		if gear.gearRatio() > 0 {
			gearsWithAdjacentPartNumbers = append(gearsWithAdjacentPartNumbers, gear)
		}
	}
	fLog.Printf("%v/%v gears have valid ratios.\n", len(gearsWithAdjacentPartNumbers), len(foundGears))

	fLog.Println("Calculating the sum of all gear ratios...")
	gearRatiosSum := 0
	for _, gear := range gearsWithAdjacentPartNumbers {
		gearRatiosSum = gearRatiosSum + gear.gearRatio()
	}

	fLog.Println("--- 2023 Day 3: Gear Ratios ---")
	fLog.Println("Part 1: Sum of all adjacent part numbers:", gearRatiosSum)
	fLog.Println("Part 2: Sum of all gear ratios:", gearRatiosSum)

	return d3ResultStruct{
		Part1_PartNumbersSum: sum,
		Part2_GearRatiosSum:  gearRatiosSum,
	}
}

func getGameDataFromInput(inputPath string) []string {
	inputF, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer inputF.Close()

	data := []string{}
	scanner := bufio.NewScanner(inputF)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return data
}

func processInputLines(lines []string) error {
	for i, line := range lines {
		// create a slice of all numbers in the line
		foundNumbers = append(foundNumbers, getNumbersFromLine(i, line)...)
		// create a slice of all gears in the line
		foundGears = append(foundGears, getGearsFromLine(i, line)...)
	}

	return nil
}

func getNumbersFromLine(lineIndex int, line string) []foundNumberStruct {
	numbers := []foundNumberStruct{}
	curIndex := 0
	for {
		// find the next number
		nextIndex := strings.IndexAny(line[curIndex:], "0123456789")
		if nextIndex == -1 {
			break
		}
		curIndex = curIndex + nextIndex
		// find the end of the number
		endIndex := strings.IndexAny(line[curIndex:], "!@#$%^&*()_+-=,./<>?;':\"[]{}\\|~` ")
		if endIndex == -1 {
			endIndex = len(line)
		} else {
			endIndex = curIndex + endIndex
		}
		// convert the number to int
		number, err := strconv.Atoi(line[curIndex:endIndex])
		if err != nil {
			panic(err)
		}
		// add the number to the slice
		numbers = append(numbers, foundNumberStruct{
			number:    number,
			index:     curIndex,
			lineIndex: lineIndex,
			length:    endIndex - curIndex,
		})
		// move the cursor to the end of the number
		curIndex = endIndex
	}
	return numbers
}

func getGearsFromLine(lineIndex int, line string) []foundGearStruct {
	gears := []foundGearStruct{}
	curIndex := 0
	for {
		// find the next gear
		nextIndex := strings.IndexAny(line[curIndex:], "*")
		if nextIndex == -1 {
			break
		}
		curIndex = curIndex + nextIndex
		// add the gear to the slice
		gears = append(gears, foundGearStruct{
			index:     curIndex,
			lineIndex: lineIndex,
		})
		// move the cursor to the end of the gear
		curIndex = curIndex + 1
	}
	return gears
}

func isSymbol(char byte) bool {
	return strings.ContainsAny(string(char), "!@#$%^&*()_+-=,/<>?;':\"[]{}\\|~`")
}

type foundNumberStruct struct {
	length    int
	number    int
	index     int
	lineIndex int
}

func (f *foundNumberStruct) qualifiesAsPartNumber() bool {
	// check if the number is adjacent to a symbol

	// check the line above
	if f.lineIndex > 0 {
		// check each char in line above
		for i := -1; i < f.length+1; i++ {
			if (f.index+i) < 0 || (f.index+i) >= len(inputData[f.lineIndex-1]) {
				continue
			}
			if isSymbol(inputData[f.lineIndex-1][f.index+i]) {
				return true
			}
		}
	}

	// check the line
	// check the number to the left
	if f.index > 0 {
		if isSymbol(inputData[f.lineIndex][f.index-1]) {
			return true
		}
	}
	// check the number to the right
	if f.index+f.length < len(inputData[f.lineIndex]) {
		if isSymbol(inputData[f.lineIndex][f.index+f.length]) {
			return true
		}
	}

	// check the line below
	if f.lineIndex < len(inputData)-1 {
		// check each char in line below
		for i := -1; i < f.length+1; i++ {
			if (f.index+i) < 0 || (f.index+i) >= len(inputData[f.lineIndex+1]) {
				continue
			}
			if isSymbol(inputData[f.lineIndex+1][f.index+i]) {
				return true
			}
		}
	}

	return false
}

type foundGearStruct struct {
	index     int
	lineIndex int
}

func (f *foundGearStruct) gearRatio() int {
	if !f.isGear() {
		return 0
	}

	adjacentNumbers := f.getAdjacentPartNumbers()
	return adjacentNumbers[0].number * adjacentNumbers[1].number

}

func (f *foundGearStruct) isGear() bool {
	return len(f.getAdjacentPartNumbers()) == 2
}

func (f *foundGearStruct) getAdjacentPartNumbers() []foundNumberStruct {
	adjacentNumbers := []foundNumberStruct{}

	for _, number := range foundNumbers {
		// check if the number is on the same line
		if number.lineIndex == f.lineIndex {
			for i := -1; i < number.length+1; i++ {
				if number.index+i == f.index {
					adjacentNumbers = append(adjacentNumbers, number)
					break
				}
			}
		} else if number.lineIndex == f.lineIndex-1 {
			// check if the number is on the line above
			for i := -1; i < number.length+1; i++ {
				if number.index+i == f.index {
					adjacentNumbers = append(adjacentNumbers, number)
					break
				}
			}
		} else if number.lineIndex == f.lineIndex+1 {
			// check if the number is on the line below
			for i := -1; i < number.length+1; i++ {
				if number.index+i == f.index {
					adjacentNumbers = append(adjacentNumbers, number)
					break
				}
			}
		}
	}

	return adjacentNumbers
}
