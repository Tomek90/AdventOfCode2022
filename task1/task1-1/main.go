package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("elves_calories.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	elfCaloriesMap := map[int]float64{}
	i := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			calories, err := strconv.ParseFloat(line, 64)
			if err != nil {
				log.Fatal(err)
			}
			elfCaloriesMap[i] += calories
		} else {
			i++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var (
		maxValue  float64
		elfNumber int
	)

	for elfNumberFromMap, calorieSum := range elfCaloriesMap {
		if maxValue < calorieSum {
			maxValue = calorieSum
			elfNumber = elfNumberFromMap + 1
		}
	}

	fmt.Println(elfNumber, maxValue)
}
