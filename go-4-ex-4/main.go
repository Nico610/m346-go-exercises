package main

import "fmt"

func convertCelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {

	fmt.Println("Erste Umrechnung von 30,5 Celsius in Fahrenheit", convertCelsiusToFahrenheit(30.5))
	fmt.Println("Zweite Umrechnung von 100 Fahrenheit in Celsius", convertFahrenheitToCelsius(100))

}
