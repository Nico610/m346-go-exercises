package main

import (
	"fmt"
)

func computeGrade(gradePoints int) int {
	if gradePoints >= 90 {
		return 6
	} else if gradePoints >= 70 && gradePoints < 90 {
		return 5
	} else if gradePoints >= 50 && gradePoints < 70 {
		return 4
	} else if gradePoints >= 30 && gradePoints < 50 {
		return 3
	} else if gradePoints >= 20 && gradePoints < 30 {
		return 2
	} else if gradePoints >= 0 && gradePoints < 20 {
		return 1
	}
	return 1
}

func main() {
	//Mit gegeben wird die Punktzahl
	fmt.Println(computeGrade(93)) //6
	fmt.Println(computeGrade(83)) //5
	fmt.Println(computeGrade(64)) //4
	fmt.Println(computeGrade(39)) //3
	fmt.Println(computeGrade(20)) //2
	fmt.Println(computeGrade(2))  //1
}
