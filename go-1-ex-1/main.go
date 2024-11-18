package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!

	//Print für Vor- und Nachname
	var firstName string = "Nico"
	var lastName string = "Marcuard"

	//Print für Geburtsdatum
	var dayOfBirth int = 9
	var monthOfBirth int = 1
	var yearOfBirth int = 2007

	//Print für Geschwister
	var numberOfSiblings int = 0

	//Print für Grösse
	var heightInMeters float32 = 1.88

	//Print für Strenzeichen
	var zodiacSign rune = '\u2651'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
