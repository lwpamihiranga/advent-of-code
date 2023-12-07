package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game_hand struct {
	red   int
	green int
	blue  int
}

const game_regex = "^Game.+:"

func main() {
	input_file, _ := os.Open("../input")

	file_scanner := bufio.NewScanner(input_file)
	file_scanner.Split(bufio.ScanLines)

	var file_lines []string

	for file_scanner.Scan() {
		file_lines = append(file_lines, file_scanner.Text())
	}

	input_file.Close()

	var min_hands []game_hand

	for _, line := range file_lines {
		// fmt.Println(line)

		r, _ := regexp.Compile(game_regex)

		game := r.FindString(line)

		var game_play_slice []game_hand

		game_rounds_string := strings.Split(line, game)[1]
		game_rounds_string = strings.TrimSpace(game_rounds_string)

		// fmt.Println(game_rounds_string)

		game_rounds := strings.Split(game_rounds_string, ";")

		for _, game_round := range game_rounds {
			game_round = strings.TrimSpace(game_round)

			// fmt.Println(game_round)

			game_play := game_hand{
				red:   0,
				green: 0,
				blue:  0,
			}

			// fmt.Printf("%+v\n", game_play)

			for _, value := range strings.Split(game_round, ",") {
				value = strings.TrimSpace(value)

				// fmt.Println(value)

				if strings.Contains(value, "red") {
					count, _ := strconv.Atoi(strings.TrimSpace(strings.Split(value, "red")[0]))

					game_play.red = count
				} else if strings.Contains(value, "green") {
					count, _ := strconv.Atoi(strings.TrimSpace(strings.Split(value, "green")[0]))

					game_play.green = count
				} else if strings.Contains(value, "blue") {
					count, _ := strconv.Atoi(strings.TrimSpace(strings.Split(value, "blue")[0]))

					game_play.blue = count
				}

			}

			game_play_slice = append(game_play_slice, game_play)
		}

		min_hand := findMinReqHand(game_play_slice)

		min_hands = append(min_hands, min_hand)
	}

	var answer int

	for _, value := range min_hands {
		power := value.red * value.green * value.blue
		answer += power
	}

	fmt.Printf("Answer for day2: %d\n", answer)
}

func findMinReqHand(hands []game_hand) (min_hand game_hand) {
	min_hand.red = 0
	min_hand.green = 0
	min_hand.blue = 0

	max_red := 0
	max_green := 0
	max_blue := 0

	for _, hand := range hands {
		if hand.red > max_red {
			max_red = hand.red
		}

		if hand.green > max_green {
			max_green = hand.green
		}

		if hand.blue > max_blue {
			max_blue = hand.blue
		}
	}

	min_hand.red = max_red
	min_hand.green = max_green
	min_hand.blue = max_blue

	return min_hand
}
