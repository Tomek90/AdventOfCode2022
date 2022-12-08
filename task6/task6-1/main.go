package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

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
		if i > 2 {
			lastChars := []string{
				string(signalText[i]),
				string(signalText[i-1]),
				string(signalText[i-2]),
				string(signalText[i-3]),
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
