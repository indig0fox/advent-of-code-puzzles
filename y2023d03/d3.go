package y2023d03

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
--- Day 3: Gear Ratios ---
You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?

Your puzzle answer was 536202.

--- Part Two ---
The engineer finds the missing part and installs it in the engine! As the engine springs to life, you jump in the closest gondola, finally ready to ascend to the water source.

You don't seem to be going very fast, though. Maybe something is still wrong? Fortunately, the gondola has a phone labeled "help", so you pick it up and the engineer answers.

Before you can explain the situation, she suggests that you look out the window. There stands the engineer, holding a phone in one hand and waving with the other. You're going so slowly that you haven't even left the station. You exit the gondola.

The missing part wasn't the only issue - one of the gears in the engine is wrong. A gear is any * symbol that is adjacent to exactly two part numbers. Its gear ratio is the result of multiplying those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so that the engineer can figure out which gear needs to be replaced.

Consider the same engine schematic again:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, there are two gears. The first is in the top left; it has part numbers 467 and 35, so its gear ratio is 16345. The second gear is in the lower right; its gear ratio is 451490. (The * adjacent to 617 is not a gear because it is only adjacent to one part number.) Adding up all of the gear ratios produces 467835.

What is the sum of all of the gear ratios in your engine schematic?

Your puzzle answer was 78272573.
*/

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

	fLog.Println("Sum of all adjacent part numbers:", gearRatiosSum)

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
