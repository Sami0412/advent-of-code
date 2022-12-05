package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
 
	if err != nil {
		fmt.Printf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	var caloriesPerElf []int = make([]int, len(txtlines))
	var elfIndex int
	var mostCalories int
	var elfWithMostCalories int
	var secondMostCalories int
	var elfWithSecondMostCalories int
	var thirdMostCalories int
	var elfWithThirdMostCalories int

	for i, line := range txtlines {
		if txtlines[i] == "" {
			elfIndex++
			continue
		}
		
		itemCalories, err := strconv.Atoi(line)
		
		if err != nil {
			fmt.Println(err)
		}

		caloriesPerElf[elfIndex] += int(itemCalories)

		if caloriesPerElf[elfIndex] > mostCalories {
			mostCalories = caloriesPerElf[elfIndex]
			elfWithMostCalories = elfIndex
		}
	}

	fmt.Println(fmt.Sprintf("Elf %d has the most calories: %d",elfWithMostCalories, mostCalories))

	for j, elfCalories := range caloriesPerElf {
		if elfCalories < mostCalories && elfCalories > secondMostCalories {
			secondMostCalories = elfCalories
			elfWithSecondMostCalories = j
		}
	}

	fmt.Println(fmt.Sprintf("Elf %d has the second most calories: %d",elfWithSecondMostCalories, secondMostCalories))

	for k, elfCalories := range caloriesPerElf {
		if elfCalories < mostCalories && elfCalories < secondMostCalories && elfCalories > thirdMostCalories {
			thirdMostCalories = elfCalories
			elfWithThirdMostCalories = k
		}
	}

	fmt.Println(fmt.Sprintf("Elf %d has the most calories: %d", elfWithThirdMostCalories, thirdMostCalories))

	totalCalories := mostCalories + secondMostCalories + thirdMostCalories

	fmt.Println(totalCalories)
	
	file.Close()
}