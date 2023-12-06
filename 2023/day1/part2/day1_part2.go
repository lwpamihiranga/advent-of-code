package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numbers_map = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	input_file, _ := os.Open("../input")

	file_scanner := bufio.NewScanner(input_file)
	file_scanner.Split(bufio.ScanLines)

	var file_lines []string

	for file_scanner.Scan() {
		file_lines = append(file_lines, file_scanner.Text())
	}

	input_file.Close()

	var calibration_values []int

	for _, line := range file_lines {
		for {
			var slice [][]int
			line, slice = replaceStrToNumber(line)

			if slice == nil {
				break
			}
		}

		var first_digit string
		var last_digit string

		for index, char := range line {
			char_str := string(char)
			_, err := strconv.Atoi(char_str)

			if err != nil {
				if index == len(line)-1 && last_digit == "" {
					last_digit = first_digit
				}

				continue
			}

			if first_digit == "" {
				first_digit = char_str
			} else {
				last_digit = char_str
			}

			if index == len(line)-1 && last_digit == "" {
				last_digit = first_digit
			}
		}

		number, _ := strconv.Atoi(first_digit + last_digit)

		calibration_values = append(calibration_values, number)
	}

	// fmt.Println(calibration_values)
	// fmt.Println(len(calibration_values))

	var answer int

	for _, value := range calibration_values {
		answer += value
	}

	fmt.Printf("Answer for day1: %d\n", answer)
}

func replaceStrToNumber(line string) (modified_line string, index_slices [][]int) {
	// var index_slices [][]int

	for key := range numbers_map {

		r, _ := regexp.Compile(key)

		result_slice := r.FindAllIndex([]byte(line), 100)

		if len(result_slice) != 0 {
			index_slices = append(index_slices, result_slice...)
		}
	}

	if len(index_slices) >= 1 {
		var min_size []int

		min_size = index_slices[0]
		for _, value := range index_slices {
			if min_size[0] > value[0] {
				min_size = value
			}
		}

		text := line[min_size[0]:min_size[1]]

		line = strings.Replace(line, text[:len(text)-1], numbers_map[text], 1)
	}

	index_slices = nil

	for key := range numbers_map {

		r, _ := regexp.Compile(key)

		result_slice := r.FindAllIndex([]byte(line), 100)

		if len(result_slice) != 0 {
			index_slices = append(index_slices, result_slice...)
		}
	}

	return line, index_slices
}
