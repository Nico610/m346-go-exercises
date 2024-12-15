package main

import (
	"fmt"
	"math"
)

func computeQuadraticFormula(a, b, c float64) (float64, float64, error) {
	diskriminant := math.Pow(b, 2) - 4*a*c

	if diskriminant < 0 {
		return 0, 0, fmt.Errorf("Fehler Ausgabe")
	}

	x1 := (-b + math.Sqrt(diskriminant)) / (2 * a)
	x2 := (-b - math.Sqrt(diskriminant)) / (2 * a)

	return x1, x2, nil
}

func main() {
	fmt.Println(computeQuadraticFormula(1, 5, 6))
}
