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
		mySign := signs[1]

		score := calculateScore(opponentSign, mySign)
		if score < 0 {
			log.Fatal(wrongScore)
		}

		myTotalResult += score + valuesMapXYZ[mySign]
	}

	fmt.Println(myTotalResult)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func calculateScore(opponentInput, myInput string) int {
	switch opponentInput {
	case "A":
		switch myInput {
		case "X":
			return 3
		case "Y":
			return 6
		case "Z":
			return 0
		}
	case "B":
		switch myInput {
		case "X":
			return 0
		case "Y":
			return 3
		case "Z":
			return 6
		}
	case "C":
		switch myInput {
		case "X":
			return 6
		case "Y":
			return 0
		case "Z":
			return 3
		}
	}
	return -1
}
