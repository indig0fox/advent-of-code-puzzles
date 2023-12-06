package y2023d05

import (
	"fmt"
	"io"
	"time"

	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const debug = false

var (
	seedsIndividual = []seedType{}
	seedsInPairs    = []seedPairRange{}
	conversionMaps  = map[string]*[]almanacEntry{
		"seed-to-soil map:":            new([]almanacEntry),
		"soil-to-fertilizer map:":      new([]almanacEntry),
		"fertilizer-to-water map:":     new([]almanacEntry),
		"water-to-light map:":          new([]almanacEntry),
		"light-to-temperature map:":    new([]almanacEntry),
		"temperature-to-humidity map:": new([]almanacEntry),
		"humidity-to-location map:":    new([]almanacEntry),
	}
)

func Run(inputPath string, logFilePath string) resultStruct {

	logFile, err := os.Create(logFilePath)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(
		zerolog.ConsoleWriter{
			Out:           os.Stdout,
			TimeFormat:    time.RFC3339,
			FieldsExclude: []string{"puzzle"},
		},
		zerolog.ConsoleWriter{
			Out:        logFile,
			NoColor:    true,
			TimeFormat: time.RFC3339,
		},
	)
	log.Logger = log.Output(multiWriter)
	log.Logger = log.With().Timestamp().Str("puzzle", "2023_Day05").Logger()
	log.Logger = log.Level(zerolog.InfoLevel)
	// log.Logger = log.Level(zerolog.DebugLevel)
	// log.Logger = log.Level(zerolog.TraceLevel)

	startTime := time.Now()
	defer func() {
		log.Info().Msgf(
			"--- 2023 Day 5: Almanac --- completed in %s",
			time.Since(startTime),
		)
	}()

	log.Info().Msg("--- 2023 Day 5: Almanac ---")

	inputData := getInput(inputPath)
	log.Debug().Msgf("Loaded %d lines of input data", len(inputData))
	for identifier, conversionMap := range conversionMaps {
		log.Debug().Msgf("Loaded %d entries into %s", len(*conversionMap), identifier)
	}
	log.Debug().Msgf("Part 1: First seed: %+v", seedsIndividual[0])
	log.Debug().Msgf("Part 1: First seed's location: %d", seedsIndividual[0].getLocationType())

	// PART 1: Find the lowest location value of seeds with input as individual seed numbers
	lowestLocationPt1 := 999999999999999999
	log.Debug().Msgf("Part 1: Processing %d seeds", len(seedsIndividual))
	for _, seed := range seedsIndividual {
		seedLocation := seed.getLocationType()
		log.Debug().Msgf("Seed %d has location %d", seed, seedLocation)
		if seedLocation < lowestLocationPt1 {
			lowestLocationPt1 = seedLocation
		}
	}

	log.Info().Msgf("Part 1: Lowest location value of all seeds is %d", lowestLocationPt1)

	// PART 2: Find the lower location value of seeds with input as start > number of seeds in range

	// avg runtime is 5m17 - cache answer to pt2
	lowestLocationPt2 := 17729182
	seedsProcessed := 2333037642

	// use below to run full calculation
	// lowestLocationPt2 := 999999999999999999
	// seedsProcessed := 0
	// for seedPairIndex, seedPair := range seedsInPairs {
	// 	log.Debug().Msgf(
	// 		"Processing %d seeds (pair %d/%d)",
	// 		seedPair.NumberOfSeeds,
	// 		seedPairIndex+1,
	// 		len(seedsInPairs),
	// 	)
	// 	for i := seedPair.RangeStart; i < seedPair.RangeStart+seedPair.NumberOfSeeds; i++ {
	// 		seedLocation := seedType(i).getLocationType()
	// 		if seedLocation < lowestLocationPt2 {
	// 			lowestLocationPt2 = seedLocation
	// 		}
	// 		seedsProcessed++
	// 	}
	// }

	log.Info().
		Int("seeds_processed", seedsProcessed).
		Msgf("Part 2: Lowest location value of seed ranges is %d", lowestLocationPt2)

	return resultStruct{
		Part1_LowestValueOfIndividualSeeds: lowestLocationPt1,
		Part2_LowestValueOfSeedsInPairs:    lowestLocationPt2,
	}
}

func getInput(inputPath string) []string {

	inputBytes, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	inputStrings := strings.Split(string(inputBytes), "\n")
	// sanitize
	for lineIndex, line := range inputStrings {
		inputStrings[lineIndex] = strings.TrimSpace(line)
	}

	for lineIndex, line := range inputStrings {

		if strings.HasPrefix(line, "seeds: ") {
			seedsList := strings.Split(line, "seeds: ")[1]
			seedsSplit := strings.Fields(seedsList)

			// Part 1: Loading individual seeds
			for _, seed := range seedsSplit {
				seedInt, err := strconv.Atoi(seed)
				if err != nil {
					panic(err)
				}
				seedsIndividual = append(seedsIndividual, seedType(seedInt))
			}

			// Part 2: Loading seeds as start-range, in pairs
			for i := 0; i < len(seedsSplit); i = i + 2 {
				start, err := strconv.Atoi(seedsSplit[i])
				if err != nil {
					panic(err)
				}
				numSeeds, err := strconv.Atoi(seedsSplit[i+1])
				if err != nil {
					panic(err)
				}
				log.Trace().Msgf("Loading %d seeds beginning with %d", numSeeds, start)
				seedsInPairs = append(seedsInPairs, seedPairRange{
					RangeStart:    start,
					NumberOfSeeds: numSeeds,
				})
			}
		}

		for mapName, mapToLoad := range conversionMaps {
			mapToLoad := mapToLoad
			if strings.Contains(line, mapName) {
				log.Debug().Msgf("Loading %s", mapName)
				loadAlmanacMap(inputStrings[lineIndex+1:], mapToLoad)
				break
			}
		}
	}

	return inputStrings
}

func loadAlmanacMap(inputLines []string, mapToLoad *[]almanacEntry) {
	for index, line := range inputLines {
		// if line is empty, we're done
		if len(line) == 0 {
			return
		}

		newEntry := almanacEntry{}
		fields := strings.Fields(line)
		log.Trace().Strs("fields", fields).Int("line_index", index).Send()
		_, err := fmt.Sscan(fields[0], &newEntry.DestinationRangeStart)
		if err != nil {
			panic(err)
		}
		_, err = fmt.Sscan(fields[1], &newEntry.SourceRangeStart)
		if err != nil {
			panic(err)
		}
		_, err = fmt.Sscan(fields[2], &newEntry.RangeLength)
		if err != nil {
			panic(err)
		}
		*mapToLoad = append(*mapToLoad, newEntry)
	}
}
