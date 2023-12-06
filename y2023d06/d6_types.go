package y2023d06

type resultStruct struct {
	Part1_ProductOfWaysToWinThreeRaces  int
	Part2_NumberOfWaysToWinVeryLongRace int
}

// raceStruct represents race details, such as the time limit and distance record
type raceStruct struct {
	Time                    int
	DistanceRecord          int
	WinningButtonPressTimes []int
}

// getWinningTimes will test different boat permutations to find
// the button press times which would beat the existing distance record
func (r *raceStruct) getWinningTimes() []int {
	// try integer button press times up to the race duration and get results that
	// are greater than the race's distance record
	winningBoats := []toyBoatStruct{}
	for i := 1; i < r.Time; i++ {
		boat := toyBoatStruct{
			TimeButtonHeld: i,
			Speed:          i,
		}
		if boat.calculateDistanceTravelledDuringRace(*r) > r.DistanceRecord {
			winningBoats = append(winningBoats, boat)
		}
	}

	r.WinningButtonPressTimes = []int{}
	for _, boat := range winningBoats {
		r.WinningButtonPressTimes = append(r.WinningButtonPressTimes, boat.TimeButtonHeld)
	}

	return r.WinningButtonPressTimes
}

// toyBoatStruct represents a toy boat with a button that can be held down to charge it up.
// The longer the button is held down, the faster the boat will go when released.
type toyBoatStruct struct {
	// TimeButtonHeld refers to the number of ms the charging button was held down
	TimeButtonHeld int
	// Speed refers to the speed of the boat in mm/s - this value is constant once the button is released
	Speed int
}

// calculateDistanceTravelledDuringRace will calculate the distance travelled by the boat
// based on its speed and the time for which the charging button was held down
func (b *toyBoatStruct) calculateDistanceTravelledDuringRace(r raceStruct) int {
	// subtract the time button was held from the available race time
	// multiply the result by the boat's speed to get the distance travelled in mm
	return (r.Time - b.TimeButtonHeld) * b.Speed
}
