package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func get_elf_total(per_elf string) (int, error) {
	per_snack := strings.Split(per_elf, "\n")
	total := 0
	for _, snack := range per_snack {
		calories, err := strconv.Atoi(snack)
		if err != nil {
			return -1, err
		}
		total += calories
	}
	return total, nil
}

func third_calories_elf(top_three_calories [3]int) (int, int) {
	min_elf_position := 0
	min_calories := top_three_calories[0]
	for elf := 1; elf < 3; elf++ {
		if top_three_calories[elf] < min_calories {
			min_calories = top_three_calories[elf]
			min_elf_position = elf
		}
	}
	return min_calories, min_elf_position
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	per_elf_string := strings.Split(string(content), "\n\n")

	max_per_elf := 0
	max_elf := 0

	for elf_index, elf := range per_elf_string {
		per_elf, err := get_elf_total(elf)
		if err != nil {
			log.Fatal(err)
		}
		if per_elf > max_per_elf {
			max_elf = elf_index
			max_per_elf = per_elf
		}
	}
	fmt.Printf("Calories by most caloried elf: %d, Most caloried elf: %d\n", max_per_elf, max_elf+1)

	top_three_calories := [3]int{0, 0, 0}

	for _, elf := range per_elf_string {
		per_elf, err := get_elf_total(elf)
		if err != nil {
			log.Fatal(err)
		}
		third_calories, third_elf_position := third_calories_elf(top_three_calories)
		if per_elf > third_calories {
			top_three_calories[third_elf_position] = per_elf
		}
	}

	top_three_sum := 0

	for elf := 0; elf < 3; elf++ {
		top_three_sum += top_three_calories[elf]
	}
	fmt.Printf("Calories on top three elves: %d\n", top_three_sum)
}
