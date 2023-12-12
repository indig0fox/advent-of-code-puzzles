package y2023d08

type resultStruct struct {
	Part1_TotalInstructionsFollowed             int
	Part2_TotalInstructionsFollowedForAllANodes int
}

type instruction int

type nodeStruct struct {
	ID             string
	LastCharacter  rune
	StepsEndingInZ []int64
	Siblings       [2]*nodeStruct
}

func (n *nodeStruct) getNextZ(instructions []instruction) int64 {
	var stepsTaken int64 = 0
	if n.StepsEndingInZ != nil {
		stepsTaken = n.StepsEndingInZ[len(n.StepsEndingInZ)-1]
	}
	var p *nodeStruct = n
	var instructionsIndex int = 0
	var instructionsLength int = len(instructions)
	for !p.endsWithZ() {
		if instructionsIndex == instructionsLength {
			instructionsIndex = 0
		}
		p = p.Siblings[instructions[instructionsIndex]]
		instructionsIndex++
		stepsTaken++
	}
	n.StepsEndingInZ = append(n.StepsEndingInZ, stepsTaken)
	return stepsTaken
}

func (n nodeStruct) endsWithZ() bool {
	return n.LastCharacter == 'Z'
}
