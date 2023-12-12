package y2023d09

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var puzzleTitle = "--- Day 9: Mirage Maintenance ---"

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
	log.Logger = log.With().Timestamp().Str("puzzle", "2023_Day09").Logger()
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

	log.Debug().Msg("Running y2023d08.Run()...")
	log.Info().Msg(puzzleTitle)

	data := getInput(inputPath)
	valuesHistory := parseValues(data)

	// Part 1: Extrapolate forwards
	for i := range valuesHistory {
		valuesHistory[i].CalculateDeltas()
		valuesHistory[i].PredictNextHistory()
	}

	part1LastHistoryValueSum := 0
	for i := range valuesHistory {
		part1LastHistoryValueSum += valuesHistory[i].CurrentHistory[len(valuesHistory[i].CurrentHistory)-1]
	}

	log.Info().Msgf("Part 1: %v", part1LastHistoryValueSum)

	// Part 2: Extrapolate backwards
	valuesHistoryPt2 := parseValues(data)

	for i := range valuesHistoryPt2 {
		valuesHistoryPt2[i].CalculateDeltas()
		valuesHistoryPt2[i].PredictPreviousHistoryPt2()
	}

	part2LastHistoryValueSum := 0
	for i := range valuesHistoryPt2 {
		part2LastHistoryValueSum += valuesHistoryPt2[i].CurrentHistory[0]
	}

	log.Info().Msgf("Part 2: %v", part2LastHistoryValueSum)

	return resultStruct{
		Part1_HistoryExtrapolatedForwardSum:  part1LastHistoryValueSum,
		Part2_HistoryExtrapolatedBackwardSum: part2LastHistoryValueSum,
	}
}

func getInput(inputPath string) []string {
	f, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

func parseValues(data []string) []*recordStruct {
	var values []*recordStruct
	for _, line := range data {
		thisRecord := recordStruct{}
		line = strings.TrimSpace(line)
		history := strings.Fields(line)
		log.Trace().Msgf("history: %v", history)
		for _, value := range history {
			historyInt, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			thisRecord.AddHistory(historyInt)
		}
		values = append(values, &thisRecord)
	}
	return values
}
