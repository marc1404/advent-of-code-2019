package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	adapters := readLinesAsAdapters("./day10/input.txt")

	log.Println("Day 10 Part 01")
	testPartOne()
	partOne(adapters)

	log.Println()

	log.Println("Day 10 Part 02")
	partTwo()
}

func testPartOne() {
	adapters := []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}

	printJoltageRatingAndDifferences(adapters)

	adapters = []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}

	printJoltageRatingAndDifferences(adapters)
}

func partOne(adapters []int) {
	_, differenceToCount := getJoltageRatingAndDifferences(adapters)
	product := differenceToCount[1] * differenceToCount[3]

	log.Println("Product of 1-jolt and 3-jolt differences:", product)
}

func partTwo() {

}

func getJoltageRatingAndDifferences(adapters []int) (deviceJoltageRating int, differenceToCount map[int]int) {
	sortAscending(adapters)

	deviceJoltageRating = determineDeviceJoltageRating(adapters)
	differenceToCount = collectJoltageDifferences(adapters, deviceJoltageRating, make(map[int]int))

	return deviceJoltageRating, differenceToCount
}

func printJoltageRatingAndDifferences(adapters []int) {
	deviceJoltageRating, differenceToCount := getJoltageRatingAndDifferences(adapters)

	log.Println("Device joltage rating:", deviceJoltageRating)
	log.Println("Joltage differences:", differenceToCount)
}

func sortAscending(values []int) {
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
}

func determineDeviceJoltageRating(adapters []int) int {
	highestAdapter := 0

	for _, adapter := range adapters {
		if adapter > highestAdapter {
			highestAdapter = adapter
		}
	}

	return highestAdapter + 3
}

func collectJoltageDifferences(adapters []int, targetJoltage int, differenceToCount map[int]int) map[int]int {
	adapters = append(adapters, targetJoltage)
	sourceJoltage := 0

	for _, adapter := range adapters {
		difference := adapter - sourceJoltage
		differenceToCount[difference]++
		sourceJoltage = adapter
	}

	return differenceToCount
}

func readLinesAsAdapters(filePath string) []int {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	adapters := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		adapter, _ := strconv.Atoi(line)
		adapters = append(adapters, adapter)
	}

	return adapters
}
