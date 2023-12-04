package y2022d02

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

/*
--- Day 2: Rock Paper Scissors ---
The Elves begin to set up camp on the beach. To decide whose tent gets to be closest to the snack storage, a giant Rock Paper Scissors tournament is already in progress.

Rock Paper Scissors is a game between two players. Each game contains many rounds; in each round, the players each simultaneously choose one of Rock, Paper, or Scissors using a hand shape. Then, a winner for that round is selected: Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock. If both players choose the same shape, the round instead ends in a draw.

Appreciative of your help yesterday, one Elf gives you an encrypted strategy guide (your puzzle input) that they say will be sure to help you win. "The first column is what your opponent is going to play: A for Rock, B for Paper, and C for Scissors. The second column--" Suddenly, the Elf is called away to help with someone's tent.

The second column, you reason, must be what you should play in response: X for Rock, Y for Paper, and Z for Scissors. Winning every time would be suspicious, so the responses must have been carefully chosen.

The winner of the whole tournament is the player with the highest score. Your total score is the sum of your scores for each round. The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

Since you can't be sure if the Elf is trying to help you or trick you, you should calculate the score you would get if you were to follow the strategy guide.

For example, suppose you were given the following strategy guide:

A Y
B X
C Z
This strategy guide predicts and recommends the following:

In the first round, your opponent will choose Rock (A), and you should choose Paper (Y). This ends in a win for you with a score of 8 (2 because you chose Paper + 6 because you won).
In the second round, your opponent will choose Paper (B), and you should choose Rock (X). This ends in a loss for you with a score of 1 (1 + 0).
The third round is a draw with both players choosing Scissors, giving you a score of 3 + 3 = 6.
In this example, if you were to follow the strategy guide, you would get a total score of 15 (8 + 1 + 6).

What would your total score be if everything goes exactly according to your strategy guide?

Your puzzle answer was 12458.

--- Part Two ---
The Elf finishes helping with the tent and sneaks back over to you. "Anyway, the second column says how the round needs to end: X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"

The total score is still calculated in the same way, but now you need to figure out what shape to choose so the round ends as indicated. The example above now goes like this:

In the first round, your opponent will choose Rock (A), and you need the round to end in a draw (Y), so you also choose Rock. This gives you a score of 1 + 3 = 4.
In the second round, your opponent will choose Paper (B), and you choose Rock so you lose (X) with a score of 1 + 0 = 1.
In the third round, you will defeat your opponent's Scissors with Rock for a score of 1 + 6 = 7.
Now that you're correctly decrypting the ultra top secret strategy guide, you would get a total score of 12.

Following the Elf's instructions for the second column, what would your total score be if everything goes exactly according to your strategy guide?

Your puzzle answer was 12683.
*/

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

	fLog.Println("--- Day 2: Rock Paper Scissors ---")
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
