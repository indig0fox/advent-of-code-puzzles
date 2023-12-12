package y2023d08

import (
	"io"
	"os"

	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var puzzleTitle = "--- Day 8: Haunted Wasteland ---"

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
	log.Logger = log.With().Timestamp().Str("puzzle", "2023_Day08").Logger()
	log.Logger = log.Level(zerolog.InfoLevel)
	// log.Logger = log.Level(zerolog.DebugLevel)
	// log.Logger = log.Level(zerolog.TraceLevel)

	startTime := time.Now()
	defer func() {
		log.Info().Msgf(
			"%s completed in %s",
			puzzleTitle,
			time.Since(startTime),
		)
	}()

	log.Debug().Msg("Running y2023d08.Run()...")
	log.Info().Msg(puzzleTitle)

	data := getInput(inputPath)
	instructions := parseInstructions(data)
	// parseInstructions(data)
	nodes := parseNodes(data)
	// parseNodes(data)

	// PART 1: Find total steps to get from AAA to ZZZ
	// now we need to go start at AAA and
	// follow through each instruction to navigate through the nodes
	// until we hit ZZZ
	// first, find AAA
	var currentNode *nodeStruct
	for _, node := range nodes {
		if node.ID == "AAA" {
			currentNode = node
			break
		}
	}

	totalInstructionsFollowed, instructionIndex := 0, 0
	for currentNode.ID != "ZZZ" {
		// find the next node to go to
		if instructionIndex >= len(instructions) {
			instructionIndex = 0
		}
		currentInstruction := instructions[instructionIndex]
		instructionIndex++
		// go to the next node
		currentNode = currentNode.Siblings[currentInstruction]
		totalInstructionsFollowed++
	}

	log.Info().Msgf("Part 1: Total instructions followed: %v", totalInstructionsFollowed)

	// PART 2: Starting on every node that ends with A,
	// run each instruction through until each node's path has hit a node ending with Z
	totalInstructionsFollowedPt2 := runPart2(instructions, nodes)

	log.Info().Msgf("Part 2: Total instructions followed starting at all [xxA] nodes: %v", totalInstructionsFollowedPt2)

	return resultStruct{
		Part1_TotalInstructionsFollowed:             totalInstructionsFollowed,
		Part2_TotalInstructionsFollowedForAllANodes: int(totalInstructionsFollowedPt2),
	}
}

func getInput(inputPath string) []string {
	inputData, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	input := string(inputData)
	inputLines := strings.Split(input, "\n")

	return inputLines
}

func parseInstructions(input []string) []instruction {
	instructions := []instruction{}

	instructionsRaw := strings.TrimSpace(input[0])
	for _, instructionRaw := range instructionsRaw {
		switch instructionRaw {
		case 'L':
			instructions = append(instructions, instruction(0))
		case 'R':
			instructions = append(instructions, instruction(1))
		default:
			panic("Unknown instruction")
		}

	}

	return instructions
}

func parseNodes(input []string) []*nodeStruct {
	nodes := []*nodeStruct{}

	nodesRaw := input[2:]

	// parse all nodes initially
	for _, line := range nodesRaw {
		nodes = append(nodes, &nodeStruct{ID: line[0:3]})
	}

	for _, line := range nodesRaw {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// find node to adjust
		node := &nodeStruct{}
		for _, testNode := range nodes {
			if testNode.ID == line[0:3] {
				node = testNode
			}
		}

		nodeSiblingsRaw := line[6:]
		// remove parens
		nodeSiblingsRaw = strings.ReplaceAll(nodeSiblingsRaw, "(", "")
		nodeSiblingsRaw = strings.ReplaceAll(nodeSiblingsRaw, ")", "")
		// split by comma
		nodeSiblingsSplit := strings.Split(nodeSiblingsRaw, ",")
		for i, nodeSibling := range nodeSiblingsSplit {
			// trim space
			nodeSibling = strings.TrimSpace(nodeSibling)
			// add address of sibling node to siblings array
			for j, testNode := range nodes {
				if testNode.ID == nodeSibling {
					node.Siblings[i] = nodes[j]
				}
			}
		}
		node.LastCharacter = rune(node.ID[len(node.ID)-1:][0])

		log.Trace().Msgf("Processed node: %+v", node)
	}

	return nodes
}

func runPart2(instructions []instruction, nodes []*nodeStruct) (instructionsFollowed int64) {
	// PART 2: Starting on every node that ends with A,
	// run each instruction through until all current nodes end with Z
	allCurrentNodesPt2 := []*nodeStruct{}
	for _, node := range nodes {
		if node.ID[len(node.ID)-1:] == "A" {
			allCurrentNodesPt2 = append(allCurrentNodesPt2, node)
		}
	}

	for _, node := range allCurrentNodesPt2 {
		node.getNextZ(instructions)
	}

	// find the lowest common multiple among all node step counts
	gcd := func(a, b int64) int64 {
		for b != 0 {
			t := b
			b = a % b
			a = t
		}
		return a
	}
	lcm := func(a, b int64) int64 {
		if a == 0 || b == 0 {
			return 0
		}
		var absA = a
		var absB = b
		if a < 0 {
			absA = -a
		}
		if b < 0 {
			absB = -b
		}
		var lcm int64 = absA * (absB / gcd(absA, absB))
		return lcm
	}

	var lowestCommonMultiple int64 = 1
	for _, node := range allCurrentNodesPt2 {
		lowestCommonMultiple = lcm(lowestCommonMultiple, node.StepsEndingInZ[len(node.StepsEndingInZ)-1])
	}

	return lowestCommonMultiple

	// var lengthInstructions = len(instructions)
	// var logInstructionsInterval = int64(math.Pow(1000, 3) / 2)
	// var instructionIndexPt2 int
	// var totalInstructionsFollowedPt2 int64

	// for {
	// 	if totalInstructionsFollowedPt2%logInstructionsInterval == 0 {
	// 		log.Trace().
	// 			Int64("current_total", totalInstructionsFollowedPt2).
	// 			Send()
	// 	}

	// 	currentInstruction := instructions[instructionIndexPt2]

	// 	navigateToNextNodeGroup(allCurrentNodesPt2, currentInstruction)
	// 	totalInstructionsFollowedPt2++

	// 	if allNodesEndWithZ(allCurrentNodesPt2) {
	// 		return totalInstructionsFollowedPt2
	// 	}

	// 	instructionIndexPt2++
	// 	if instructionIndexPt2 >= lengthInstructions {
	// 		instructionIndexPt2 = 0
	// 	}
	// }
}
