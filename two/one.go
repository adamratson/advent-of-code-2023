package two

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func One() {
	file, err := os.Open("two/input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxOfColours := map[string]int{"red": 12, "green": 13, "blue": 14}

	var successfulGames []string

	for scanner.Scan() {
		text := strings.Replace(scanner.Text(), "Game ", "", -1)
		gameIdAndParams := strings.Split(text, ": ")

		gameId := gameIdAndParams[0]
		gameParams := gameIdAndParams[1]

		revelations := strings.Split(gameParams, "; ")
		if revelationsAreValid(maxOfColours, revelations) {
			successfulGames = append(successfulGames, gameId)
		}
	}

	sum := 0
	for _, game := range successfulGames {
		gameInt, _ := strconv.Atoi(game)
		sum += gameInt
	}

	fmt.Println(sum)
}

func revelationsAreValid(maxOfColours map[string]int, revelations []string) bool {
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

			if num > maxOfColours[colour] {
				return false
			}
		}
	}

	return true
}
