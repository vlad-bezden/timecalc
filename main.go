// Package main calculates time for user provided distance and speed
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	minute = 60          // minutes in an hour
	second = minute * 60 // seconds in an hour
)

var errInvalidArguments = errors.New("Invalid arguments")
var errReadingInput = errors.New("Error reading input")

func main() {
	for {
		miles := getInput("Enter distance (miles): ")
		speed := getInput("Enter speed (m/h): ")

		calcResult(miles, speed)
	}
}

// calcResult calculate hours, minutes, and seconds and print them out
func calcResult(miles float64, speed float64) {
	// time in hours
	time := miles / speed
	hours := int(time)
	minutes := int(time*minute) % 60
	seconds := int(time*second) % 60

	fmt.Printf("Time: %02d:%02d:%02d\n\n", hours, minutes, seconds)
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

// getInput prompts user for speed or milage input validate it and returns value
func getInput(prompt string) float64 {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		if err == io.EOF {
			fmt.Println("\n\nExiting...")
			os.Exit(0)
		}
		fmt.Println(err)
		printError(errReadingInput)
	}
	// convert string to float64 value
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		printError(errInvalidArguments)
	}
	return value
}
