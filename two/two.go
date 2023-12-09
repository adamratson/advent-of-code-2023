package two

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Two() {
	file, err := os.Open("two/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		text := strings.Replace(scanner.Text(), "Game ", "", -1)
		gameIdAndParams := strings.Split(text, ": ")

		gameParams := gameIdAndParams[1]

		revelations := strings.Split(gameParams, "; ")
		gamePower := powerOfCubes(revelations)

		sum += gamePower
	}

	fmt.Println(sum)
}

func powerOfCubes(revelations []string) int {
	power := 1
	for _, v := range findMinCubes(revelations) {
		power = power * v
	}
	return power
}

func findMinCubes(revelations []string) map[string]int {
	minCubes := map[string]int{}
	for _, reachRevelation := range revelations {
		reachRevelationSplit := strings.Split(reachRevelation, ", ")
		for _, colourRevelation := range reachRevelationSplit {
			colourRevelationSplit := strings.Split(colourRevelation, " ")
			num, err := strconv.Atoi(colourRevelationSplit[0])

			if err != nil {
				fmt.Println(num, err)
				panic(err)
			}

			colour := colourRevelationSplit[1]

			currentMin, present := minCubes[colour]

			if num > currentMin || !present {
				minCubes[colour] = num
			}
		}
	}

	return minCubes
}
