package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const fourteenLast = 14

func main() {
	f, err := os.Open("signal.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	signal, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	signalText := string(signal)

	for i := 0; i < len(signalText); i++ {
		var lastChars []string
		if i > 13 {
			for j := i; j > i-14; j-- {
				lastChars = append(lastChars, string(signalText[j]))
			}

			if checkUniquenessOfLastChars(lastChars) {
				fmt.Println(i + 1)

				break
			}
		}
	}
}

func checkUniquenessOfLastChars(lastChars []string) bool {
	unique := true
	charMap := map[string]bool{}

	for _, char := range lastChars {
		if _, ok := charMap[char]; !ok {
			charMap[char] = true
		} else {
			return false
		}
	}

	return unique
}
