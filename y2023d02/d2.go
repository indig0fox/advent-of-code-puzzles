package y2023d02

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
--- Day 2: Cube Conundrum ---
You're launched high into the atmosphere! The apex of your trajectory just barely reaches the surface of a large island floating in the sky. You gently land in a fluffy pile of leaves. It's quite cold, but you don't see much snow. An Elf runs over to greet you.

The Elf explains that you've arrived at Snow Island and apologizes for the lack of snow. He'll be happy to explain the situation, but it's a bit of a walk, so you have some time. They don't get many visitors up here; would you like to play a game in the meantime?

As you walk, the Elf shows you a small bag and some cubes which are either red, green, or blue. Each time you play this game, he will hide a secret number of cubes of each color in the bag, and your goal is to figure out information about the number of cubes.

To get information, once a bag has been loaded with cubes, the Elf will reach into the bag, grab a handful of random cubes, show them to you, and then put them back in the bag. He'll do this a few times per game.

You play several games and record the information from each game (your puzzle input). Each game is listed with its ID number (like the 11 in Game 11: ...) followed by a semicolon-separated list of subsets of cubes that were revealed from the bag (like 3 red, 5 green, 4 blue).

For example, the record of a few games might look like this:

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
In game 1, three sets of cubes are revealed from the bag (and then put back again). The first set is 3 blue cubes and 4 red cubes; the second set is 1 red cube, 2 green cubes, and 6 blue cubes; the third set is only 2 green cubes.

The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

In the example above, games 1, 2, and 5 would have been possible if the bag had been loaded with that configuration. However, game 3 would have been impossible because at one point the Elf showed you 20 red cubes at once; similarly, game 4 would also have been impossible because the Elf showed you 15 blue cubes at once. If you add up the IDs of the games that would have been possible, you get 8.

Determine which games would have been possible if the bag had been loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?
*/

type d2ResultStruct struct {
	Part1_NumberOfPossibleGames int
	Part2_CubePowerSum          int
}

var fLog *log.Logger

func Run(inputFile string, logFilePath string) d2ResultStruct {

	logFile, err := os.Create(logFilePath)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	fLog = log.New(multiWriter, "", log.LstdFlags|log.Lshortfile)

	data := getGameDataFromInput(inputFile)
	// dataJson, err := json.MarshalIndent(data[0], "", "  ")
	// if err != nil {
	// 	panic(err)
	// }
	// fLog.Printf("Data: %v\n", string(dataJson))

	for _, game := range data {
		game.gameIsPossible()

	}

	possibleCount := 0
	for _, game := range data {
		if game.IsPossible {
			possibleCount = possibleCount + game.ID
		}
	}

	for _, game := range data {
		minCubes := game.getMinimumCubeCount()
		power := minCubes.Red * minCubes.Green * minCubes.Blue
		game.CubePower = power
	}

	// for _, game := range data {
	// 	idStr := fmt.Sprintf("Game %03d", game.ID)
	// 	fLog.Printf("%s: isPossible: %-8t cubePower: %-8d", idStr, game.IsPossible, game.CubePower)
	// }

	cubePowerSum := 0
	for _, game := range data {
		cubePowerSum = cubePowerSum + game.CubePower
	}

	fLog.Println("--- Day 2: Cube Conundrum ---")
	fLog.Printf("Number of possible games: %d\n", possibleCount)
	fLog.Printf("Cube power sum: %d\n", cubePowerSum)

	return d2ResultStruct{
		Part1_NumberOfPossibleGames: possibleCount,
		Part2_CubePowerSum:          cubePowerSum,
	}
}

type gameData struct {
	ID         int
	GameSets   []gameSetArray
	IsPossible bool
	CubePower  int
}

func (g *gameData) gameIsPossible() bool {
	for _, gameSet := range g.GameSets {
		for _, cube := range gameSet {
			if cube.Count > 12 && cube.Color == "red" {
				g.IsPossible = false
				return false
			} else if cube.Count > 13 && cube.Color == "green" {
				g.IsPossible = false

				return false
			} else if cube.Count > 14 && cube.Color == "blue" {
				g.IsPossible = false
				return false
			}
		}
	}
	g.IsPossible = true
	return true
}

type cubeMinimumCountStruct struct {
	Red   int
	Green int
	Blue  int
}

func (g *gameData) getMinimumCubeCount() *cubeMinimumCountStruct {
	colors := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, gameSet := range g.GameSets {
		for _, cube := range gameSet {
			if cube.Count > colors[cube.Color] {
				colors[cube.Color] = cube.Count
			}
		}
	}

	return &cubeMinimumCountStruct{
		Red:   colors["red"],
		Green: colors["green"],
		Blue:  colors["blue"],
	}
}

type gameSetArray []cubeDataInstance

type cubeDataInstance struct {
	Count int
	Color string
}

func getGameDataFromInput(inputPath string) []*gameData {
	f, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var outputGameData []*gameData

	for scanner.Scan() {

		thisGame := gameData{
			ID:       0,
			GameSets: []gameSetArray{},
		}

		line := scanner.Text()
		// fLog.Println(line)

		reg := regexp.MustCompile(`Game (\d+): (.*)`)
		matches := reg.FindStringSubmatch(line)
		if len(matches) != 3 {
			panic("Invalid input")
		}

		thisGame.ID, err = strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}

		var thisGameSetsOfCubes []gameSetArray

		gameSetsRaw := strings.Split(matches[2], ";")
		for _, set := range gameSetsRaw {
			cubes := strings.Split(set, ",")
			reg := regexp.MustCompile(`(\d+) (\w+)`)

			thisGameSets := gameSetArray{}
			for _, cube := range cubes {
				match := reg.FindString(cube)
				// fLog.Println(match)

				// count
				count := 0
				count, err = strconv.Atoi(strings.Split(match, " ")[0])
				if err != nil {
					panic(err)
				}

				thisGameSets = append(thisGameSets, cubeDataInstance{
					Color: strings.Split(match, " ")[1],
					Count: count,
				})
			}
			thisGameSetsOfCubes = append(thisGameSetsOfCubes, thisGameSets)
		}

		thisGame.GameSets = thisGameSetsOfCubes
		outputGameData = append(outputGameData, &thisGame)
	}

	return outputGameData
}
