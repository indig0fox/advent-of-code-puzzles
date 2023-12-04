package y2022d03

import (
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

/*
--- Day 3: Rucksack Reorganization ---
One Elf has the important job of loading all of the rucksacks with supplies for the jungle journey. Unfortunately, that Elf didn't quite follow the packing instructions, and so a few items now need to be rearranged.

Each rucksack has two large compartments. All items of a given type are meant to go into exactly one of the two compartments. The Elf that did the packing failed to follow this rule for exactly one item type per rucksack.

The Elves have made a list of all of the items currently in each rucksack (your puzzle input), but they need your help finding the errors. Every item type is identified by a single lowercase or uppercase letter (that is, a and A refer to different types of items).

The list of items for each rucksack is given as characters all on a single line. A given rucksack always has the same number of items in each of its two compartments, so the first half of the characters represent items in the first compartment, while the second half of the characters represent items in the second compartment.

For example, suppose you have the following list of contents from six rucksacks:

vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
The first rucksack contains the items vJrwpWtwJgWrhcsFMMfFFhFp, which means its first compartment contains the items vJrwpWtwJgWr, while the second compartment contains the items hcsFMMfFFhFp. The only item type that appears in both compartments is lowercase p.
The second rucksack's compartments contain jqHRNqRjqzjGDLGL and rsFMfFZSrLrFZsSL. The only item type that appears in both compartments is uppercase L.
The third rucksack's compartments contain PmmdzqPrV and vPwwTWBwg; the only common item type is uppercase P.
The fourth rucksack's compartments only share item type v.
The fifth rucksack's compartments only share item type t.
The sixth rucksack's compartments only share item type s.
To help prioritize item rearrangement, every item type can be converted to a priority:

Lowercase item types a through z have priorities 1 through 26.
Uppercase item types A through Z have priorities 27 through 52.
In the above example, the priority of the item type that appears in both compartments of each rucksack is 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s); the sum of these is 157.

Find the item type that appears in both compartments of each rucksack. What is the sum of the priorities of those item types?

Your puzzle answer was 8072.

--- Part Two ---
As you finish identifying the misplaced items, the Elves come to you with another issue.

For safety, the Elves are divided into groups of three. Every Elf carries a badge that identifies their group. For efficiency, within each group of three Elves, the badge is the only item type carried by all three Elves. That is, if a group's badge is item type B, then all three Elves will have item type B somewhere in their rucksack, and at most two of the Elves will be carrying any other item type.

The problem is that someone forgot to put this year's updated authenticity sticker on the badges. All of the badges need to be pulled out of the rucksacks so the new authenticity stickers can be attached.

Additionally, nobody wrote down which item type corresponds to each group's badges. The only way to tell which item type is the right one is by finding the one item type that is common between all three Elves in each group.

Every set of three lines in your list corresponds to a single group, but each group can have a different badge item type. So, in the above example, the first group's rucksacks are the first three lines:

vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
And the second group's rucksacks are the next three lines:

wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
In the first group, the only item type that appears in all three rucksacks is lowercase r; this must be their badges. In the second group, their badge item type must be Z.

Priorities for these items must still be found to organize the sticker attachment efforts: here, they are 18 (r) for the first group and 52 (Z) for the second group. The sum of these is 70.

Find the item type that corresponds to the badges of each three-Elf group. What is the sum of the priorities of those item types?

Your puzzle answer was 2567.
*/

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

	fLog.Println("--- Day 3: Rucksack Reorganization ---")
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
