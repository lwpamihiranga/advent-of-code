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

	total_cubes := game_hand{
		red:   12,
		green: 13,
		blue:  14,
	}

	var valid_game_ids []int

	for _, line := range file_lines {
		// fmt.Println(line)

		r, _ := regexp.Compile(game_regex)

		game := r.FindString(line)

		// fmt.Println(game)

		game_split := strings.Split(game, "Game")[1]
		game_split = strings.TrimSpace(game_split)
		game_split = strings.Split(game_split, ":")[0]

		// fmt.Println(game_split)
		// fmt.Printf("%T\n", game_split)

		game_id, _ := strconv.Atoi(game_split)
		is_valid_game := true

		// fmt.Println(game_id)
		// fmt.Printf("%T\n", game_id)

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

			// fmt.Printf("%+v\n", game_play)

			is_valid_round := checkValidGamePlay(total_cubes, game_play)

			// fmt.Println(is_valid_round)

			if !is_valid_round {
				is_valid_game = false
				break
			}
		}

		if is_valid_game {
			valid_game_ids = append(valid_game_ids, game_id)
		}
	}

	// fmt.Println(valid_game_ids)

	var answer int

	for _, value := range valid_game_ids {
		answer += value
	}

	fmt.Printf("Answer for day2: %d\n", answer)
}

func checkValidGamePlay(total game_hand, round game_hand) (is_valid_game bool) {
	is_valid_game = true

	if total.red < round.red {
		is_valid_game = false
	}

	if total.green < round.green {
		is_valid_game = false
	}

	if total.blue < round.blue {
		is_valid_game = false
	}

	return is_valid_game
}
