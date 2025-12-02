package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
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
