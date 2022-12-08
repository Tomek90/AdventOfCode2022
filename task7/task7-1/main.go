package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const maxSize = 100000

func main() {
	var allLines []string

	f, err := os.Open("directories.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		allLines = append(allLines, scanner.Text())
	}

	directoriesMap := map[int]int{}
	directorySeq := 0

	for i, line := range allLines {
		lineParts := strings.Split(line, " ")
		num, err := checkIfStringIsNumeric(lineParts[0])
		if err != nil && i > 0 {
			previousLineWords := strings.Split(allLines[i-1], " ")

			_, err := checkIfStringIsNumeric(previousLineWords[0])
			if err == nil {
				directorySeq++
			}
		} else {
			directoriesMap[directorySeq] += num
		}
	}

	fmt.Println(directoriesMap)

	allSmallDirectoriesSum := calculateSmallDirectoriesSum(directoriesMap)
	fmt.Println(allSmallDirectoriesSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkIfStringIsNumeric(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func calculateSmallDirectoriesSum(directoriesMap map[int]int) int {
	allSmallDirectoriesSum := 0

	for _, sumValue := range directoriesMap {
		if sumValue < maxSize {
			allSmallDirectoriesSum += sumValue
		}
	}

	return allSmallDirectoriesSum
}
