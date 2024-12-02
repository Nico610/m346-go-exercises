package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		KlassenBezeichnug string
		Students          []Student
	}
	// TODO: declare a map of modules being attended by multiple classes

	modules := map[string][]Class{
		"M346": {
			{
				KlassenBezeichnug: "INA23bl",
				Students: []Student{
					{FirstName: "Nico", LastName: "Marcuard"},
					{FirstName: "Steven", LastName: "Mattmann"},
				},
			},
			{
				KlassenBezeichnug: "INA24BL",
				Students: []Student{
					{FirstName: "Yves", LastName: "Troxler"},
					{FirstName: "Domenic", LastName: "Troxler"},
					{FirstName: "Riccardo", LastName: "Greco"},
				},
			},
		},
		"Mathematik": {
			{
				KlassenBezeichnug: "INF92LA",
				Students: []Student{
					{FirstName: "Elia", LastName: "Mustermann"},
					{FirstName: "Elisa", LastName: "Musterfrau"},
				},
			},
			{
				KlassenBezeichnug: "INF93LAC",
				Students: []Student{
					{FirstName: "Gabriel", LastName: "Buqaj"},
					{FirstName: "Dominik", LastName: "Scharer"},
					{FirstName: "Tiago", LastName: "Caroso"},
				},
			},
		},
	}

	// TODO: output everything using fmt.Println()
	fmt.Println(modules)
}
