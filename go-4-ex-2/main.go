package main

import (
	"fmt"
	"math"
)

func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	fmt.Println(computeHypotenuse(4, 5))
	fmt.Println(computeHypotenuse(8, 10))
	fmt.Println(computeHypotenuse(34, 51))
	fmt.Println(computeHypotenuse(7, 5.75))

}
