package y2022d03

import (
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

type resultStruct struct {
	Part1_SharedItemPrioritySum int
	Part2_BadgePrioritySum      int
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

	lines := getInput(inputPath)

	// PART 1: Rucksack Reorganization
	sharedItems := []rune{}
	sharedItemPriorities := []int{}
	for _, line := range lines {
		pocket1 := []rune(line[:len(line)/2])
		pocket2 := []rune(line[len(line)/2:])
		sharedItem, found := determineSharedItemInRucksack(pocket1, pocket2)
		if !found {
			panic("No shared item found")
		} else {
			sharedItems = append(sharedItems, sharedItem)
		}
	}
	for _, item := range sharedItems {
		sharedItemPriorities = append(sharedItemPriorities, getPriorityOfItem(item))
	}
	sharedItemPrioritiesSum := 0
	for _, priority := range sharedItemPriorities {
		sharedItemPrioritiesSum += priority
	}

	// PART 2: Badge Search
	for (len(lines) % 3) != 0 {
		lines = append(lines, "")
	}
	elfGroups := [][3]string{}
	for i := 0; i < len(lines); i += 3 {
		elfGroups = append(elfGroups, [3]string{lines[i], lines[i+1], lines[i+2]})
	}
	sharedItemsInElfGroups := []string{}
	for _, elfGroup := range elfGroups {
		sharedItemsInElfGroups = append(sharedItemsInElfGroups, getSharedItemInElfGroup(elfGroup))
	}
	sharedBadgePrioritiesSum := 0
	for _, item := range sharedItemsInElfGroups {
		sharedBadgePrioritiesSum += getPriorityOfItem(rune(item[0]))
	}

	fLog.Println("--- 2022 Day 3: Rucksack Reorganization ---")
	fLog.Println("Part 1: The sum of the priorities of the shared items is:", sharedItemPrioritiesSum)
	fLog.Println("Part 2: The sum of the priorities of the shared badge items is:", sharedBadgePrioritiesSum)

	return resultStruct{
		Part1_SharedItemPrioritySum: sharedItemPrioritiesSum,
		Part2_BadgePrioritySum:      sharedBadgePrioritiesSum,
	}
}

func getInput(inputPath string) []string {
	bytes, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\n")
	return lines
}

func countItemsInSlice[T any](slice []T, item T) int {
	count := 0
	for _, sliceItem := range slice {
		if reflect.DeepEqual(sliceItem, item) {
			count++
		}
	}
	return count
}

func determineSharedItemInRucksack(itemGroupsToCheck ...[]rune) (rune, bool) {

	checkFor := func(item rune, itemGroupsToCheck ...[]rune) bool {
		groupsHave := 0
		for _, itemGroup := range itemGroupsToCheck {
			if countItemsInSlice(itemGroup, item) > 0 {
				groupsHave++
			}
		}
		return groupsHave == len(itemGroupsToCheck)
	}

	for _, itemGroup := range itemGroupsToCheck {
		for _, item := range itemGroup {
			if checkFor(item, itemGroupsToCheck...) {
				return item, true
			}
		}
	}
	return ' ', false
}

func getPriorityOfItem(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1)
	}
	if item >= 'A' && item <= 'Z' {
		return int(item - 'A' + 27)
	}
	panic("Invalid item")
}

func getSharedItemInElfGroup(elfGroup [3]string) string {
	sharedItem, found := determineSharedItemInRucksack([]rune(elfGroup[0]), []rune(elfGroup[1]), []rune(elfGroup[2]))
	if !found {
		panic("No shared item found")
	}
	return string(sharedItem)
}
