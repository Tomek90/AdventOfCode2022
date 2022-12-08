package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
		allCalories            []float64
		sumOfThreePigsCalories float64
	)

	for _, calorieSum := range elfCaloriesMap {
		allCalories = append(allCalories, calorieSum)
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(allCalories)))

	for i := 0; i < 3; i++ {
		sumOfThreePigsCalories += allCalories[i]
	}

	fmt.Println(sumOfThreePigsCalories)
}
