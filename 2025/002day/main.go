package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// string_holder, _ := Give_out_array()
	//
	// fmt.Println("size of the origial set", len(string_holder))
	// var valid_val_arr []Range
	// for _, val := range string_holder {
	// 	range_val, _ := Check_range_odd(val)
	//
	// 	if range_val != nil {
	// 		valid_val_arr = append(valid_val_arr, *range_val)
	// 	}
	// }
	//
	// fmt.Println("here is valid array", valid_val_arr)
	// fmt.Println("size of the cleared one", len(valid_val_arr))

	Is_valid_id("1188511885")
}

func Give_out_array() ([]string, error) {
	data, err := os.ReadFile("./data.txt")
	if err != nil {
		log.Println("error occured while reading the file")
		return nil, err
	}

	var full_string string = string(data)
	num_store := strings.Split(full_string, ",")
	return num_store, nil
}

type Range struct {
	start string
	end   string
}

// this should check for the validity of the string
// if not valid then remove them from the input

func Check_range_odd(number_string string) (*Range, error) {
	store := strings.Split(number_string, "-")
	start := store[0]
	end := store[1]

	if len(start)%2 == 0 || len(end)%2 == 0 {
		valid := &Range{
			start: start,
			end:   end,
		}

		return valid, nil
	}

	return nil, errors.New("not a valid range")
}

func Is_valid_id(num string) bool {
	if len(num)%2 == 0 {
		mid := len(num) / 2
		first_half := num[:mid]
		second_half := num[mid:]

		first_half_num, err := strconv.Atoi(first_half)
		if err != nil {
			fmt.Println("error occured while checking validation", err)
			os.Exit(1)
		}

		second_half_num, err := strconv.Atoi(second_half)
		if err != nil {
			fmt.Println("error occured while checking validation", err)
			os.Exit(1)
		}

		if first_half_num == second_half_num {
			return false
		}

		return true

	}

	return true
}
