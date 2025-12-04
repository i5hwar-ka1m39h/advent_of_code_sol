package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	Part_two()
}

func Part_one() {
	data, err := os.ReadFile("./001part1.txt")
	if err != nil {
		log.Println("error occured:", err)
	}

	complete_string := string(data)
	trimed_string := strings.TrimSpace(complete_string)
	positions := strings.Split(trimed_string, "\n")

	fmt.Println(len(positions))

	current_pos := 50
	zero_reach_count := 0
	for _, val := range positions {
		dir := val[:1]
		moves, err := strconv.Atoi(val[1:])
		if err != nil {
			log.Println("error occured while converting to number", err)
		}
		if dir == "L" {
			current_pos = current_pos - moves
		} else {
			current_pos = current_pos + moves
		}

		current_pos = current_pos % 100

		if current_pos == 0 {
			zero_reach_count++
		}

	}

	fmt.Println("zeroth position reached for ", zero_reach_count)
}

func Part_two() {
	data, err := os.ReadFile("./001part1.txt")
	if err != nil {
		log.Println("error occured:", err)
		return
	}

	complete := strings.TrimSpace(string(data))
	lines := strings.Split(complete, "\n")

	currentPos := 50
	totalZeroEvents := 0
	dialSize := 100

	for _, line := range lines {
		dir := line[:1]
		moves, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Println("error reading movement:", err)
			continue
		}

		if dir == "R" {
			currentPos = currentPos + moves
			totalZeroEvents += currentPos / dialSize
			currentPos = currentPos % dialSize

		} else {
			old_pos := currentPos
			currentPos = currentPos - moves

			if currentPos == 0 {
				totalZeroEvents++
			}

			if currentPos < 0 {
				if old_pos != 0 {
					totalZeroEvents++
				}

				extra_cycle := (moves - old_pos) / dialSize
				if extra_cycle > 0 {
					totalZeroEvents += extra_cycle
				}

				currentPos = ((currentPos % dialSize) + dialSize) % dialSize
			}

		}

	}
	fmt.Println("Total zero events:", totalZeroEvents)
}
