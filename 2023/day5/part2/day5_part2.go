package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type range_desc struct {
	s_start int
	s_end   int
	d_start int
	d_end   int
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

	var seeds_slice []int
	var seed_to_soil_slice []range_desc
	var soil_to_fertilizer_slice []range_desc
	var fertilizer_to_water_slice []range_desc
	var water_to_light_slice []range_desc
	var light_to_temperature_slice []range_desc
	var temperature_to_humidity_slice []range_desc
	var humidity_to_location_slice []range_desc

	for index, line := range file_lines {
		if index == 0 {
			seeds_string := strings.Split(line, "seeds:")[1]
			seeds_string = strings.TrimSpace(seeds_string)
			for _, seed := range strings.Fields(seeds_string) {
				seed_val, _ := strconv.Atoi(seed)
				// seeds_map[seed_val] = seed_val
				seeds_slice = append(seeds_slice, seed_val)
			}
		}

		if line == "seed-to-soil map:" {
			for i := 1; ; i++ {
				next_line := file_lines[index+i]

				numbers_slice := strings.Fields(next_line)

				if len(numbers_slice) == 0 {
					break
				}

				dest_range_start, _ := strconv.Atoi(numbers_slice[0])
				source_range_start, _ := strconv.Atoi(numbers_slice[1])
				range_value, _ := strconv.Atoi(numbers_slice[2])

				value := range_desc{
					s_start: source_range_start,
					s_end:   source_range_start + range_value - 1,
					d_start: dest_range_start,
					d_end:   dest_range_start + range_value - 1,
				}

				seed_to_soil_slice = append(seed_to_soil_slice, value)
			}
		}

		if line == "soil-to-fertilizer map:" {
			for i := 1; ; i++ {
				next_line := file_lines[index+i]

				numbers_slice := strings.Fields(next_line)

				if len(numbers_slice) == 0 {
					break
				}

				dest_range_start, _ := strconv.Atoi(numbers_slice[0])
				source_range_start, _ := strconv.Atoi(numbers_slice[1])
				range_value, _ := strconv.Atoi(numbers_slice[2])

				value := range_desc{
					s_start: source_range_start,
					s_end:   source_range_start + range_value - 1,
					d_start: dest_range_start,
					d_end:   dest_range_start + range_value - 1,
				}

				soil_to_fertilizer_slice = append(soil_to_fertilizer_slice, value)
			}
		}

		if line == "fertilizer-to-water map:" {
			for i := 1; ; i++ {
				next_line := file_lines[index+i]

				numbers_slice := strings.Fields(next_line)

				if len(numbers_slice) == 0 {
					break
				}

				dest_range_start, _ := strconv.Atoi(numbers_slice[0])
				source_range_start, _ := strconv.Atoi(numbers_slice[1])
				range_value, _ := strconv.Atoi(numbers_slice[2])

				value := range_desc{
					s_start: source_range_start,
					s_end:   source_range_start + range_value - 1,
					d_start: dest_range_start,
					d_end:   dest_range_start + range_value - 1,
				}

				fertilizer_to_water_slice = append(fertilizer_to_water_slice, value)
			}
		}

		if line == "water-to-light map:" {
			for i := 1; ; i++ {
				next_line := file_lines[index+i]

				numbers_slice := strings.Fields(next_line)

				if len(numbers_slice) == 0 {
					break
				}

				dest_range_start, _ := strconv.Atoi(numbers_slice[0])
				source_range_start, _ := strconv.Atoi(numbers_slice[1])
				range_value, _ := strconv.Atoi(numbers_slice[2])

				value := range_desc{
					s_start: source_range_start,
					s_end:   source_range_start + range_value - 1,
					d_start: dest_range_start,
					d_end:   dest_range_start + range_value - 1,
				}

				water_to_light_slice = append(water_to_light_slice, value)
			}
		}

		if line == "light-to-temperature map:" {
			for i := 1; ; i++ {
				next_line := file_lines[index+i]

				numbers_slice := strings.Fields(next_line)

				if len(numbers_slice) == 0 {
					break
				}

				dest_range_start, _ := strconv.Atoi(numbers_slice[0])
				source_range_start, _ := strconv.Atoi(numbers_slice[1])
				range_value, _ := strconv.Atoi(numbers_slice[2])

				value := range_desc{
					s_start: source_range_start,
					s_end:   source_range_start + range_value - 1,
					d_start: dest_range_start,
					d_end:   dest_range_start + range_value - 1,
				}

				light_to_temperature_slice = append(light_to_temperature_slice, value)
			}
		}

		if line == "temperature-to-humidity map:" {
			for i := 1; ; i++ {
				next_line := file_lines[index+i]

				numbers_slice := strings.Fields(next_line)

				if len(numbers_slice) == 0 {
					break
				}

				dest_range_start, _ := strconv.Atoi(numbers_slice[0])
				source_range_start, _ := strconv.Atoi(numbers_slice[1])
				range_value, _ := strconv.Atoi(numbers_slice[2])

				value := range_desc{
					s_start: source_range_start,
					s_end:   source_range_start + range_value - 1,
					d_start: dest_range_start,
					d_end:   dest_range_start + range_value - 1,
				}

				temperature_to_humidity_slice = append(temperature_to_humidity_slice, value)
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

				dest_range_start, _ := strconv.Atoi(numbers_slice[0])
				source_range_start, _ := strconv.Atoi(numbers_slice[1])
				range_value, _ := strconv.Atoi(numbers_slice[2])

				value := range_desc{
					s_start: source_range_start,
					s_end:   source_range_start + range_value - 1,
					d_start: dest_range_start,
					d_end:   dest_range_start + range_value - 1,
				}

				humidity_to_location_slice = append(humidity_to_location_slice, value)
			}
		}
	}

	var seeds_range_slice []struct {
		start int
		end   int
	}
	is_skip := false

	for index, val := range seeds_slice {
		if is_skip {
			is_skip = false
			continue
		}

		value := struct {
			start int
			end   int
		}{
			start: val,
			end:   val + seeds_slice[index+1] - 1,
		}

		seeds_range_slice = append(seeds_range_slice, value)

		is_skip = true
	}

	lowest_location := -1

	for index, item := range seeds_range_slice {
		for i := item.start; i <= item.end; i++ {
			soil_value := findDestinationValue(i, seed_to_soil_slice)
			fertilizer_val := findDestinationValue(soil_value, soil_to_fertilizer_slice)
			water_val := findDestinationValue(fertilizer_val, fertilizer_to_water_slice)
			light_val := findDestinationValue(water_val, water_to_light_slice)
			temp_val := findDestinationValue(light_val, light_to_temperature_slice)
			humidity_val := findDestinationValue(temp_val, temperature_to_humidity_slice)
			location_val := findDestinationValue(humidity_val, humidity_to_location_slice)

			// fmt.Println(seed_val, soil_value, fertilizer_val, water_val, light_val, temp_val, humidity_val, location_val)

			if index == 0 && lowest_location == -1 {
				lowest_location = location_val
			} else if lowest_location > location_val {
				lowest_location = location_val
			}
		}
	}

	// NOTE: Part2 answer took about 5 mins to run
	fmt.Printf("Answer for day5: %d\n", lowest_location)
}

func findDestinationValue(source_value int, range_desc_slice []range_desc) (dest_value int) {
	dest_value = source_value
	for _, item := range range_desc_slice {
		if source_value >= item.s_start && source_value <= item.s_end {
			dest_value = item.d_start + (source_value - item.s_start)
			break
		}
	}

	return dest_value
}
