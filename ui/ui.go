package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/xie3245/merge_intervals_go/types"
)

const end_word string = "done"
const (
	undefined uint8 = iota
	had_lower
	had_upper
)

var state uint8 = undefined

func GetIntervals(data_ch chan types.Interval) {
	fmt.Println("Please enter integers in range 0 ~ 63. Finish by inputting: done")
	fmt.Println("lower bound: ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		if is_end(input) {
			break;
		}

		if worked, num := get_number_if_in_range(input); worked {
			interpret_to_interval(num, data_ch) 
		}

		print_next_indication()		
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	close(data_ch)
}

func print_next_indication() {
	switch state {
	case had_lower: {
		fmt.Println("upper bound:")
	}
	case had_upper: {
		fmt.Println("lower bound:")
	}
	}
}

func PrintResult(ch chan types.Interval) {
	var result strings.Builder
	for intv := range ch {
		result.WriteString(types.ToString(intv))
	}
	fmt.Println("merged: ", result.String())
}

var lower uint8
var upper uint8

func interpret_to_interval(num uint8, data_ch chan types.Interval) {
	switch state {
	case undefined: {
		lower = num
		state = had_lower
	}
	case had_lower: {
		upper = num
		state = had_upper
		intv := types.Interval{lower, upper}
		data_ch <- intv
	}
	case had_upper: {
		lower = num
		state = had_lower
	}
	}
}

func is_end(word string) bool {
	return word == end_word
}

func get_number_if_in_range(word string) (bool, uint8) {
	const base int = 10
	const typesize int = 8

	if num, err := strconv.ParseUint(word, base, typesize); err == nil {
		if num < 64 {
			return true, uint8(num)
		} else {
			fmt.Println("number out of range")
		}
	} else {
		fmt.Println(err)
	}
	return false, 0
}