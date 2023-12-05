package main

import (
	"encoding/json"
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

}

func run2022puzzles(results *resultsStruct) *resultsStruct {

	results.Year2022["Day01"] = y2022d01.Run("y2022d01/input.txt", "y2022d01/log.txt")
	results.Year2022["Day02"] = y2022d02.Run("y2022d02/input.txt", "y2022d02/log.txt")
	results.Year2022["Day03"] = y2022d03.Run("y2022d03/input.txt", "y2022d03/log.txt")
	results.Year2022["Day04"] = y2022d04.Run("y2022d04/input.txt", "y2022d04/log.txt")
	results.Year2022["Day05"] = y2022d05.Run("y2022d05/input.txt", "y2022d05/log.txt")

	return results
}

func run2023puzzles(results *resultsStruct) *resultsStruct {
	results.Year2023["Day01"] = y2023d01.Run("y2023d01/input.txt", "y2023d01/log.txt")
	results.Year2023["Day02"] = y2023d02.Run("y2023d02/input.txt", "y2023d02/log.txt")
	results.Year2023["Day03"] = y2023d03.Run("y2023d03/input.txt", "y2023d03/log.txt")
	results.Year2023["Day04"] = y2023d04.Run("y2023d04/input.txt", "y2023d04/log.txt")

	return results
}
