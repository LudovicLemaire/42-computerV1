package main

import (
	"fmt"
	"math"
)

func degreeTwo(a float64, b float64, c float64) {
	delta := (b * b) - 4*(a*c)
	fmt.Printf("\n%sPolynomial degree:%s 2", string(colorB), string(colorReset))
	fmt.Printf("\n%sΔ:%s ", string(colorT), string(colorReset))
	fmt.Printf("%g\n", delta)

	// Δ > 0
	if delta > 0 {
		x1 := ((-b) + math.Sqrt(delta)) / (2 * a)
		x2 := ((-b) - math.Sqrt(delta)) / (2 * a)
		fmt.Println("\nDiscriminant is strictly positive, the two Real solutions are:")
		fmt.Printf("%sx1:%s %g\n", string(colorG), string(colorReset), x1)
		fmt.Printf("%sx2:%s %g\n", string(colorG), string(colorReset), x2)
	}

	// Δ = 0
	if delta == 0 {
		x := -b / (2 * a)
		fmt.Println("\nDiscriminant is null, the one Real solution is:")
		fmt.Printf("%sx:%s %g\n", string(colorG), string(colorReset), x)
	}

	// Δ < 0
	if delta < 0 {
		x1 := fmt.Sprintf("%g + i√%+g / %g\n", -b, delta, 2*a)
		x2 := fmt.Sprintf("%g - i√%+g / %g\n", -b, delta, 2*a)
		fmt.Println("\nDiscriminant is negative, the equation has no real solution, but has 2 complex solutions:")
		fmt.Printf("%sx1:%s %s", string(colorG), string(colorReset), x1)
		fmt.Printf("%sx2:%s %s\n", string(colorG), string(colorReset), x2)
	}
}

func degreeOne(x float64, v float64) {
	fmt.Printf("\n%sPolynomial degree:%s 1\nThe solution is:\n", string(colorB), string(colorReset))
	fmt.Printf("%sx:%s %g\n", string(colorG), string(colorReset), v/x)
}

func degreeZero(v float64) {
	if v == 0 {
		fmt.Println("\nSolutions are every reel numbers.")
	} else {
		fmt.Println("\nI am pretty sure that what you wrote does not mean anything mathematicaly speaking.")
	}
}
