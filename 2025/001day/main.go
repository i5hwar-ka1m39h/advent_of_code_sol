package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("./001part1.txt")
	if err != nil {
		log.Println("error occured:", err)
	}

	fmt.Println(string(data))
}
