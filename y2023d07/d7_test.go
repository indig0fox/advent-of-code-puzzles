package y2023d07

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestHands(t *testing.T) {
	tt := []struct {
		name             string
		hands            []*handOfCardsStruct
		expectedBids     []int
		expectedWinnings int
	}{
		{
			name: "Test 1",
			hands: []*handOfCardsStruct{
				{
					cards: [5]cardStruct{'3', '2', 'T', '3', 'K'},
					bid:   765,
				},
				{
					cards: [5]cardStruct{'T', '5', '5', 'J', '5'},
					bid:   684,
				},
				{
					cards: [5]cardStruct{'K', 'K', '6', '7', '7'},
					bid:   28,
				},
				{
					cards: [5]cardStruct{'K', 'T', 'J', 'J', 'T'},
					bid:   220,
				},
				{
					cards: [5]cardStruct{'Q', 'Q', 'Q', 'J', 'A'},
					bid:   483,
				},
			},
			expectedBids:     []int{765, 220, 28, 684, 483},
			expectedWinnings: 6440,
		},
		{
			name: "Five of a kind",
			hands: []*handOfCardsStruct{
				{
					cards: [5]cardStruct{'3', '3', '3', '3', '3'},
					bid:   765,
				},
				{
					cards: [5]cardStruct{'2', '2', '2', '2', '2'},
					bid:   684,
				},
				{
					cards: [5]cardStruct{'T', 'T', 'T', 'T', 'T'},
					bid:   28,
				},
			},
			expectedBids:     []int{684, 765, 28},
			expectedWinnings: 684*1 + 765*2 + 28*3,
		},
		{
			name: "Four of a kind",
			hands: []*handOfCardsStruct{
				{
					cards: [5]cardStruct{'3', '3', '3', '3', '2'},
					bid:   765,
				},
				{
					cards: [5]cardStruct{'2', '2', '2', '2', '3'},
					bid:   684,
				},
				{
					cards: [5]cardStruct{'T', 'T', 'T', 'T', '3'},
					bid:   28,
				},
			},
			expectedBids:     []int{684, 765, 28},
			expectedWinnings: 684*1 + 765*2 + 28*3,
		},
		{
			name: "Full house",
			hands: []*handOfCardsStruct{
				{
					cards: [5]cardStruct{'3', '3', '3', '2', '2'},
					bid:   765,
				},
				{
					cards: [5]cardStruct{'2', '2', '2', '3', '3'},
					bid:   684,
				},
				{
					cards: [5]cardStruct{'T', 'T', 'T', '3', '3'},
					bid:   28,
				},
			},
			expectedBids:     []int{684, 765, 28},
			expectedWinnings: 684*1 + 765*2 + 28*3,
		},
		{
			name: "High Card",
			hands: []*handOfCardsStruct{
				{
					cards: [5]cardStruct{'3', '2', 'T', '4', 'K'},
					bid:   765,
				},
				{
					cards: [5]cardStruct{'T', '5', '6', 'J', '8'},
					bid:   684,
				},
				{
					cards: [5]cardStruct{'K', 'Q', '6', '7', 'T'},
					bid:   28,
				},
			},
			expectedBids:     []int{765, 684, 28},
			expectedWinnings: 765*1 + 684*2 + 28*3,
		},
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:           os.Stdout,
		TimeFormat:    time.RFC3339,
		FieldsExclude: []string{"puzzle"},
	})
	// log.Logger = log.Level(zerolog.DebugLevel)
	log.Logger = log.Level(zerolog.TraceLevel)

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			hands := sortHands(test.hands, false)
			// ranks are index + 1
			bidOrder := []int{}
			for _, hand := range hands {
				bidOrder = append(bidOrder, hand.bid)
			}
			if !reflect.DeepEqual(bidOrder, test.expectedBids) {
				t.Errorf("Test %s failed. Expected bid order to be %v, got %v", test.name, test.expectedBids, bidOrder)
			}

			// apply rank based on sorting
			for i := range hands {
				hands[i].rank = i + 1
			}

			for _, hand := range hands {
				log.Debug().Object("hand", hand).Msg("Sorted hand")
			}

			// Part 1: Find total winnings
			if test.expectedWinnings != 0 {
				totalWinnings := 0
				for i, hand := range hands {
					totalWinnings += hand.bid * (i + 1)
				}

				if totalWinnings != test.expectedWinnings {
					t.Errorf("Test %s failed. Expected total winnings to be %d, got %d", test.name, test.expectedWinnings, totalWinnings)
				}
			}

		})
	}
}
