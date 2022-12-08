package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var wrongScore = errors.New("a wrong score calculated")

var valuesMapXYZ = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func main() {
	var myTotalResult int

	f, err := os.Open("paper_rock_scissors.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		signs := strings.Split(scanner.Text(), " ")
		opponentSign := signs[0]
		myChar := signs[1]

		score, signUsed := calculateScore(opponentSign, myChar)
		if score < 0 {
			log.Fatal(wrongScore)
		}

		myTotalResult += score + valuesMapXYZ[signUsed]
	}

	fmt.Println(myTotalResult)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func calculateScore(opponentInput, myChar string) (int, string) {
	switch myChar {
	case "X":
		switch opponentInput {
		case "A":
			return 0, "Z"
		case "B":
			return 0, "X"
		case "C":
			return 0, "Y"
		}
	case "Y":
		switch opponentInput {
		case "A":
			return 3, "X"
		case "B":
			return 3, "Y"
		case "C":
			return 3, "Z"
		}
	case "Z":
		switch opponentInput {
		case "A":
			return 6, "Y"
		case "B":
			return 6, "Z"
		case "C":
			return 6, "X"
		}
	}
	return -1, ""
}
