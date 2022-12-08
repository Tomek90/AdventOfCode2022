package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		allLines           []string
		lastCrateLine      = 7
		firstMoveStatement = 10
	)

	f, err := os.Open("crates.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		allLines = append(allLines, scanner.Text())
	}

	crates := getCrates(allLines[:lastCrateLine+1])
	crates, err = executeMovesOnCrates(crates, allLines[firstMoveStatement:])
	if err != nil {
		log.Fatal(err)
	}

	topCrates := getTopCrates(crates)
	fmt.Print(topCrates)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getCrates(inputLines []string) (crates [][]string) {
	for i := 0; i <= len(inputLines); i++ {
		crates = append(crates, []string{})
	}

	for i := len(inputLines) - 1; i >= 0; i-- {
		cratesInx := 0
		inputLine := inputLines[i]
		for j := 1; j <= len(inputLine); j += 4 {
			if inputLine[j:j+1] != " " {
				crates[cratesInx] = append(crates[cratesInx], inputLine[j:j+1])
			}
			cratesInx++
		}
	}

	return
}

func executeMovesOnCrates(crates [][]string, moveStatements []string) (newCrates [][]string, err error) {
	for _, moveStatement := range moveStatements {
		qtyMoved, from, to, err := getCoordinatesFromString(moveStatement)
		if err != nil {
			return nil, err
		}

		crates[to] = append(crates[to], crates[from][len(crates[from])-qtyMoved:len(crates[from])]...)

		crates[from] = crates[from][:len(crates[from])-qtyMoved]
	}
	newCrates = crates

	return
}

func getCoordinatesFromString(moveStatement string) (int, int, int, error) {
	words := strings.Split(moveStatement, " ")

	qtyMoved, err := strconv.Atoi(words[1])
	if err != nil {
		return 0, 0, 0, err
	}

	from, err := strconv.Atoi(words[3])
	if err != nil {
		return 0, 0, 0, err
	}

	to, err := strconv.Atoi(words[5])
	if err != nil {
		return 0, 0, 0, err
	}

	return qtyMoved, from - 1, to - 1, nil
}

func getTopCrates(crates [][]string) (topCrates string) {
	for _, crate := range crates {
		topCrates += crate[len(crate)-1]
	}

	return
}
