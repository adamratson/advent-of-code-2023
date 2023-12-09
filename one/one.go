package one

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func One() {
	file, err := os.Open("one/one.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result = 0
	for scanner.Scan() {
		var num string
		var lastNum rune
		var foundFirstNum = false
		for _, char := range scanner.Text() {
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
			panic(err)
		}

		result += i
	}

	fmt.Println(result)
}
