package main

import (
	"fmt"
	"strings"
)

func main() {
	palindrom("radar")
	maxNumber([]int{9, 2, 4})
	star(5)
	factorial(4)
}

// 1.
func palindrom(text string) string {

	textSplit := strings.Split(text, "")
	textSplitArray1 := []string{}
	textSplitArray2 := []string{}

	for i := len(textSplit) - 1; i >= 0; i-- {
		textSplitArray2 = append(textSplitArray2, textSplit[i])
	}
	for _, value := range textSplit {
		textSplitArray1 = append(textSplitArray1, value)
	}
	if strings.Join(textSplitArray1, "") != strings.Join(textSplitArray2, "") {
		fmt.Println("false")
		return "false"
	}
	fmt.Println("true")
	return "true"

}

// 2
func maxNumber(number []int) int {

	max := 0
	for _, value := range number {
		if max < value {
			max = value
		}
	}
	fmt.Println(max)
	return max
}

// 3
func star(input int) string {
	asterix := "*"
	result := ""
	for i := 0; i <= (input); i++ {
		fmt.Println(strings.Repeat(asterix, i))
		result = strings.Repeat(asterix, i)
	}
	fmt.Println(result)

	return result
}

// no 4
func factorial(input int) int {
	result := 0
	if input == 0 {
		fmt.Println(1)
		return 1
	} else {
		result = input * factorial(input-1)
	}
	fmt.Println(result)

	return result
}
