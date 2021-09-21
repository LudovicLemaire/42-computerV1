package main

import (
	"fmt"
	"math"
)

// Newton’s method
func Sqrt(x float64) float64 {
	var z float64 = 1
	for i := 1; i <= 10; i++ {
		z = (z - (z*z-x)/(2*z))
	}
	return z
}

func degreeTwo(a float64, b float64, c float64) {
	delta := (b * b) - 4*(a*c)
	fmt.Printf("\n%sPolynomial degree:%s 2", string(colorB), string(colorReset))
	fmt.Printf("\n%sΔ:%s ", string(colorT), string(colorReset))
	fmt.Printf("%g\n", delta)

	// Δ > 0
	if delta > 0 {
		x1 := ((-b) + Sqrt(delta)) / (2 * a)
		x2 := ((-b) - Sqrt(delta)) / (2 * a)
		fmt.Println("\nDiscriminant is strictly positive, the two Real solutions are:")
		fmt.Printf("%sx1:%s %g", string(colorG), string(colorReset), x1)
		v, div := floatToFrac(x1)
		if !(div == 1.0 || int(div)%10 == 0) {
			fmt.Printf(" %s:%s %g/%g", string(colorG), string(colorReset), v, div)
		}
		fmt.Println("")
		fmt.Printf("%sx2:%s %g", string(colorG), string(colorReset), x2)
		v, div = floatToFrac(x2)
		if !(div == 1.0 || int(div)%10 == 0) {
			fmt.Printf(" %s:%s %g/%g", string(colorG), string(colorReset), v, div)
		}
		fmt.Println("")
	}

	// Δ = 0
	if delta == 0 {
		x := -b / (2 * a)
		fmt.Println("\nDiscriminant is null, the one Real solution is:")
		fmt.Printf("%sx:%s %g", string(colorG), string(colorReset), x)
		v, div := floatToFrac(x)
		if !(div == 1.0 || int(div)%10 == 0) {
			fmt.Printf(" %s:%s %g/%g\n", string(colorG), string(colorReset), v, div)
		}
		fmt.Println("")
	}

	// Δ < 0
	if delta < 0 {

		// α β
		//x1 := fmt.Sprintf("(%g + i√%+g) / %g\n", -b, -delta, 2*a)
		//x2 := fmt.Sprintf("(%g - i√%+g) / %g\n", -b, -delta, 2*a)

		x1 := fmt.Sprintf("%g - i%g\n", -b/(2*a), Sqrt(-delta)/(2*a))
		x2 := fmt.Sprintf("%g + i%g\n", -b/(2*a), Sqrt(-delta)/(2*a))

		//x1 := fmt.Sprintf("%g\n", (-b/2*a)-(Sqrt(-delta)/2*a))
		//x2 := fmt.Sprintf("%g\n", (-b/2*a)+(Sqrt(-delta)/2*a))

		fmt.Println("\nDiscriminant is negative, the equation has no real solution, but has 2 complex solutions:")
		fmt.Printf("%sx1:%s α - iβ\n", string(colorT), string(colorReset))
		fmt.Printf("%sx2:%s α + iβ\n\n", string(colorT), string(colorReset))

		fmt.Printf("%sx1:%s %s\n", string(colorT), string(colorReset), "(-b / 2a) - (i√(Δ) / 2a)")
		fmt.Printf("%sx2:%s %s\n\n", string(colorT), string(colorReset), "(-b / 2a) + (i√(Δ) / 2a)")

		fmt.Printf("%sx1:%s -%g/2*%g - i√(%g)/2*%g\n", string(colorT), string(colorReset), b, a, delta, a)
		fmt.Printf("%sx1:%s -%g/2*%g + i√(%g)/2*%g\n\n", string(colorT), string(colorReset), b, a, delta, a)

		fmt.Printf("%sx1:%s -%g/%g - i√(%g)/%g\n", string(colorT), string(colorReset), b, 2*a, delta, 2*a)
		fmt.Printf("%sx1:%s -%g/%g + i√(%g)/%g\n\n", string(colorT), string(colorReset), b, 2*a, delta, 2*a)

		fmt.Printf("%sx1:%s %s", string(colorG), string(colorReset), x1)
		fmt.Printf("%sx2:%s %s\n", string(colorG), string(colorReset), x2)

		var pt1 string
		var pt2 string
		v, div := floatToFrac(-b / (2 * a))
		if !(div == 1.0 || int(div)%10 == 0) {
			pt1 = fmt.Sprintf("%g/%g", v, div)
		} else {
			pt1 = fmt.Sprintf("%g", -b/2*a)
		}
		v, div = floatToFrac(Sqrt(-delta) / (2 * a))
		if !(div == 1.0 || int(div)%10 == 0) {
			pt2 = fmt.Sprintf("%g/%g", v, div)
		} else {
			pt2 = fmt.Sprintf("%g", Sqrt(-delta)/(2*a))
		}

		fmt.Printf("%sx1:%s %s - i(%s)\n", string(colorP), string(colorReset), pt1, pt2)
		fmt.Printf("%sx1:%s %s + i(%s)\n", string(colorP), string(colorReset), pt1, pt2)

	}
}

func degreeOne(x float64, v float64) {
	fmt.Printf("\n%sPolynomial degree:%s 1\nThe solution is:\n", string(colorB), string(colorReset))
	fmt.Printf("%sx:%s %g", string(colorG), string(colorReset), v/x)
	v, div := floatToFrac(v / x)
	if !(div == 1.0 || int(div)%10 == 0) {
		fmt.Printf(" %s:%s %g/%g\n", string(colorG), string(colorReset), v, div)
	}
	fmt.Println("")
}

func degreeZero(v float64) {
	if v == 0 {
		fmt.Println("\nSolutions are every reel numbers.")
	} else {
		fmt.Println("\nI am pretty sure that what you wrote does not mean anything mathematicaly speaking.")
	}
}

func floatToFrac(x float64) (float64, float64) {
	isNeg := false
	if x < 0 {
		x = -x
		isNeg = true
	}
	tolerance := 1.0e-6
	h1 := 1.0
	k1 := 0.0
	h2 := 0.0
	k2 := 1.0
	b := x
	for {
		a := math.Floor(b)
		aux := h1
		h1 = a*h1 + h2
		h2 = aux
		aux = k1
		k1 = a*k1 + k2
		k2 = aux
		b = 1 / (b - a)
		if !(math.Abs(x-h1/k1) > x*tolerance) {
			break
		}
	}
	if isNeg {
		h1 = -h1
	}
	return h1, k1
}
