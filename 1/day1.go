package main

import (
	"log"
	"os"
	"strings"
)

func get_elf_total(per_elf string) int {
	per_snack := strings.Split(per_elf, "\n")
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	per_elf := strings.Split(string(content), "\n\n")

}
