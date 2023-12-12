package y2023d09

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestPart1(t *testing.T) {

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:           os.Stdout,
		TimeFormat:    time.RFC3339,
		FieldsExclude: []string{"puzzle"},
	})
	log.Logger = log.Level(zerolog.TraceLevel)

	testData := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
	wantSum := 114

	t.Run("TestPart1", func(t *testing.T) {
		records := parseValues(testData)
		for _, history := range records {
			history.CalculateDeltas()
			for _, line := range history.VisualizeDeltas() {
				fmt.Println(line)
			}
			history.PredictNextHistory()
			for _, line := range history.VisualizeDeltas() {
				fmt.Println(line)
			}
		}

		endSum := 0
		for _, history := range records {
			endSum += history.CurrentHistory[len(history.CurrentHistory)-1]
		}

		if endSum != wantSum {
			t.Errorf("Expected %v, got %v", wantSum, endSum)
		} else {
			t.Logf("Expected %v, got %v", wantSum, endSum)
		}
	})
}

func TestPart2(t *testing.T) {

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:           os.Stdout,
		TimeFormat:    time.RFC3339,
		FieldsExclude: []string{"puzzle"},
	})
	log.Logger = log.Level(zerolog.TraceLevel)

	testData := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
	wantSum := 2

	t.Run("TestPart1", func(t *testing.T) {
		records := parseValues(testData)
		for _, history := range records {
			history.CalculateDeltas()
			for _, line := range history.VisualizeDeltas() {
				fmt.Println(line)
			}
			history.PredictPreviousHistoryPt2()
			for _, line := range history.VisualizeDeltas() {
				fmt.Println(line)
			}
		}

		endSum := 0
		for _, history := range records {
			endSum += history.CurrentHistory[0]
		}

		if endSum != wantSum {
			t.Errorf("Expected %v, got %v", wantSum, endSum)
		} else {
			t.Logf("Expected %v, got %v", wantSum, endSum)
		}
	})
}
