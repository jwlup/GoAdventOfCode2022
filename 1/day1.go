package main

import (
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

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	per_elf := strings.Split(string(content), "\n\n")

	max_per_elf := 0
	max_elf := 0

	for elf_index, elf := range per_elf {
		per_elf, err := get_elf_total(elf)
		if err != nil {
			log.Fatal(err)
		}
		if per_elf > max_per_elf {
			max_elf = elf_index
			max_per_elf = per_elf
		}
	}
	println("%v,%v", max_per_elf, max_elf+1)
}
