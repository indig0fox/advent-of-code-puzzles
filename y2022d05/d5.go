package y2022d05

import (
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type resultStruct struct {
	Part1_TopCratesOfEachStack                string
	Part2_TopCratesOfEachStackMultipleAtATime string
}

var (
	fLog *log.Logger
)

type crateObjectStruct struct {
	ID string
}

type crateMove struct {
	QuantityOfCrates int
	SourceStack      int
	TargetStack      int
}

type stackObject struct {
	ID     int
	Crates []crateObjectStruct
}

func (s *stackObject) getTopCrate() crateObjectStruct {
	return s.Crates[0]
}

func (s *stackObject) moveTopCrate(targetStack *stackObject) {
	targetStack.addCrateToTop(s.getTopCrate())
	s.Crates = s.Crates[1:]
}

func (s *stackObject) moveTopXCrates(targetStack *stackObject, x int) {
	// get the crates to move starting from the top
	cratesToMove := []crateObjectStruct{}
	for i := 0; i < x; i++ {
		cratesToMove = append(cratesToMove, s.Crates[i])
	}
	// add them to the target stack
	targetStack.addCratesToTop(cratesToMove)

	// remove them from the source stack
	for i := 0; i < x; i++ {
		s.Crates = s.Crates[1:]
	}
}

func (s *stackObject) addCrateToTop(crate crateObjectStruct) {
	newCrates := []crateObjectStruct{}
	newCrates = append(newCrates, crate)
	newCrates = append(newCrates, s.Crates...)
	s.Crates = newCrates
}

func (s *stackObject) addCratesToTop(crates []crateObjectStruct) {
	newCrates := []crateObjectStruct{}
	newCrates = append(newCrates, crates...)
	newCrates = append(newCrates, s.Crates...)
	s.Crates = newCrates
}

func (s *stackObject) showStack() string {
	outString := ""
	for _, crate := range s.Crates {
		outString += crate.ID
		outString += "\n"
	}
	return outString
}

func Run(inputPath string, logFilePath string) resultStruct {

	logFile, err := os.Create(logFilePath)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	fLog = log.New(multiWriter, "", log.LstdFlags|log.Lshortfile)

	data := getInput(inputPath)
	stacks := createStacks(9)
	var moves []crateMove
	stacks, moves = processInputLines(data, stacks)

	// Test moving crates
	// stacks[0].moveTopCrate(stacks[1])
	// fLog.Println("Stack 1 new top crate:", stacks[0].getTopCrate())
	// fLog.Println("Stack 2 new top crate:", stacks[1].getTopCrate())
	// fLog.Println("Stack 1:", stacks[0].showStack())
	// fLog.Println("Stack 2:", stacks[1].showStack())

	// Part 1: Find top crate in each stack after moveset

	// We have our stacks and our moves. Now we need to execute the moves.
	for _, move := range moves {
		// fLog.Println("Moving", move.QuantityOfCrates, "crates from stack", move.SourceStack, "to stack", move.TargetStack)
		for i := 0; i < move.QuantityOfCrates; i++ {
			for _, stack := range stacks {
				if stack.ID == move.SourceStack {
					for _, targetStack := range stacks {
						if targetStack.ID == move.TargetStack {
							stack.moveTopCrate(targetStack)
						}
					}
				}
			}
		}
		// showAllStacks(stacks)
	}

	// Get the top crates of each stack and concatenate them into a string
	part1OutString := ""
	for _, stack := range stacks {
		part1OutString += stack.getTopCrate().ID
	}

	//...............................................
	data = getInput(inputPath)
	stacks = createStacks(9)
	stacks, moves = processInputLines(data, stacks)

	// Part 2: Find top crate in each stack after moveset (moving multiple crates at a time!)
	for _, move := range moves {
		// fLog.Println("Moving", move.QuantityOfCrates, "crates from stack", move.SourceStack, "to stack", move.TargetStack)
		for _, stack := range stacks {
			if stack.ID == move.SourceStack {
				for _, targetStack := range stacks {
					if targetStack.ID == move.TargetStack {
						stack.moveTopXCrates(targetStack, move.QuantityOfCrates)
					}
				}
			}
		}
		// showAllStacks(stacks)
	}

	// Get the top crates of each stack and concatenate them into a string
	part2OutString := ""
	for _, stack := range stacks {
		part2OutString += stack.getTopCrate().ID
	}

	fLog.Println("--- 2022 Day 5: Ship Assignments ---")
	fLog.Println("Part 1: The top crates of each stack are:", part1OutString)
	fLog.Println("Part 2: The top crates of each stack (moving multiple at a time) are:", part2OutString)

	return resultStruct{
		Part1_TopCratesOfEachStack:                part1OutString,
		Part2_TopCratesOfEachStackMultipleAtATime: part2OutString,
	}
}

func getInput(inputPath string) []string {
	fLog.Println("Getting game data from input...")
	bytes, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	inputData := strings.Split(string(bytes), "\n")
	return inputData
}

func createStacks(stackCount int) []*stackObject {
	stacks := []*stackObject{}
	for i := 1; i < stackCount+1; i++ {
		stacks = append(stacks, &stackObject{
			ID:     i,
			Crates: []crateObjectStruct{},
		})
	}
	return stacks
}

func showAllStacks(stacks []*stackObject) {
	showString := "\n"
	for _, stack := range stacks {
		showString += "Stack " + strconv.Itoa(stack.ID) + ": "
		for i := 0; i < 8; i++ {
			if len(stack.Crates) > i {
				showString += stack.Crates[i].ID
				showString += "   "
			} else {
				showString += "    "
			}
		}
		showString += "\n"
	}
	fLog.Println(showString)
}

func processInputLines(input []string, stacks []*stackObject) ([]*stackObject, []crateMove) {
	moves := []crateMove{}

	// lines 1-8 are crates. the order indicates the stack they're in (starting at 1)

	// get crates from input (lines 1-8)
	regex := regexp.MustCompile(`\[(\w{1})\]\s|\s{3}`)
	for i := 0; i < 8; i++ {
		crateScan := regex.FindAllStringSubmatch(input[i], -1)
		// fLog.Println(crateScan)

		for j, crate := range crateScan {
			if crate[0] != "   " {
				for k, stack := range stacks {
					if stack.ID == j+1 {
						stacks[k].Crates = append(stack.Crates, crateObjectStruct{
							ID: crate[1],
						})
						// fLog.Println("Added crate", crate[1], "to stack", stack.ID)
					}
				}
			}
		}
	}

	// line 9 are the stack numbers

	// line 11-n are the moves to make
	for i := 10; i < len(input); i++ {
		regex := regexp.MustCompile(`\d{1,2}`)
		moveScan := regex.FindAllString(input[i], -1)

		// parse ints
		moveInts := [3]int{}
		for j, move := range moveScan {
			moveInt, err := strconv.Atoi(move)
			if err != nil {
				panic(err)
			}
			moveInts[j] = moveInt
		}

		moveItem := crateMove{
			QuantityOfCrates: moveInts[0],
			SourceStack:      moveInts[1],
			TargetStack:      moveInts[2],
		}

		// fLog.Printf("%+v", moveItem)

		moves = append(moves, moveItem)
	}

	return stacks, moves
}
