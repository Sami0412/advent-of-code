package main

import (
	"fmt"
	"bufio"
	"os"
)

var roundScoreMap = map[string]int{
	"A Y": 6,
	"B Z": 6,
	"C X": 6,
	"A Z": 0,
	"B X": 0,
	"C Y": 0,
	"A X": 3,
	"B Y": 3,
	"C Z": 3,
}

var handScoreMap = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func main() {
	file, err := os.Open("input.txt")
 
	if err != nil {
		fmt.Printf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var roundResult string

	for scanner.Scan() {
		roundResult = scanner.Text()
	}
}