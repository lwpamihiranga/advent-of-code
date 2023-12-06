package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input_file, _ := os.Open("input")

	file_scanner := bufio.NewScanner(input_file)
	file_scanner.Split(bufio.ScanLines)

	var file_lines []string

	for file_scanner.Scan() {
		file_lines = append(file_lines, file_scanner.Text())
	}

	input_file.Close()

	var calibration_values []int

	for _, line := range file_lines {
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
