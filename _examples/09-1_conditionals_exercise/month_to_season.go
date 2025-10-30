package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		printUsage()
		os.Exit(1)
	}
	monthNStr := os.Args[1]
	monthNum, err := strconv.Atoi(monthNStr)
	if err != nil {
		fmt.Printf("Error: Invalid month value '%s'. Must be an integer.\n", monthNStr)
		os.Exit(1)
	}

	season, err := Season(monthNum)
	if err != nil {
		fmt.Printf("Error: Invalid month value '%d'. Must be a number from 1 to 12.\n", monthNum)
		os.Exit(1)
	}
	fmt.Printf("%s is in %s\n", time.Month(monthNum).String(), season)
}

func Season(month int) (string, error) {
	switch month {
	case 12, 1, 2:
		return "Winter", nil
	case 3, 4, 5:
		return "Spring", nil
	case 6, 7, 8:
		return "Summer", nil
	case 9, 10, 11:
		return "Autumn", nil
	default:
		return "", fmt.Errorf("parameter 'month' value needs to be from 1 to 12")
	}
}

func printUsage() {
	fmt.Println("Usage: go run month_to_season.go <month>")
	fmt.Println("  <month> : The month number.")
}
