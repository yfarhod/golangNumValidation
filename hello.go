package main

import (
	"fmt"
)

func main() {

	fmt.Println("Please enter your 16-digit credit card number:")

	var userInput int
	fmt.Scanln(&userInput)

	if checkNum(userInput) {
		fmt.Println("Congratulations! Your credit card number:", userInput, "is valid")
	} else {
		fmt.Println("Sorry! Your credit card number:", userInput, "is not valid")
	}

}

func checkNum(num int) bool {
	if checkLength(num) != 16 {
		return false
	}

	var arr [16]int
	for i := 0; i < 16; i++ {
		arr[i] = num % 10
		num = num / 10
	}

	for j := 1; j < 16; j += 2 {
		arr[j] = arr[j] * 2
	}

	sum := 0
	for _, v := range arr {
		sum += (v / 10) + (v % 10)
	}

	if sum%10 == 0 {
		return true
	}

	return false

}

func checkLength(num int) int {
	count := 0
	for num != 0 {
		num = num / 10
		count++
	}
	return count
}
