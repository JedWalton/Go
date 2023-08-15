package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var (
	figuresCodes = map[string]string{
		"A": "R",
		"B": "P",
		"C": "S",
		"X": "L",
		"Y": "D",
		"Z": "W",
	}
)

func calculateScore(a, b string) int {
	p1, p2 := figuresCodes[a], figuresCodes[b]

	roundScore := 0
	if p2 == "D" { // draw
		roundScore += 3
		if p1 == "R" {
			roundScore += 1 // I must pick rock
		} else if p1 == "P" {
			roundScore += 2 // I must pick paper
		} else if p1 == "S" {
			roundScore += 3 // I must pick scissors
		}
	} else if p2 == "L" { // lose
		roundScore += 0
		if p1 == "R" {
			roundScore += 3 // I must pick scissors
		} else if p1 == "P" {
			roundScore += 1 // I must pick rock
		} else if p1 == "S" {
			roundScore += 2 // I must pick paper
		}
	} else { // win
		roundScore += 6
		if p1 == "R" {
			roundScore += 2 // I must pick paper
		} else if p1 == "P" {
			roundScore += 3 // I must pick scissors
		} else if p1 == "S" {
			roundScore += 1 // I must pick rock
		}
	}

	return roundScore
}

func main() {
	input, err := os.Open("./strategy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	totalScore := 0

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		totalScore += calculateScore(tokens[0], tokens[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(totalScore)
}
