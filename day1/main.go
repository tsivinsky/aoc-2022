package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.OpenFile("./input.txt", os.O_RDONLY, 0644)
	defer f.Close()

	s := bufio.NewScanner(f)

	elves := make(map[int]int)

	elfIndex := 0
	for s.Scan() {
		line := s.Text()
		if line == "" {
			elfIndex++
			continue
		}

		caloriesCount, _ := strconv.Atoi(line)
		elves[elfIndex] += caloriesCount
	}

	max := 0
	for _, elf := range elves {
		if max <= elf {
			max = elf
		}
	}

	fmt.Printf("max: %v\n", max) // part 1 answer

	var top3 []int
	for i := 0; i < 3; i++ {
		max := 0
		id := 0

		for i, elf := range elves {
			if max <= elf {
				max = elf
				id = i
			}
		}

		delete(elves, id)

		top3 = append(top3, max)
	}

	top3Total := 0
	for _, max := range top3 {
		top3Total += max
	}

	fmt.Printf("top3Total: %v\n", top3Total) // part 2 answer
}
