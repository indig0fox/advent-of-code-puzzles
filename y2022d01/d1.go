package y2022d01

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type resultStruct struct {
	Part1_MaxElfCalorieSum         int
	Part2_TopThreeElvesCaloriesSum int
}

var fLog *log.Logger

func Run(inputPath string, logFilePath string) resultStruct {
	logFile, err := os.Create(logFilePath)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	fLog = log.New(multiWriter, "", log.LstdFlags|log.Lshortfile)

	elfCalorieCounts := processInput(inputPath)

	// find the elf with the most calories
	var maxElf *elfStruct
	maxCalories := 0
	for _, elf := range elfCalorieCounts {
		if elf.TotalCalories > maxCalories {
			maxCalories = elf.TotalCalories
			maxElf = elf
		}
	}

	// sort the elves by total calories
	sort.Slice(elfCalorieCounts, func(i, j int) bool {
		return elfCalorieCounts[i].TotalCalories > elfCalorieCounts[j].TotalCalories
	})

	fLog.Println("---2022  Day 1: Calorie Counting ---")
	fLog.Printf("Part 1: The %dth elf was carrying the most calories (%d)", maxElf.Number, maxElf.TotalCalories)
	fLog.Printf("Part 2: The top three elves were carrying %d, %d, and %d calories with a total of %d",
		elfCalorieCounts[0].TotalCalories,
		elfCalorieCounts[1].TotalCalories,
		elfCalorieCounts[2].TotalCalories,
		elfCalorieCounts[0].TotalCalories+elfCalorieCounts[1].TotalCalories+elfCalorieCounts[2].TotalCalories,
	)

	return resultStruct{
		Part1_MaxElfCalorieSum:         maxElf.TotalCalories,
		Part2_TopThreeElvesCaloriesSum: elfCalorieCounts[0].TotalCalories + elfCalorieCounts[1].TotalCalories + elfCalorieCounts[2].TotalCalories,
	}
}

type elfStruct struct {
	Number        int
	TotalCalories int
}

func processInput(inputPath string) []*elfStruct {

	f, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scan := bufio.NewScanner(f)
	defer f.Close()

	elfCalorieCounts := []*elfStruct{}

	tempElf := &elfStruct{
		Number: 1,
	}

	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			elfCalorieCounts = append(elfCalorieCounts, tempElf)
			tempElf = &elfStruct{
				Number: tempElf.Number + 1,
			}
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		tempElf.TotalCalories += calories
	}

	return elfCalorieCounts
}
