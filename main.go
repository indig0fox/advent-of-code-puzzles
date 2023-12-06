package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/indig0fox/advent-of-code-puzzles/y2022d01"
	"github.com/indig0fox/advent-of-code-puzzles/y2022d02"
	"github.com/indig0fox/advent-of-code-puzzles/y2022d03"
	"github.com/indig0fox/advent-of-code-puzzles/y2022d04"
	"github.com/indig0fox/advent-of-code-puzzles/y2022d05"
	"github.com/indig0fox/advent-of-code-puzzles/y2023d01"
	"github.com/indig0fox/advent-of-code-puzzles/y2023d02"
	"github.com/indig0fox/advent-of-code-puzzles/y2023d03"
	"github.com/indig0fox/advent-of-code-puzzles/y2023d04"
	"github.com/indig0fox/advent-of-code-puzzles/y2023d05"
	"github.com/indig0fox/advent-of-code-puzzles/y2023d06"
)

type resultsStruct struct {
	Year2022 map[string]interface{}
	Year2023 map[string]interface{}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var results = &resultsStruct{
		Year2022: make(map[string]interface{}),
		Year2023: make(map[string]interface{}),
	}
	resultsFile, err := os.Create("results.json")
	if err != nil {
		panic(err)
	}
	defer resultsFile.Close()

	run2022puzzles(results)
	run2023puzzles(results)

	jsonResults, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		panic(err)
	}

	resultsFile.Write(jsonResults)
	// log.Printf("Results: %v\n", string(jsonResults))

	// await user input
	fmt.Println("Press enter to exit...")
	fmt.Scanln()
}

func run2022puzzles(results *resultsStruct) *resultsStruct {

	getInputFile := func(day string) string {
		return "y2022d" + day + "/input.txt"
	}
	getLogFile := func(day string) string {
		return "y2022d" + day + "/log.txt"
	}
	results.Year2022 = map[string]interface{}{
		"Day01": y2022d01.Run(getInputFile("01"), getLogFile("01")),
		"Day02": y2022d02.Run(getInputFile("02"), getLogFile("02")),
		"Day03": y2022d03.Run(getInputFile("03"), getLogFile("03")),
		"Day04": y2022d04.Run(getInputFile("04"), getLogFile("04")),
		"Day05": y2022d05.Run(getInputFile("05"), getLogFile("05")),
	}

	return results
}

func run2023puzzles(results *resultsStruct) *resultsStruct {
	getInputFile := func(day string) string {
		return "y2023d" + day + "/input.txt"
	}
	getLogFile := func(day string) string {
		return "y2023d" + day + "/log.txt"
	}
	results.Year2023 = map[string]interface{}{
		"Day01": y2023d01.Run(getInputFile("01"), getLogFile("01")),
		"Day02": y2023d02.Run(getInputFile("02"), getLogFile("02")),
		"Day03": y2023d03.Run(getInputFile("03"), getLogFile("03")),
		"Day04": y2023d04.Run(getInputFile("04"), getLogFile("04")),
		"Day05": y2023d05.Run(getInputFile("05"), getLogFile("05")),
		"Day06": y2023d06.Run(getInputFile("06"), getLogFile("06")),
	}

	return results
}
