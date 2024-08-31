package main

import (
	"fmt"
)

func romanToInteger(origin_roman string) int {
	var roman map[string]int = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	
	var result int = 0
	var length_of_roman int = len(origin_roman)

	for i := 0; i < length_of_roman; {
		var number1 int = roman[string(origin_roman[i])]

		if (i + 1 < length_of_roman) {
			var number2 int = roman[string(origin_roman[i + 1])]
			if (number1 >= number2){
				result += number1
				i++
			} else {
				result += number2 - number1
				i += 2
			}
		} else {
			result += number1
			i++
		}
	}
	return result
}

func main() {
	var roman string = "MCMXCIV"
	var arabic int = romanToInteger(roman)
	fmt.Print(arabic)
}