package y2023d04

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type resultStruct struct {
	Part1_ScratchcardsPointSum  int
	Part2_TotalNumberOfWonCards int
}

var (
	fLog         *log.Logger
	scratchCards = []*scratchCardStruct{}
	newPile      = make(map[int]int)
)

func Run(inputPath string, logFilePath string) resultStruct {

	logFile, err := os.Create(logFilePath)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	fLog = log.New(multiWriter, "", log.LstdFlags|log.Lshortfile)

	data := getInput(inputPath)

	// PART 1: Get sum of scratch card points
	for _, line := range data {
		scratchCards = append(scratchCards, processScratchCardFromInputLine(line))
	}
	scratchCardPointsSum := 0
	for _, card := range scratchCards {
		scratchCardPointsSum += card.Score
	}

	// PART 2: Get total scratchcard count with won copies based on count of matching numbers

	for _, card := range scratchCards {
		// fLog.Println("Card", card.CardNumber, "has", len(card.MatchingNumbers), "matching numbers")
		addCardToNewPile(card)
	}
	totalScratchCardCount := 0
	for _, count := range newPile {
		totalScratchCardCount += count
	}

	fLog.Println("--- 2023 Day 4: Scratch Cards ---")
	fLog.Println("Part 1: Scratchcard points sum is", scratchCardPointsSum)
	fLog.Println("Part 2: Total scratchcard count (including copies) is", totalScratchCardCount)
	return resultStruct{
		Part1_ScratchcardsPointSum:  scratchCardPointsSum,
		Part2_TotalNumberOfWonCards: totalScratchCardCount,
	}
}

func addCardToNewPile(card *scratchCardStruct) {
	if _, ok := newPile[card.CardNumber]; ok {
		newPile[card.CardNumber] = newPile[card.CardNumber] + 1
	} else {
		newPile[card.CardNumber] = 1
	}

	for i := 0; i < len(card.MatchingNumbers); i++ {
		copiedCard := scratchCards[card.CardNumber+i]
		addCardToNewPile(copiedCard)
	}
}

func getInput(inputPath string) []string {
	f, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(f), "\n")

}

func processScratchCardFromInputLine(line string) *scratchCardStruct {

	splitLine := strings.Split(line, "|")
	// First half (header + winning numbers)
	leaderSubstring := splitLine[0]
	leaderRegex := regexp.MustCompile(`\d+`)
	leaderMatches := leaderRegex.FindAllString(leaderSubstring, -1)
	cardNum, err := strconv.Atoi(leaderMatches[0])
	if err != nil {
		panic(err)
	}

	winningNumbersInt := []int{}
	for _, leaderMatch := range leaderMatches[1:] {
		winningNumberInt := 0
		_, err := fmt.Sscanf(leaderMatch, "%d", &winningNumberInt)
		if err != nil {
			panic(err)
		}
		winningNumbersInt = append(winningNumbersInt, winningNumberInt)
	}

	// Second half (has numbers)
	trailerSubstring := splitLine[1]
	hasNumbersRegex := regexp.MustCompile(`\d+`)
	hasMatches := hasNumbersRegex.FindAllString(trailerSubstring, -1)
	hasMatchesInt := []int{}
	for _, hasMatch := range hasMatches {
		hasMatchInt := 0
		_, err := fmt.Sscanf(hasMatch, "%d", &hasMatchInt)
		if err != nil {
			panic(err)
		}
		hasMatchesInt = append(hasMatchesInt, hasMatchInt)
	}

	thisCard := &scratchCardStruct{
		CardNumber:      cardNum,
		WinningNumbers:  winningNumbersInt,
		HasNumbers:      hasMatchesInt,
		MatchingNumbers: []int{},
		Score:           0,
	}

	for _, winningNumber := range winningNumbersInt {
		if thisCard.hasMatch(winningNumber) {
			thisCard.MatchingNumbers = append(thisCard.MatchingNumbers, winningNumber)
		}
	}

	thisCard.updateScore()
	// fLog.Println(thisCard)
	return thisCard
}

type scratchCardStruct struct {
	CardNumber      int
	WinningNumbers  []int
	HasNumbers      []int
	MatchingNumbers []int
	Score           int
}

func (s *scratchCardStruct) String() string {
	return fmt.Sprintf(
		"Card %d has %d matching numbers, worth %d points.",
		s.CardNumber,
		len(s.MatchingNumbers),
		s.Score,
	)
}

func (s *scratchCardStruct) hasMatch(winningNumber int) bool {
	for _, hasNumber := range s.HasNumbers {
		if hasNumber == winningNumber {
			return true
		}
	}
	return false
}

func (s *scratchCardStruct) updateScore() int {
	s.Score = 0
	for range s.MatchingNumbers {
		if s.Score == 0 {
			s.Score = 1
		} else {
			s.Score = s.Score * 2
		}
	}
	return s.Score
}
