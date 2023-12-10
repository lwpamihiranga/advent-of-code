package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	var valid_part_numbers []string

	for line_index, line_to_check := range file_lines {
		before_line := ""
		after_line := ""

		if line_index > 0 {
			before_line = file_lines[line_index-1]
		}

		if line_index < len(file_lines)-1 {
			after_line = file_lines[line_index+1]
		}

		for char_index, char := range line_to_check {
			character := string(char)
			if character == "." {
				continue
			}

			_, err := strconv.Atoi(character)

			if err == nil {
				continue
			}

			// fmt.Println("symbol", char_index, string(char))

			adjacent_char_index_slice := struct {
				before_index  int
				current_index int
				next_index    int
			}{
				char_index - 1,
				char_index,
				char_index + 1,
			}

			// fmt.Printf("%+v\n", adjacent_char_index_slice)

			if before_line != "" {
				result_left_before_line := ""
				result_middle_before_line := ""
				result_right_before_line := ""

				result_left_before_line = readToLeft(adjacent_char_index_slice.before_index, before_line)
				result_middle_before_line = readToLeft(adjacent_char_index_slice.current_index, before_line)
				result_right_before_line = readToRight(adjacent_char_index_slice.next_index, before_line)

				if result_middle_before_line != "" {
					if result_right_before_line != "" {
						valid_part_numbers = append(valid_part_numbers, result_middle_before_line+result_right_before_line)
					} else {
						valid_part_numbers = append(valid_part_numbers, result_middle_before_line)
					}
				} else {
					if result_left_before_line != "" {
						valid_part_numbers = append(valid_part_numbers, result_left_before_line)
					}

					if result_right_before_line != "" {
						valid_part_numbers = append(valid_part_numbers, result_right_before_line)
					}
				}
			}

			if after_line != "" {
				result_left_after_line := ""
				result_middle_after_line := ""
				result_right_after_line := ""

				result_left_after_line = readToLeft(adjacent_char_index_slice.before_index, after_line)
				result_middle_after_line = readToLeft(adjacent_char_index_slice.current_index, after_line)
				result_right_after_line = readToRight(adjacent_char_index_slice.next_index, after_line)

				if result_middle_after_line != "" {
					if result_right_after_line != "" {
						valid_part_numbers = append(valid_part_numbers, result_middle_after_line+result_right_after_line)
					} else {
						valid_part_numbers = append(valid_part_numbers, result_middle_after_line)
					}
				} else {
					if result_left_after_line != "" {
						valid_part_numbers = append(valid_part_numbers, result_left_after_line)
					}

					if result_right_after_line != "" {
						valid_part_numbers = append(valid_part_numbers, result_right_after_line)
					}
				}
			}

			result_left := ""
			result_middle := ""
			result_right := ""

			result_left = readToLeft(adjacent_char_index_slice.before_index, line_to_check)
			result_middle = readToLeft(adjacent_char_index_slice.current_index, line_to_check)
			result_right = readToRight(adjacent_char_index_slice.next_index, line_to_check)

			if result_middle != "" {
				if result_right != "" {
					valid_part_numbers = append(valid_part_numbers, result_middle+result_right)
				} else {
					valid_part_numbers = append(valid_part_numbers, result_middle)
				}
			} else {
				if result_left != "" {
					valid_part_numbers = append(valid_part_numbers, result_left)
				}

				if result_right != "" {
					valid_part_numbers = append(valid_part_numbers, result_right)
				}
			}

			// fmt.Println("result left", result_left)
			// fmt.Println("result middle", result_middle)
			// fmt.Println("result right", result_right)

		}

		// fmt.Println(before_line)
		// fmt.Println(line_index, line_to_check)
		// fmt.Println(after_line)

		// if line_index == 10 {
		// 	break
		// }
	}

	// fmt.Printf("%+v\n", valid_part_numbers)

	var answer int

	for _, part_number_string := range valid_part_numbers {
		part_number, err := strconv.Atoi(part_number_string)

		if err != nil {
			continue
		}

		answer += part_number
	}

	fmt.Printf("Answer for day3: %d\n", answer)
}

func checkIfNumber(index int, line string) (is_number bool) {
	is_number = true

	_, err := strconv.Atoi(string(line[index]))

	if err != nil {
		is_number = false
	}

	return is_number
}

func readToLeft(index int, line string) string {
	result := ""
	start_index := index

	for {
		if index < 0 {
			break
		}

		is_number := checkIfNumber(index, line)

		if !is_number {
			break
		}

		result = line[index : start_index+1]

		index -= 1
	}

	return result
}

func readToRight(index int, line string) string {
	result := ""
	start_index := index

	for {
		if index >= len(line) {
			break
		}

		is_number := checkIfNumber(index, line)

		if !is_number {
			break
		}

		index += 1
	}

	if index > len(line) {
		index -= 1
	}

	result = line[start_index:index]

	return result
}
