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
		sameCharactersArr []rune
		prioritiesSum     int
	)

	f, err := os.Open("items.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rucksackItems := scanner.Text()
		halfIndex := len(rucksackItems) / 2

		rucksackOneItems := rucksackItems[:halfIndex]
		rucksackTwoItems := rucksackItems[halfIndex:]

		rucksackOneUniqueItemsMap := map[rune]bool{}

		for _, char := range []rune(rucksackOneItems) {
			rucksackTwoIndex := strings.Index(rucksackTwoItems, string(char))
			if rucksackTwoIndex > -1 {
				if _, ok := rucksackOneUniqueItemsMap[char]; ok {
					continue
				}

				sameCharactersArr = append(sameCharactersArr, char)
				rucksackOneUniqueItemsMap[char] = true
			}
		}
	}

	for _, char := range sameCharactersArr {
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
