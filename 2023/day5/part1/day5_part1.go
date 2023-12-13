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

	var seeds_map = map[int]int{}
	var seed_to_soil_map = map[int]int{}
	var soil_to_fertilizer_map = map[int]int{}
	var fertilizer_to_water_map = map[int]int{}
	var water_to_light_map = map[int]int{}
	var light_to_temperature_map = map[int]int{}
	var temperature_to_humidity_map = map[int]int{}
	var humidity_to_location_map = map[int]int{}

	for index, line := range file_lines {
		if index == 0 {
			seeds_string := strings.Split(line, "seeds:")[1]
			seeds_string = strings.TrimSpace(seeds_string)
			for _, seed := range strings.Fields(seeds_string) {
				seed_val, _ := strconv.Atoi(seed)
				seeds_map[seed_val] = seed_val
			}
		}

		if line == "seed-to-soil map:" {
			for i := 1; ; i++ {
				next_line := file_lines[index+i]

				numbers_slice := strings.Fields(next_line)

				if len(numbers_slice) == 0 {
					break
				}

				seed_to_soil_map = generateMap(seed_to_soil_map, numbers_slice)
			}
		}

		if line == "soil-to-fertilizer map:" {
			for i := 1; ; i++ {
				next_line := file_lines[index+i]

				numbers_slice := strings.Fields(next_line)

				if len(numbers_slice) == 0 {
					break
				}

				soil_to_fertilizer_map = generateMap(soil_to_fertilizer_map, numbers_slice)

			}
		}

		if line == "fertilizer-to-water map:" {
			for i := 1; ; i++ {
				next_line := file_lines[index+i]

				numbers_slice := strings.Fields(next_line)

				if len(numbers_slice) == 0 {
					break
				}

				fertilizer_to_water_map = generateMap(fertilizer_to_water_map, numbers_slice)

			}
		}

		if line == "water-to-light map:" {
			for i := 1; ; i++ {
				next_line := file_lines[index+i]

				numbers_slice := strings.Fields(next_line)

				if len(numbers_slice) == 0 {
					break
				}

				water_to_light_map = generateMap(water_to_light_map, numbers_slice)
			}
		}

		if line == "light-to-temperature map:" {
			for i := 1; ; i++ {
				next_line := file_lines[index+i]

				numbers_slice := strings.Fields(next_line)

				if len(numbers_slice) == 0 {
					break
				}

				light_to_temperature_map = generateMap(light_to_temperature_map, numbers_slice)
			}
		}

		if line == "temperature-to-humidity map:" {
			for i := 1; ; i++ {
				next_line := file_lines[index+i]

				numbers_slice := strings.Fields(next_line)

				if len(numbers_slice) == 0 {
					break
				}

				temperature_to_humidity_map = generateMap(temperature_to_humidity_map, numbers_slice)
			}
		}

		if line == "humidity-to-location map:" {
			for i := 1; ; i++ {
				if index+i >= len(file_lines) {
					break
				}

				next_line := file_lines[index+i]

				numbers_slice := strings.Fields(next_line)

				if len(numbers_slice) == 0 {
					break
				}

				humidity_to_location_map = generateMap(humidity_to_location_map, numbers_slice)
			}
		}
	}

	// fmt.Printf("seeds_map: %+v\n", seeds_map)
	// fmt.Printf("seed_to_soil_map: %+v\n", seed_to_soil_map)
	// fmt.Printf("soil_to_fertilizer_map: %+v\n", soil_to_fertilizer_map)
	// fmt.Printf("fertilizer_to_water_map: %+v\n", fertilizer_to_water_map)
	// fmt.Printf("water_to_light_map: %+v\n", water_to_light_map)
	// fmt.Printf("light_to_temperature_map: %+v\n", light_to_temperature_map)
	// fmt.Printf("temperature_to_humidity_map: %+v\n", temperature_to_humidity_map)
	// fmt.Printf("humidity_to_location_map: %+v\n", humidity_to_location_map)

	lowest_location := 0

	for seed_val := range seeds_map {
		index := 0

		soil_val, ok := seed_to_soil_map[seed_val]

		if !ok {
			soil_val = seed_val
		}

		fertilizer_val, ok := soil_to_fertilizer_map[soil_val]

		if !ok {
			fertilizer_val = soil_val
		}

		water_val, ok := fertilizer_to_water_map[fertilizer_val]

		if !ok {
			water_val = fertilizer_val
		}

		light_val, ok := water_to_light_map[water_val]

		if !ok {
			light_val = water_val
		}

		temp_val, ok := light_to_temperature_map[light_val]

		if !ok {
			temp_val = light_val
		}

		humidity_val, ok := temperature_to_humidity_map[temp_val]

		if !ok {
			humidity_val = temp_val
		}

		location_val, ok := humidity_to_location_map[humidity_val]

		if !ok {
			location_val = humidity_val
		}

		fmt.Println(seed_val, soil_val, fertilizer_val, water_val, light_val, temp_val, humidity_val, location_val)

		if index == 0 {
			lowest_location = location_val
		} else if lowest_location > location_val {
			lowest_location = location_val
		}

		index += 1
	}

	fmt.Println("Lowerst Location: ", lowest_location)
}

// func generateMap(map_to_update *map[int]int, numbers_slice []string) (generated_map map[int]int) {
// 	generated_map = map[int]int{}

// 	dest_range_start, _ := strconv.Atoi(numbers_slice[0])
// 	source_range_start, _ := strconv.Atoi(numbers_slice[1])
// 	range_value, _ := strconv.Atoi(numbers_slice[2])

// 	for i := 0; i < range_value; i++ {
// 		generated_map[source_range_start+i] = dest_range_start + i
// 	}

// 	return generated_map
// }

func generateMap(map_to_update map[int]int, numbers_slice []string) map[int]int {
	dest_range_start, _ := strconv.Atoi(numbers_slice[0])
	source_range_start, _ := strconv.Atoi(numbers_slice[1])
	range_value, _ := strconv.Atoi(numbers_slice[2])

	for i := 0; i < range_value; i++ {
		map_to_update[source_range_start+i] = dest_range_start + i
	}

	return map_to_update
}
