package one

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Two() {
	file, err := os.Open("one/one.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	wordNums := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	result := 0
	for scanner.Scan() {
		text := scanner.Text()

		var accText string
		for _, char := range text {
			accText += string(char)
			for key, value := range wordNums {
				accText = strings.Replace(accText, key, strconv.Itoa(value), -1)
			}
		}

		var num string
		var lastNum rune
		foundFirstNum := false
		for _, char := range accText {
			if unicode.IsDigit(char) {
				lastNum = char
				if !foundFirstNum {
					num = string(char)
					foundFirstNum = true
				}
			}
		}

		num += string(lastNum)

		i, err := strconv.Atoi(num)

		if err != nil {
			fmt.Println(num, err)
			panic(err)
		}

		result += i
	}

	fmt.Println(result)
}
