package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	input_file, _ := os.Open("../input")

	file_scanner := bufio.NewScanner(input_file)
	file_scanner.Split(bufio.ScanLines)

	var file_lines []string

	for file_scanner.Scan() {
		file_lines = append(file_lines, file_scanner.Text())
	}

	input_file.Close()

	var number_match_count_splice []int

	for _, card_line := range file_lines {
		card_line_split := strings.Split(card_line, ":")

		winning_number_string := strings.Split(card_line_split[1], "|")[0]
		winning_number_string = strings.TrimSpace(winning_number_string)

		available_number_string := strings.Split(card_line_split[1], "|")[1]
		available_number_string = strings.TrimSpace(available_number_string)

		winning_number_slice := strings.Fields(winning_number_string)
		available_number_slice := strings.Fields(available_number_string)

		number_match_count := -1
		for _, winning_number := range winning_number_slice {
			for _, available_number := range available_number_slice {
				if strings.TrimSpace(available_number) == strings.TrimSpace(winning_number) {
					number_match_count += 1
					break
				}
			}
		}

		if number_match_count != -1 {
			number_match_count_splice = append(number_match_count_splice, number_match_count)
		}
	}

	var answer float64

	for _, number := range number_match_count_splice {
		answer += math.Pow(2, float64(number))
	}

	fmt.Printf("Answer for day4: %.0f\n", answer)
}
