package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	var card_map = map[int]int{}
	card_length := len(file_lines)

	for _, card_line := range file_lines {
		card_line_split := strings.Split(card_line, ":")

		card_number_string := card_line_split[0]
		card_number_string, _ = strings.CutPrefix(card_number_string, "Card")
		card_number_string = strings.TrimSpace(card_number_string)

		card_number, _ := strconv.Atoi(strings.TrimSpace(card_number_string))

		winning_number_string := strings.Split(card_line_split[1], "|")[0]
		winning_number_string = strings.TrimSpace(winning_number_string)

		available_number_string := strings.Split(card_line_split[1], "|")[1]
		available_number_string = strings.TrimSpace(available_number_string)

		winning_number_slice := strings.Fields(winning_number_string)
		available_number_slice := strings.Fields(available_number_string)

		number_match_count := 0
		for _, winning_number := range winning_number_slice {
			for _, available_number := range available_number_slice {
				if strings.TrimSpace(available_number) == strings.TrimSpace(winning_number) {
					number_match_count += 1
					break
				}
			}
		}

		card_map[card_number] = number_match_count
	}

	var card_map_copy = map[int]int{}

	for k := range card_map {
		card_map_copy[k] = 1
	}

	// for k, v := range card_map_copy {
	// 	for i := 0; i < v; i++ {
	// 		for j := 0; j < card_map[k]; j++ {
	// 			if k+j+1 > card_length {
	// 				break
	// 			}
	// 			card_map_copy[k+j+1] += 1
	// 		}
	// 	}
	// }

	// NOTE: If we try to loop thourgh card_map_copy map using range as above commented block
	// it will give various answers in each run. Hence chose to loop using a for loop
	for index := 1; index <= card_length; index++ {
		for i := 0; i < card_map_copy[index]; i++ {
			for j := 0; j < card_map[index]; j++ {
				if index+j+1 > card_length {
					break
				}
				card_map_copy[index+j+1] += 1
			}
		}
	}

	var answer int
	for _, v := range card_map_copy {
		answer += v
	}

	fmt.Printf("Answer for day4: %d\n", answer)
}
