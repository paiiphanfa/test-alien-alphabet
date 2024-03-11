package main

import (
	"fmt"
	"strings"
)

func main() {
	alienValues := map[rune]int{
		'R': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'Z': 10,
		'B': 5,
		'A': 1,
	}
	var input string

	fmt.Print("input: s = ")
	fmt.Scanln(&input)

	input = strings.ToUpper(input)

	var sum int
	var alienNum []int
	var explanation []string

	for i := 0; i < len(input); i++ {
		char := rune(input[i])
		nextChar := ' '

		if i < len(input)-1 {
			nextChar = rune(input[i+1])
		}

		switch char {
		case 'A':
			if nextChar == 'B' || nextChar == 'Z' {
				alienNum = append(alienNum, alienValues[nextChar]-alienValues[char])
				explanation = append(explanation, fmt.Sprintf("%s%s = %d", string(char), string(nextChar), alienValues[nextChar]-alienValues[char]))
				i++
			} else {
				alienNum = append(alienNum, alienValues[char])
				explanation = append(explanation, fmt.Sprintf("%s = %d", string(char), alienValues[char]))
			}
		case 'Z':
			if nextChar == 'L' || nextChar == 'C' {
				alienNum = append(alienNum, alienValues[nextChar]-alienValues[char])
				explanation = append(explanation, fmt.Sprintf("%s%s = %d", string(char), string(nextChar), alienValues[nextChar]-alienValues[char]))
				i++
			} else {
				alienNum = append(alienNum, alienValues[char])
				explanation = append(explanation, fmt.Sprintf("%s = %d", string(char), alienValues[char]))
			}
		case 'C':
			if nextChar == 'D' || nextChar == 'R' {
				alienNum = append(alienNum, alienValues[nextChar]-alienValues[char])
				explanation = append(explanation, fmt.Sprintf("%s%s = %d", string(char), string(nextChar), alienValues[nextChar]-alienValues[char]))
				i++
			} else {
				alienNum = append(alienNum, alienValues[char])
				explanation = append(explanation, fmt.Sprintf("%s = %d", string(char), alienValues[char]))
			}
		default:
			alienNum = append(alienNum, alienValues[char])
			explanation = append(explanation, fmt.Sprintf("%s = %d", string(char), alienValues[char]))
		}
	}

	for _, value := range alienNum {
		sum += value
	}

	fmt.Printf("Output: %d\n", sum)
	fmt.Printf("Explanation: %s.\n", strings.Join(explanation, ", "))
}
