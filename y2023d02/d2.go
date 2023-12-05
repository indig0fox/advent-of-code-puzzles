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

	fLog.Println("--- 2023 Day 2: Cube Conundrum ---")
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
