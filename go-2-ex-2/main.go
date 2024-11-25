package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	fibs[2] = fibs[0] + fibs[1]
	fibs[3] = fibs[1] + fibs[2]
	fibs[4] = fibs[2] + fibs[3]

	fibs = append(fibs, fibs[3]+fibs[4]) // Hier rechne ich die 3 Zahl und die 4 Zahl zussamenn um die nächste zu bekommen

	fibs = append(fibs, fibs[len(fibs)-2]+fibs[len(fibs)-1]) // muss 13 sein
	fibs = append(fibs, fibs[len(fibs)-2]+fibs[len(fibs)-1]) // muss 21 sein
	fibs = append(fibs, fibs[len(fibs)-2]+fibs[len(fibs)-1]) // muss 34 sein

	fmt.Println(fibs) // expected output: [1 1 2 3 5 8 13 21 34]
}
