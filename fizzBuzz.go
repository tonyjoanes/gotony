package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 1
	for i <= 15 {
		fmt.Println(fizzBuzz(i))
		i = i + 1
	}
}

func fizzBuzz(testNumber int) string {
	var returnString = ""

	if divisibleByThree(testNumber) {
		returnString += "Fizz"
	}

	if divisibleByFive(testNumber) {
		returnString += "Buzz"
	}

	if len(returnString) > 0 {
		return returnString
	}

	return strconv.Itoa(testNumber)
}

func divisibleByThree(testNum int) bool {
	return testNum%3 == 0
}

func divisibleByFive(testNum int) bool {
	return testNum%5 == 0
}
