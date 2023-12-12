package y2023d08

import (
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestPart2(t *testing.T) {

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:           os.Stdout,
		TimeFormat:    time.RFC3339,
		FieldsExclude: []string{"puzzle"},
	})
	log.Logger = log.Level(zerolog.TraceLevel)

	tt := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "example",
			input: []string{
				"LR",
				"",
				"11A = (11B, XXX)",
				"11B = (XXX, 11Z)",
				"11Z = (11B, XXX)",
				"22A = (22B, XXX)",
				"22B = (22C, 22C)",
				"22C = (22Z, 22Z)",
				"22Z = (22B, 22B)",
				"XXX = (XXX, XXX)",
			},
			expected: 6,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			instructions := parseInstructions(tc.input)
			// parseInstructions(data)
			nodes := parseNodes(tc.input)

			totalInstructionsFollowedPt2 := runPart2(instructions, nodes)

			if totalInstructionsFollowedPt2 != int64(tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, totalInstructionsFollowedPt2)
			} else {
				t.Logf("Expected %v, got %v", tc.expected, totalInstructionsFollowedPt2)
			}
		})
	}
}
