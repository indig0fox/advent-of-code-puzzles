package y2023d07

import (
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var puzzleTitle = "--- Day 7: Camel Cards ---"

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
	log.Logger = log.With().Timestamp().Str("puzzle", "2023_Day07").Logger()
	log.Logger = log.Level(zerolog.InfoLevel)
	log.Logger = log.Level(zerolog.DebugLevel)
	// log.Logger = log.Level(zerolog.TraceLevel)

	startTime := time.Now()
	defer func() {
		log.Info().Msgf(
			"%s completed in %s",
			puzzleTitle,
			time.Since(startTime),
		)
	}()

	log.Debug().Msg("Running y2023d07.Run()...")
	log.Info().Msg(puzzleTitle)

	data := getInput(inputPath)
	hands := parseHands(data)
	// sort hands by strength > first higher card > second higher card > etc.
	// don't count jokers for part 1
	hands = sortHands(hands, false)

	// Part 1: Find total winnings by multiplying bid with rank
	totalWinningsPt1 := 0
	for i, hand := range hands {
		totalWinningsPt1 += hand.bid * (i + 1)
	}

	log.Info().Msgf("Part 1: Total winnings: %d", totalWinningsPt1)

	// Part 2: Find total winnings by multiplying bid with rank using Jokers as wildcards!
	hands = sortHands(hands, true)

	totalWinningsPt2 := 0
	for i, hand := range hands {
		totalWinningsPt2 += hand.bid * (i + 1)
	}

	log.Info().Msgf("Part 2 (NOT CORRECT): Total winnings using Jokers as wildcards: %d", totalWinningsPt2)

	return resultStruct{
		Part1_SumOfTotalWinnings: totalWinningsPt1,
		// Part2_SumOfTotalWinningsWithJokers: totalWinningsPt2,
	}
}

func getInput(inputPath string) []string {
	inputData, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(inputData), "\n")
}

func parseHands(inputData []string) []*handOfCardsStruct {
	// Parse the hands from the input data
	hands := []*handOfCardsStruct{}
	for _, line := range inputData {
		lineSplit := strings.Fields(line)
		hand := handOfCardsStruct{
			cards: [5]cardStruct{},
			bid:   0,
		}
		for i, card := range lineSplit[0] {
			hand.cards[i] = cardStruct(card)
		}
		var err error
		hand.bid, err = strconv.Atoi(lineSplit[1])
		if err != nil {
			panic(err)
		}
		hands = append(hands, &hand)
		log.Trace().Object("hand", &hand).Msg("Parsed hand")
	}
	return hands
}

func sortHands(hands []*handOfCardsStruct, shouldCountJokers bool) []*handOfCardsStruct {
	// apply strength to each hand
	for i, hand := range hands {
		hands[i].strength = hand.getStrength(shouldCountJokers)
	}

	newHands := hands

	// sort hands by strength > first higher card > second higher card > etc.
	sort.SliceStable(newHands, func(i, j int) bool {
		if newHands[i].strength != newHands[j].strength {
			return newHands[i].strength < newHands[j].strength
		}

		// hands have the same strength, so sort by the first card which is higher
		for k := 0; k < 5; k++ {
			if cardTypes.getCardIndex(newHands[i].cards[k]) != cardTypes.getCardIndex(newHands[j].cards[k]) {
				return cardTypes.getCardIndex(newHands[i].cards[k]) < cardTypes.getCardIndex(newHands[j].cards[k])
			}
		}

		log.Warn().Msgf("Hands %v and %v are the same", newHands[i], newHands[j])
		return false
	})

	return newHands
}
