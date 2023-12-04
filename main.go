package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/indig0fox/advent-of-code-puzzles-2023/d1"
	"github.com/indig0fox/advent-of-code-puzzles-2023/d2"
	"github.com/indig0fox/advent-of-code-puzzles-2023/d3"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	results := make(map[string]interface{})

	// Day 1
	results["Day01"] = d1.Run("d1/input.txt", "d1/log.txt")
	// Day 2
	results["Day02"] = d2.Run("d2/input.txt", "d2/log.txt")
	// Day 3
	results["Day03"] = d3.Run("d3/input.txt", "d3/log.txt")

	jsonResults, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		panic(err)
	}
	log.Printf("Results: %v\n", string(jsonResults))
	resultsFile, err := os.Create("results.json")
	if err != nil {
		panic(err)
	}
	defer resultsFile.Close()
	resultsFile.Write(jsonResults)
}
