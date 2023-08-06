package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func circular_array(array [3]int, index int) int {
	return array[(3+index)%3]
}

func score_round(round string) (int, int, error) {
	actions := strings.Split(round, " ")
	if len(actions) != 2 {
		return -1, -1, errors.New("Not a valid round - expect exactly 2 actions")
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
	score_1, err := outcome_score_p1(action_codes)
	if err != nil {
		return -1, -1, err
	}
	score_2, err := outcome_score_p2(action_codes)
	if err != nil {
		return -1, -1, err
	}

	return score_1, score_2, nil
}

func outcome_score_p1(actions [2]int) (int, error) {
	your_action := actions[1]
	opp_action := actions[0]
	err := action_range_checker(your_action, opp_action)
	if err != nil {
		return -1, err
	}
	outcome := [3]int{3, 6, 0}
	outcome_index := your_action - opp_action
	result := circular_array(outcome, outcome_index)
	return result + your_action + 1, nil
}

func action_range_checker(action_1 int, action_2 int) error {
	if action_1 < 0 || action_1 > 2 || action_2 < 0 || action_2 > 2 {
		return errors.New("Action Code out of Range")
	}
	return nil
}

func outcome_score_p2(actions [2]int) (int, error) {
	outcome := actions[1]
	opp_action := actions[0]
	err := action_range_checker(outcome, opp_action)
	if err != nil {
		return -1, err
	}
	if outcome == 1 {
		return opp_action + 3 + 1, nil
	}
	your_action_choice := [3]int{0, 1, 2}
	your_action_index := opp_action + 2 - (outcome / 2)
	your_action := circular_array(your_action_choice, your_action_index)
	return outcome*3 + your_action + 1, nil
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	per_round := strings.Split(string(content), "\n")

	total_score_p1 := 0
	total_score_p2 := 0
	for _, round := range per_round {
		score_1, score_2, err := score_round(round)
		if err != nil {
			log.Fatal(err)
		}
		total_score_p1 += score_1
		total_score_p2 += score_2
	}
	fmt.Printf("Total Score (Part 1): %d\nTotal Score (Part 2): %d", total_score_p1, total_score_p2)

}
