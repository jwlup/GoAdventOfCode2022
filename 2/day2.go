package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func circular_array(array [3]int, index int) int {
	return (3 + index) % 3
}

func score_round(round string) (int, error) {
	actions := strings.Split(round, " ")
	if len(actions) != 2 {
		return -1, errors.New("Not a valid round - expect exactly 2 actions")
	}
	encode_your_shape := map[string]int{
		"X": 0,
		"Y": 1,
		"Z": 2,
	}
	encode_opp_shape := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
	}
	action_codes := [2]int{encode_opp_shape[actions[0]], encode_your_shape[actions[1]]}
	score, err := outcome_score_p1(action_codes)
	if err != nil {
		return -1, err
	}

	return score, nil
}

func outcome_score_p1(actions [2]int) (int, error) {
	your_action := actions[1]
	opp_action := actions[0]
	if your_action < 0 || your_action > 2 || opp_action < 0 || opp_action > 2 {
		return -1, errors.New("Action Code out of Range")
	}
	outcome := [3]int{3, 6, 0}
	outcome_index := circular_array(outcome, your_action-opp_action)
	result := outcome[outcome_index]
	return result + your_action + 1, nil
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
