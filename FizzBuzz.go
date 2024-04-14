package main

import (
	"fmt"
)

func FizzBuzz(number int) string {
	if number%3 == 0 && number%5 == 0 {
		return "FizzBuzz"
	} else if number%3 == 0 {
		return "Fizz"
	} else if number%5 == 0 {
		return "Buzz"
	} else {
		return fmt.Sprintf("%d", number)
	}
}
