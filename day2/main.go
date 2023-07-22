package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Rock     = "A"
	Paper    = "B"
	Scissors = "C"
)

const (
	Lose = "X"
	Draw = "Y"
	Win  = "Z"
)

const (
	ScoreRock     = 1
	ScorePaper    = 2
	ScoreScissors = 3

	ScoreLose = 0
	ScoreDraw = 3
	ScoreWin  = 6
)

func main() {
	f, _ := os.OpenFile("./input.txt", os.O_RDONLY, 0644)
	defer f.Close()

	s := bufio.NewScanner(f)

	total := 0

	for s.Scan() {
		line := s.Text()
		sl := strings.Split(line, " ")
		p1 := sl[0]
		outcome := sl[1]

		if outcome == Win {
			total += ScoreWin

			if p1 == Rock {
				total += ScorePaper
			} else if p1 == Paper {
				total += ScoreScissors
			} else if p1 == Scissors {
				total += ScoreRock
			}
		}

		if outcome == Draw {
			total += ScoreDraw

			if p1 == Rock {
				total += ScoreRock
			} else if p1 == Paper {
				total += ScorePaper
			} else if p1 == Scissors {
				total += ScoreScissors
			}
		}

		if outcome == Lose {
			total += ScoreLose

			if p1 == Rock {
				total += ScoreScissors
			} else if p1 == Paper {
				total += ScoreRock
			} else if p1 == Scissors {
				total += ScorePaper
			}
		}
	}

	fmt.Printf("total: %v\n", total)
}
