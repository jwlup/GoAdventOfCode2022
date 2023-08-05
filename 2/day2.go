package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func score_round(round string) (int, error) {
	actions := strings.Split(round, " ")
	if len(actions) != 2 {
		return -1, errors.New("Not a valid round - expect exactly 2 actions")
	}
	shape_score := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	score, err := outcome_score(actions)
	if err != nil {
		return -1, err
	}
	shape_score_round := shape_score["A"]
	score += shape_score_round
	return score, nil
}

func outcome_score(actions []string) (int, error) {
	results := 0
	your_shape := actions[1]
	//outcome := [3]int{3, 6, 0}

	switch actions[0] {
	case "A":
		if your_shape == "Y" {
			results = 6
		} else if your_shape == "X" {
			results = 3
		}
	case "B":
		if your_shape == "Z" {
			results = 6
		} else if your_shape == "Y" {
			results = 3
		}
	case "C":
		if your_shape == "X" {
			results = 6
		} else if your_shape == "Z" {
			results = 3
		}
	default:
		return -1, errors.New("Invalid action by opponent")
	}
	return results, nil
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	per_round := strings.Split(string(content), "\n")

	total_score := 0
	for _, round := range per_round {
		score, err := score_round(round)
		if err != nil {
			log.Fatal(err)
		}
		total_score += score
	}
	fmt.Printf("Total Score (Part 1): %d", total_score)

}
