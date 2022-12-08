package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func main() {
	var (
		allRucksacks  []string
		allGroupChars []rune
		prioritiesSum int
	)

	f, err := os.Open("items.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		allRucksacks = append(allRucksacks, scanner.Text())
	}

	for i := 0; i < len(allRucksacks); i += 3 {
		for _, char := range []rune(allRucksacks[i]) {
			if strings.Index(allRucksacks[i+1], string(char)) > -1 && strings.Index(allRucksacks[i+2], string(char)) > -1 {
				allGroupChars = append(allGroupChars, char)

				break
			}
		}
	}

	for _, char := range allGroupChars {
		charPriority := strings.Index(alphabet, strings.ToLower(string(char))) + 1
		if unicode.IsUpper(char) {
			charPriority += 26
		}

		prioritiesSum += charPriority
	}

	fmt.Println(prioritiesSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
