package main

import (
	"fmt"
	"os"
	"strconv"
)

type Celsius float32
type Fahrenheit float32

func main() {
	if len(os.Args) != 3 {
		printUsage()
	}

	tempStr := os.Args[1]
	temperature, err := strconv.ParseFloat(tempStr, 32)
	if err != nil {
		fmt.Printf("Error: Invalid temperature value '%s'. Must be a number.\n", tempStr)
		os.Exit(1)
	}

	flag := os.Args[2]

	switch flag {
	case "-f":
		result := toCelcius(Fahrenheit(temperature))
		fmt.Printf("%.2f°F is %.2f°C\n", temperature, result)
	case "-c":
		result := toFahrenheit(Celsius(temperature))
		fmt.Printf("%.2f°C is %.2f°F\n", temperature, result)
	default:
		fmt.Printf("Error: Unknown flag '%s'. Use -f or -c.\n", flag)
		printUsage()
	}

}

func toCelcius(t Fahrenheit) Celsius {
	return Celsius((t - 32) * 5 / 9)
}

func toFahrenheit(t Celsius) Fahrenheit {
	return Fahrenheit((t * 9 / 5) + 32)
}

func printUsage() {
	fmt.Println("Usage: go run temperature_converter.go <temperature> [-f|-c]")
	fmt.Println("  <temperature> : The numeric temperature value.")
	fmt.Println("  -f            : Convert temperature (Fahrenheit) to Celsius.")
	fmt.Println("  -c            : Convert temperature (Celsius) to Fahrenheit.")
	os.Exit(1)
}
