package y2022d02

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type resultStruct struct {
	Part1_Score              int
	Part2_StrategyGuideScore int
}

var fLog *log.Logger

func Run(inputPath string, logFilePath string) resultStruct {

	logFile, err := os.Create(logFilePath)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	fLog = log.New(multiWriter, "", log.LstdFlags|log.Lshortfile)

	data := getInputData(inputPath)

	player1g1 := playerStruct{
		name:  "Player 1",
		score: 0,
	}

	player2g1 := playerStruct{
		name:  "Player 2",
		score: 0,
	}

	player1g2 := playerStruct{
		name:  "Player 1",
		score: 0,
	}

	player2g2 := playerStruct{
		name:  "Player 2",
		score: 0,
	}

	// PART 1
	for _, line := range data {
		// fLog.Printf("Line: %v\n", line)
		lineMoves := strings.Split(line, " ")
		player1g1.currentMove = string(lineMoves[0])
		player2g1.currentMove = string(lineMoves[1])

		playRockPaperScissors(&player1g1, &player2g1)
	}

	// PART 2
	for _, line := range data {
		// fLog.Printf("Line: %v\n", line)
		lineMoves := strings.Split(line, " ")
		player1g2.currentMove = string(lineMoves[0])
		player2g2.currentMove = string(lineMoves[1])

		player2g2.currentMove = player2g2.getStrategyGuideMove(player1g2.currentMove)
		playRockPaperScissors(&player1g2, &player2g2)
	}

	fLog.Println("--- 2022 Day 2: Rock Paper Scissors ---")
	fLog.Println("Part 1: Player 2 Game 1's score after", len(data), "rounds:", player2g1.score)
	fLog.Println("Part 2: Player 2 Game 2's score after", len(data), "rounds (using strategy guide!):", player2g2.score)

	return resultStruct{
		Part1_Score:              player2g1.score,
		Part2_StrategyGuideScore: player2g2.score,
	}
}

func getInputData(inputPath string) []string {

	f, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	data := []string{}

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}

type rpsMove struct {
	name                  string
	pointsEarnedByPlaying int
	versusRock            int
	versusScissors        int
	versusPaper           int
}

var moveRock rpsMove = rpsMove{
	name:                  "Rock",
	pointsEarnedByPlaying: 1,
	versusRock:            3,
	versusScissors:        6,
	versusPaper:           0,
}

var movePaper rpsMove = rpsMove{
	name:                  "Paper",
	pointsEarnedByPlaying: 2,
	versusRock:            6,
	versusScissors:        0,
	versusPaper:           3,
}

var moveScissors rpsMove = rpsMove{
	name:                  "Scissors",
	pointsEarnedByPlaying: 3,
	versusRock:            0,
	versusScissors:        3,
	versusPaper:           6,
}

var moveSet map[string]*rpsMove = map[string]*rpsMove{
	"A": &moveRock,
	"X": &moveRock,
	"B": &movePaper,
	"Y": &movePaper,
	"C": &moveScissors,
	"Z": &moveScissors,
}

type playerStruct struct {
	name        string
	moves       []string
	currentMove string
	score       int
}

var strategyGuide map[string]string = map[string]string{
	"X": "Lose",
	"Y": "Draw",
	"Z": "Win",
}

func (p *playerStruct) getStrategyGuideMove(opponentMove string) string {
	opponentMoveActual := moveSet[opponentMove]
	if strategyGuide[p.currentMove] == "Win" {
		if opponentMoveActual.name == "Rock" {
			return "Y"
		} else if opponentMoveActual.name == "Paper" {
			return "Z"
		} else if opponentMoveActual.name == "Scissors" {
			return "X"
		}
	} else if strategyGuide[p.currentMove] == "Lose" {
		if opponentMoveActual.name == "Rock" {
			return "Z"
		} else if opponentMoveActual.name == "Paper" {
			return "X"
		} else if opponentMoveActual.name == "Scissors" {
			return "Y"
		}
	} else if strategyGuide[p.currentMove] == "Draw" {
		if opponentMoveActual.name == "Rock" {
			return "X"
		} else if opponentMoveActual.name == "Paper" {
			return "Y"
		} else if opponentMoveActual.name == "Scissors" {
			return "Z"
		}
	}
	return ""
}

func playRockPaperScissors(player1 *playerStruct, player2 *playerStruct) {
	player1Move := moveSet[player1.currentMove]
	player2Move := moveSet[player2.currentMove]

	if player1Move == nil || player2Move == nil {
		panic("Invalid move")
	}

	if player1Move.name == "Rock" {
		player2.score += player2Move.versusRock + player2Move.pointsEarnedByPlaying
	} else if player1Move.name == "Paper" {
		player2.score += player2Move.versusPaper + player2Move.pointsEarnedByPlaying
	} else if player1Move.name == "Scissors" {
		player2.score += player2Move.versusScissors + player2Move.pointsEarnedByPlaying
	}

	if player2Move.name == "Rock" {
		player1.score += player1Move.versusRock + player1Move.pointsEarnedByPlaying
	}
	if player2Move.name == "Paper" {
		player1.score += player1Move.versusPaper + player1Move.pointsEarnedByPlaying
	}
	if player2Move.name == "Scissors" {
		player1.score += player1Move.versusScissors + player1Move.pointsEarnedByPlaying
	}

	player1.moves = append(player1.moves, player1.currentMove)
	player2.moves = append(player2.moves, player2.currentMove)

	player1.currentMove = ""
	player2.currentMove = ""
}
