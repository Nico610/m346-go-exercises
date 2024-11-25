package main

import "fmt"

type FullName struct {
	Firstname string
	Lastname  string
}

type Birthdate struct {
	DayOfBirth   byte
	MonthOfBirth byte
	YearOfBirth  int16
}

type Profile struct {
	Name             FullName
	Born             Birthdate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{

		Name: FullName{
			Firstname: "Nico",
			Lastname:  "Marcuard",
		},
		Born: Birthdate{
			DayOfBirth:   9,
			MonthOfBirth: 1,
			YearOfBirth:  2007,
		},
		NumberOfSiblings: 0,
		ZodiacSign:       '\u2651',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings = me.NumberOfSiblings + 1
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
