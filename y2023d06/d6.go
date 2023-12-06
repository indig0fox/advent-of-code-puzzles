package y2023d06

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const puzzleTitle = "--- Day 6: Wait For It ---"

var availableRaces = []*raceStruct{
	{Time: 60, DistanceRecord: 475},
	{Time: 94, DistanceRecord: 2138},
	{Time: 78, DistanceRecord: 1015},
	{Time: 82, DistanceRecord: 1650},
}

var part2Race = raceStruct{
	Time:           60947882,
	DistanceRecord: 475213810151650,
}

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
	log.Logger = log.With().Timestamp().Str("puzzle", "2023_Day06").Logger()
	log.Logger = log.Level(zerolog.InfoLevel)
	// log.Logger = log.Level(zerolog.DebugLevel)
	// log.Logger = log.Level(zerolog.TraceLevel)

	startTime := time.Now()
	defer func() {
		log.Info().Msgf(
			"%s completed in %s",
			puzzleTitle,
			time.Since(startTime),
		)
	}()

	log.Debug().Msg("Running y2023d06.Run()...")
	log.Info().Msg(puzzleTitle)

	// iterate through races so that each race has a slice of winning button press times
	for _, thisRace := range availableRaces {
		thisRace.getWinningTimes()
	}

	// Part 1: Find the number of ways each race can be won, multiplied by each other
	part1Result := 1
	for _, thisRace := range availableRaces {
		part1Result = part1Result * len(thisRace.WinningButtonPressTimes)
	}

	log.Info().Msgf("Part 1: Product of the number of ways to win each race: %v", part1Result)

	// Part 2: Find the number of ways to beat a single, much longer race
	part2Result := len(part2Race.getWinningTimes())

	log.Info().Msgf("Part 2: Number of ways to win the very long race: %v", part2Result)

	return resultStruct{
		Part1_ProductOfWaysToWinThreeRaces:  part1Result,
		Part2_NumberOfWaysToWinVeryLongRace: part2Result,
	}
}
