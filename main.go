package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var errInvalidArguments = errors.New("Invalid arguments")
var errReadingInput = errors.New("Error reading input")

func main() {
	for {
		miles := getInput("Enter Miles: ")
		speed := getInput("Enter Speed (m/h): ")

		calcResult(miles, speed)

		isContinue := askForContinue()

		if !isContinue {
			fmt.Println("Good bye!")
			break
		}
	}
}

func calcResult(miles float64, speed float64) {
	// time in seconds
	time := miles / speed * 60 * 60
	hours := math.Floor(time / 60 / 60)
	minutes := math.Floor(time / 60)
	seconds := int(time) % 60

	fmt.Printf("%v:%v:%v\n\n", hours, minutes, seconds)
}

func askForContinue() bool {
	var response string
	fmt.Print("Would you like to calculate another time? (Y/y) ")

	_, err := fmt.Scanln(&response)
	if err != nil {
		printError(errReadingInput)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}

	if containsString(okayResponses, response) {
		return true
	}
	return false
}

// containsString returns true if slice contains element
func containsString(slice []string, element string) bool {
	for _, elem := range slice {
		if elem == element {
			return true
		}
	}
	return false
}

// printError prints error and exit application with status 1
func printError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}

func getInput(prompt string) float64 {
	fmt.Print(prompt)
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		printError(errReadingInput)
	}
	return inputToValue(input)
}

func inputToValue(input string) float64 {
	// trim the newline charecter
	input = strings.TrimSpace(input)
	// convert string to float64 value
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		printError(errInvalidArguments)
	}
	return value
}
