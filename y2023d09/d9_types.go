package y2023d09

import (
	"fmt"
)

type resultStruct struct {
	Part1_HistoryExtrapolatedForwardSum  int
	Part2_HistoryExtrapolatedBackwardSum int
}

type recordStruct struct {
	OriginalHistory []int
	CurrentHistory  []int
	Deltas          [][]int
}

func (r *recordStruct) AddHistory(value int) {
	r.OriginalHistory = append(r.OriginalHistory, value)
	r.CurrentHistory = append(r.CurrentHistory, value)
}

func allIntsAreZero(data []int) bool {
	for _, value := range data {
		if value != 0 {
			return false
		}
	}
	return true
}

func (r *recordStruct) CalculateDeltas() {
	r.Deltas = [][]int{r.CurrentHistory}

	i := 0
	for !allIntsAreZero(r.Deltas[i]) {
		newDelta := []int{}
		for j := range r.Deltas[i] {
			if j == 0 {
				continue
			}
			newDelta = append(newDelta, r.Deltas[i][j]-r.Deltas[i][j-1])
		}
		r.Deltas = append(r.Deltas, newDelta)
		i++
	}
}

func (r *recordStruct) VisualizeDeltas() []string {
	var output []string

	var source [][]int = r.Deltas

	// for each row
	for sourceIndex, intArr := range source {
		var line string
		line += fmt.Sprintf("%-*s", 2*sourceIndex, "")
		// for each number, calculate the distance to the next number
		for _, value := range intArr {
			line += fmt.Sprintf("%-*d", 4, value)
		}
		output = append(output, line)
	}

	return output
}

func (r *recordStruct) PredictNextHistory() {
	// add 0 to the end of each row
	for i := range r.Deltas {
		r.Deltas[i] = append(r.Deltas[i], 0)
	}

	// calculate the last value in each row above (indexed before) the last row
	for i := len(r.Deltas) - 2; i >= 0; i-- {
		// moving from bottom to top (right to left), calculate the next value
		// by summing the previous value in this row with the last value in the previous row
		lastTwoValsCurRow := r.Deltas[i][len(r.Deltas[i])-2:]
		lastValPrevRow := r.Deltas[i+1][len(r.Deltas[i+1])-1]
		// set last val cur row
		r.Deltas[i][len(r.Deltas[i])-1] = lastTwoValsCurRow[0] + lastValPrevRow
	}

	// set current history
	r.CurrentHistory = r.Deltas[0]
	// calculate new deltas
	r.CalculateDeltas()
}

func (r *recordStruct) PredictPreviousHistoryPt2() {
	// add 0 to the beginning of each row
	for i := range r.Deltas {
		r.Deltas[i] = append([]int{0}, r.Deltas[i]...)
	}

	// calculate the first value in each row above (indexed before) the last row
	for i := len(r.Deltas) - 2; i >= 0; i-- {
		// moving from bottom to top (right to left), calculate the next value
		// by summing the previous value in this row with the last value in the previous row
		firstTwoValsCurRow := r.Deltas[i][:2]
		firstValPrevRow := r.Deltas[i+1][0]
		// set last val cur row
		r.Deltas[i][0] = firstTwoValsCurRow[1] - firstValPrevRow
	}

	// set current history
	r.CurrentHistory = r.Deltas[0]
	// calculate new deltas
	r.CalculateDeltas()
}
