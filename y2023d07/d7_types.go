package y2023d07

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type resultStruct struct {
	Part1_SumOfTotalWinnings           int
	Part2_SumOfTotalWinningsWithJokers int
}

var strengthsMap = map[string]int{
	"five of a kind":  6,
	"four of a kind":  5,
	"full house":      4,
	"three of a kind": 3,
	"two pair":        2,
	"one pair":        1,
	"high card":       0,
}

type cardStruct rune // 2-9, T, J, Q, K, A

var cardTypes = cardTypesStruct{
	'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A',
}

type cardTypesStruct []cardStruct

func (c cardTypesStruct) getCardIndex(card cardStruct) int {
	for i, cardType := range c {
		if cardType == card {
			return i
		}
	}
	return 0
}

func testCardMatches(cards [5]cardStruct) []cardMatchStruct {
	matches := []cardMatchStruct{}
	for _, cardType := range cardTypes {
		thisCardTypeMatch := cardMatchStruct{
			card:  cardType,
			count: 0,
		}
		for _, card := range cards {
			if card == cardType {
				thisCardTypeMatch.count++
			}
		}
		if thisCardTypeMatch.count > 0 {
			matches = append(matches, thisCardTypeMatch)
		}
	}
	return matches
}

type cardMatchStruct struct {
	card  cardStruct
	count int
}

type handOfCardsStruct struct {
	cards          [5]cardStruct
	numberOfJokers int
	bid            int
	rank           int
	strength       int
}

func (h *handOfCardsStruct) MarshalZerologObject(e *zerolog.Event) {
	e.Str("cards", string(h.cards[:]))
	e.Int("bid", h.bid)
	e.Int("rank", h.rank)
}

// getStrength will check for any valid hand types and return the highest strength (as int).
// param shouldCountJokers: if true, jokers will be counted as wildcards
func (h *handOfCardsStruct) getStrength(shouldCountJokers bool) int {
	// first, get all hand types
	// then, get the highest strength
	matches := testCardMatches(h.cards)

	if shouldCountJokers {
		for _, match := range matches {
			if match.card == 'J' {
				h.numberOfJokers = match.count
			}
		}
	}

	log.Trace().Msgf("Matches: %+v", matches)
	if h.isFiveOfAKind(matches) {
		log.Trace().Msg("Five of a kind")
		return strengthsMap["five of a kind"]
	}
	if h.isFourOfAKind(matches) {
		log.Trace().Msg("Four of a kind")
		return strengthsMap["four of a kind"]
	}
	if h.isFullHouse(matches) {
		log.Trace().Msg("Full house")
		return strengthsMap["full house"]
	}
	if h.isThreeOfAKind(matches) {
		log.Trace().Msg("Three of a kind")
		return strengthsMap["three of a kind"]
	}
	if h.isTwoPair(matches) {
		log.Trace().Msg("Two pair")
		return strengthsMap["two pair"]
	}
	if h.isOnePair(matches) {
		log.Trace().Msg("One pair")
		return strengthsMap["one pair"]
	}
	if h.isHighCard(matches) {
		log.Trace().Msg("High card")
		return strengthsMap["high card"]
	}
	log.Trace().Msg("No type found")
	return 0
}

func (h *handOfCardsStruct) isFiveOfAKind(matches []cardMatchStruct) bool {
	// all cards have the same value
	if h.numberOfJokers == 0 {
		return len(matches) == 1 && matches[0].count == 5
	}

	// if there are jokers, there should be 2 matches
	return len(matches) != 2
}

// one card is different
func (h *handOfCardsStruct) isFourOfAKind(matches []cardMatchStruct) bool {
	for _, match := range matches {
		if match.count == 4 {
			return len(matches) == 2
		}
	}

	return 5-h.numberOfJokers-matches[0].count == 1 || 5-h.numberOfJokers-matches[1].count == 1
}

// three cards have the same label, and the remaining two cards share a different label
func (h *handOfCardsStruct) isFullHouse(matches []cardMatchStruct) bool {
	var match1, match2, match1wJ, match2wJ bool
	for _, match := range matches {
		if match.count == 3 {
			match1 = true
		} else if match.count+h.numberOfJokers == 3 {
			match1wJ = true
		}
		if match.count == 2 {
			match2 = true
		} else if match.count+h.numberOfJokers == 2 {
			match2wJ = true
		}
	}
	return match1 && match2wJ || match1wJ && match2
}

func (h *handOfCardsStruct) isThreeOfAKind(matches []cardMatchStruct) bool {
	// three match, other two are different
	var shouldMatch3 cardMatchStruct
	var notMatch1, notMatch2 cardStruct
	for _, match := range matches {
		if match.count == 3 {
			shouldMatch3 = match
		} else if match.count+h.numberOfJokers == 3 {
			shouldMatch3 = match
		}
	}
	for _, card := range h.cards {
		if card != shouldMatch3.card {
			if notMatch1 == 0 {
				notMatch1 = card
			} else {
				notMatch2 = card
			}
		}
	}
	if notMatch1 == notMatch2 {
		return false
	}
	return shouldMatch3 != (cardMatchStruct{}) && notMatch1 != 0 && notMatch2 != 0
}

func (h *handOfCardsStruct) isTwoPair(matches []cardMatchStruct) bool {
	// two match, two match, one different
	var match1, match2 cardMatchStruct
	var notMatch1 cardStruct
	var usedJokers bool
	for _, match := range matches {
		if match.count == 2 {
			if match1 == (cardMatchStruct{}) {
				match1 = match
			} else {
				match2 = match
			}
		} else if match.count+h.numberOfJokers == 2 {
			if match1 == (cardMatchStruct{}) {
				match1 = match
				usedJokers = true
			} else {
				if !usedJokers {
					match2 = match
				}
			}
		}
	}
	for _, card := range h.cards {
		if card != match1.card && card != match2.card {
			notMatch1 = card
		}
	}
	return match1 != (cardMatchStruct{}) && match2 != (cardMatchStruct{}) && notMatch1 != 0
}

func (h *handOfCardsStruct) isOnePair(matches []cardMatchStruct) bool {
	// two cards match, the rest are unique
	var shouldMatch1 cardMatchStruct
	var notMatch1, notMatch2, notMatch3 cardStruct
	for _, match := range matches {
		if match.count == 2 {
			shouldMatch1 = match
		} else if match.count+h.numberOfJokers == 2 {
			shouldMatch1 = match
		}
	}
	for _, card := range h.cards {
		if card != shouldMatch1.card {
			if notMatch1 == 0 {
				notMatch1 = card
			} else if notMatch2 == 0 {
				notMatch2 = card
			} else {
				notMatch3 = card
			}
		}
	}
	if notMatch1 == notMatch2 || notMatch1 == notMatch3 || notMatch2 == notMatch3 {
		return false
	}
	return shouldMatch1 != (cardMatchStruct{}) && notMatch1 != 0 && notMatch2 != 0 && notMatch3 != 0
}

func (h *handOfCardsStruct) isHighCard(matches []cardMatchStruct) bool {
	// all cards are unique
	return len(matches) == 5
}
