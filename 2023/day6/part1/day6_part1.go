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

	var races_slice []struct {
		time     int
		distance int
	}

	time_string := strings.Split(file_lines[0], "Time:")[1]
	time_slice := strings.Fields(time_string)

	distance_string := strings.Split(file_lines[1], "Distance:")[1]
	distance_slice := strings.Fields(distance_string)

	for index, value := range time_slice {
		t, _ := strconv.Atoi(value)
		d, _ := strconv.Atoi(distance_slice[index])

		race := struct {
			time     int
			distance int
		}{
			time:     t,
			distance: d,
		}

		races_slice = append(races_slice, race)
	}

	var win_chances_slice []int

	for _, race := range races_slice {
		win_chances := 0

		for i := 1; i < race.time; i++ {
			possible_distance := i * (race.time - i)

			if possible_distance > race.distance {
				win_chances += 1
			}
		}

		win_chances_slice = append(win_chances_slice, win_chances)
	}

	var answer int

	for index, value := range win_chances_slice {
		if index == 0 {
			answer = value
			continue
		}

		answer = answer * value
	}

	fmt.Printf("Answer for day6: %d\n", answer)
}
