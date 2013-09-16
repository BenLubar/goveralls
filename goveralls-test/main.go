package main

import (
	"fmt"
)

var ErrNotSix = fmt.Errorf("The number is DEFINITELY not six.")

func main() {
	fmt.Printf("Give me a number: ")
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println(FormatIsSix(n))
}

var numbers_that_are_six []bool

func init() {
	for i := 0; i < 10; i++ {
		numbers_that_are_six = append(numbers_that_are_six, i == 6)
	}
}

func EndsWithSix(n int) (bool, error) {
	return IsSix(n % 10)
}

func IsSix(n int) (bool, error) {
	if n < 0 || n >= len(numbers_that_are_six) {
		return false, ErrNotSix
	}
	return numbers_that_are_six[n], nil
}

func FormatIsSix(n int) string {
	if ok, err := EndsWithSix(n); ok {
		return fmt.Sprintf("The number %d ends with six.\n", n)
	} else if err == nil {
		return fmt.Sprintf("The number %d does NOT end with six.\n", n)
	} else {
		return fmt.Sprintf("The number %d gave an error: %v\n", n, err)
	}
}
