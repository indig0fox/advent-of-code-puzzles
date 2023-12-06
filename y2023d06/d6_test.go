package y2023d06

import (
	"reflect"
	"testing"
)

func TestRaces(t *testing.T) {

	tt := []struct {
		name                        string
		race                        raceStruct
		rangeButtonTimesThatWillWin []int
	}{
		{
			name:                        "Test 1",
			race:                        raceStruct{Time: 7, DistanceRecord: 9},
			rangeButtonTimesThatWillWin: []int{2, 5},
		},
		{
			name:                        "Test 2",
			race:                        raceStruct{Time: 15, DistanceRecord: 40},
			rangeButtonTimesThatWillWin: []int{4, 11},
		},
		{
			name:                        "Test 3",
			race:                        raceStruct{Time: 30, DistanceRecord: 200},
			rangeButtonTimesThatWillWin: []int{11, 19},
		},
	}

	for _, test := range tt {

		t.Run(test.name, func(t *testing.T) {
			// test for the winning range of button times that will beat the distance record
			// try integer button press times up to the race duration and get results that
			// are greater than the race's distance record
			winningBoats := []toyBoatStruct{}
			for i := 1; i < test.race.Time; i++ {
				boat := toyBoatStruct{
					TimeButtonHeld: i,
					Speed:          i,
				}
				if boat.calculateDistanceTravelledDuringRace(test.race) > test.race.DistanceRecord {
					winningBoats = append(winningBoats, boat)
				}
			}

			winningTimes := []int{}
			for _, boat := range winningBoats {
				winningTimes = append(winningTimes, boat.TimeButtonHeld)
			}

			winningRange := []int{winningTimes[0], winningTimes[len(winningTimes)-1]}

			if !reflect.DeepEqual(test.rangeButtonTimesThatWillWin, winningRange) {
				t.Errorf("Test %s failed. Expected range of winning button times to be %v, got %v", test.name, test.rangeButtonTimesThatWillWin, winningTimes)
			}
		})
	}
}
