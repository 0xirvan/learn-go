package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Custom error

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("input is empty")
	}
	return true, nil

}

func main() {
	var input string
	fmt.Println("Enter numhber: ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println("You entered a number: ", number)
	} else {
		fmt.Println("You entered a string: ", input)
		fmt.Println(err.Error())
	}

	isOk, err := validate("")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Is OK: ", isOk)
	}
}
