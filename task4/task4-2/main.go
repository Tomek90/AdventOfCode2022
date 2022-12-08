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
	var fullCoverageCount int

	f, err := os.Open("elves_assignments.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		assignments := strings.Split(scanner.Text(), ",")
		elf1Assignment := assignments[0]
		elf2Assignment := assignments[1]

		elf1SectionRangeStart, elf1SectionRangeEnd, err := getSectionRangeNumbers(elf1Assignment)
		if err != nil {
			log.Fatal(err)
		}

		elf2SectionRangeStart, elf2SectionRangeEnd, err := getSectionRangeNumbers(elf2Assignment)
		if err != nil {
			log.Fatal(err)
		}

		elf1Sections := getAllElfSections(elf1SectionRangeStart, elf1SectionRangeEnd)
		elf2Sections := getAllElfSections(elf2SectionRangeStart, elf2SectionRangeEnd)

		if checkIfSectionContainedWithin(elf1Sections, elf2Sections) {
			fullCoverageCount++
		}
	}

	fmt.Println(fullCoverageCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getSectionRangeNumbers(assignment string) (int, int, error) {
	firstSectionStr := strings.Split(assignment, "-")[0]
	endSectionStr := strings.Split(assignment, "-")[1]

	firstSectionInt, err := strconv.Atoi(firstSectionStr)
	if err != nil {
		return 0, 0, err
	}

	endSectionInt, err := strconv.Atoi(endSectionStr)
	if err != nil {
		return 0, 0, err
	}

	return firstSectionInt, endSectionInt, nil
}

func getAllElfSections(startSection, endSection int) (allSectionList []int) {
	for i := startSection; i <= endSection; i++ {
		allSectionList = append(allSectionList, i)
	}

	return
}

func checkIfSectionContainedWithin(elf1Sections, elf2Sections []int) bool {
	section1ContainedIn2 := []bool{}

	for _, sectionNumberElf1 := range elf1Sections {
		for _, sectionNumberElf2 := range elf2Sections {
			if sectionNumberElf1 == sectionNumberElf2 {
				section1ContainedIn2 = append(section1ContainedIn2, true)

				break
			}
		}
	}

	section2ContainedIn1 := []bool{}

	for _, sectionNumberElf2 := range elf2Sections {
		for _, sectionNumberElf1 := range elf1Sections {
			if sectionNumberElf2 == sectionNumberElf1 {
				section2ContainedIn1 = append(section2ContainedIn1, true)

				break
			}
		}
	}

	return len(section1ContainedIn2) == len(elf1Sections) || len(section2ContainedIn1) == len(elf2Sections)
}
