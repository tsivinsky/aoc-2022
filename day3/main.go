package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

func main() {
	f, _ := os.OpenFile("./input.txt", os.O_RDONLY, 0644)
	defer f.Close()

	s := bufio.NewScanner(f)

	rucksacks := make(map[int][]string)

	rucksackIndex := 0
	row := 1
	for s.Scan() {
		line := s.Text()

		rucksacks[rucksackIndex] = append(rucksacks[rucksackIndex], line)

		if row%3 == 0 {
			rucksackIndex++
		}

		row++
	}

	total := 0

	for _, rucksack := range rucksacks {
		var badge string

		for _, letter := range rucksack[0] {
			for _, letter2 := range rucksack[1] {
				for _, letter3 := range rucksack[2] {
					if letter == letter2 && letter == letter3 {
						badge = string(letter)
					}
				}
			}
		}

		badgePriority := 0
		for i, letter := range letters {
			if badge == letter {
				badgePriority = i + 1
			}
		}

		total += badgePriority
	}

	fmt.Printf("total: %v\n", total)
}

func isLowerCase(str string) bool {
	return str == strings.ToLower(str)
}
