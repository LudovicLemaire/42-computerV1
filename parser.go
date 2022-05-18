package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func spacetrim(str string) string {
	return strings.Join(strings.Fields(str), "")
}

var degreeMap, degreeMap2 map[float64]float64

func parser() (float64, float64, float64) {
	degreeMap = make(map[float64]float64, len(degreeMap))
	degreeMap2 = make(map[float64]float64, len(degreeMap2))

	s1 := ""
	if len(os.Args) == 2 {
		//init equation
		s1 = os.Args[1]
		if s1 != "" {
			//print formula
			s := fmt.Sprintf("%sFormula:%s %s\n", string(colorB), string(colorReset), s1)
			fmt.Println(s)
		}
	}
	for s1 == "" {
		s1 = "4 * X^2 + 4 * X^1 + 1 * X^0 = 0" // Δ = 0
		//s1 = "-5 * X^0 - 9 * X^1 + 2 * X^2 = 0 * X^0" // Δ > 0
		//s1 = "7 * X^0 + 5 * X^1 + 3 * X^2 = 0" // Δ < 0
		//s1 = "2x^2 + 9x - 5 + 4 - 4 = 0" // Δ > 0 human readable
		//s1 = "2x^2 + 9x - 5 + 4 - 4 + 4x^5 = 0" // error too high degree
		//s1 = "2x^2 + 9x - 5 + 4 - 4 + 4x^2 = 0 + 5 + 7 + 9 + 45 + 4x + 9x" // random shit
		//s1 = "1 + 4 + 74 + 8 - 45 - 18 + 4 - 41 + x - 3x = 25x + 4 - 9 - 2x + 12 - 15x" // degree one
		//s1 = "42 = 42" // mehh
		fmt.Printf("%sUsage examples:%s\n", string(colorY), string(colorReset))
		fmt.Println("-5 * X^0 - 9 * X^1 + 2 * X^2 = 0 * X^0")
		fmt.Println("4 * X^2 + 4 * X^1 + 1 * X^0 = 0")
		fmt.Printf("7 * X^0 + 5 * X^1 + 3 * X^2 = 0\n\n")
		fmt.Println("2x^2 + 9x - 5 = 0")
		fmt.Println("2x^2 + 9x - 5 + 4 - 4 + 4x^2 = 0 + 5 + 7 + 9 + 45 + 4x + 9x")
		fmt.Println("1 + 4 + 74 + 8 - 45 - 18 + 4 - 41 + x - 3x = 25x + 4 - 9 - 2x + 12 - 15x")

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("\n%sEnter your equation: %s", string(colorB), string(colorReset))
		s1, _ = reader.ReadString('\n')
		if len(s1) == 2 {
			s1 = ""
		}
		fmt.Println("")
	}
	//
	// Split
	//

	//reorder string
	s1 = strings.Replace(s1, " * ", "*", -1)
	s1 = strings.Replace(s1, " + ", " +", -1)
	s1 = strings.Replace(s1, " - ", " -", -1)

	//cut the equation in 2 parts
	rmEqual := regexp.MustCompile(`=`)
	splited := rmEqual.Split(s1, -1)

	if len(splited) != 2 {
		fmt.Printf("\n%sError parsing:%s formula is wrong\n", string(colorR), string(colorReset))
		os.Exit(42)
	} else if len(splited[0]) == 0 || len(splited[1]) == 0 {
		fmt.Printf("\n%sError parsing:%s formula is wrong\n", string(colorR), string(colorReset))
		os.Exit(42)
	}

	//cut per part
	firstPart := strings.Split(strings.TrimSpace(splited[0]), ` `)
	secondPart := strings.Split(strings.TrimSpace(splited[1]), ` `)

	//trim parts
	for i := range firstPart {
		firstPart[i] = strings.TrimSpace(firstPart[i])
	}
	for i := range secondPart {
		secondPart[i] = strings.TrimSpace(secondPart[i])
	}

	if len(firstPart) == 0 || len(secondPart) == 0 {
		fmt.Printf("\n%sError parsing:%s formula is wrong\n", string(colorR), string(colorReset))
		os.Exit(42)
	}

	if len(spacetrim(firstPart[0])) == 0 || len(spacetrim(secondPart[0])) == 0 {
		fmt.Printf("\n%sError parsing:%s formula is wrong\n", string(colorR), string(colorReset))
		os.Exit(42)
	}

	//add symbol for first part
	if !(firstPart[0][0] == '-' || firstPart[0][0] == '+') {
		firstPart[0] = "+" + firstPart[0]
	}
	if !(secondPart[0][0] == '-' || secondPart[0][0] == '+') {
		secondPart[0] = "+" + secondPart[0]
	}

	for i := range firstPart {
		firstPart[i] = strings.ToLower(firstPart[i])
		if strings.Contains(firstPart[i], "*x^") {
			v, _ := strconv.ParseFloat(strings.Split(firstPart[i], "*")[0], 64)
			degree, _ := strconv.ParseFloat(strings.Split(firstPart[i], "^")[1], 64)
			degreeMap[degree] += v
		} else if strings.Contains(firstPart[i], "x^") {
			v, _ := strconv.ParseFloat(strings.Split(firstPart[i], "x")[0], 64)
			degree, _ := strconv.ParseFloat(strings.Split(firstPart[i], "^")[1], 64)
			if v == 0 && len(strings.Split(firstPart[i], "x")[0]) == 1 {
				v = 1
			}
			degreeMap[degree] += v
		} else if strings.Contains(firstPart[i], "x") {
			v, _ := strconv.ParseFloat(strings.Split(firstPart[i], "x")[0], 64)
			var degree float64 = 1
			if v == 0 && len(strings.Split(firstPart[i], "x")[0]) == 1 {
				v = 1
			}
			degreeMap[degree] += v
		} else {
			v, _ := strconv.ParseFloat(firstPart[i], 64)
			var degree float64 = 0
			degreeMap[degree] += v
		}
	}

	for i := range secondPart {
		secondPart[i] = strings.ToLower(secondPart[i])
		if strings.Contains(secondPart[i], "*x^") {
			v, _ := strconv.ParseFloat(strings.Split(secondPart[i], "*")[0], 64)
			degree, _ := strconv.ParseFloat(strings.Split(secondPart[i], "^")[1], 64)
			degreeMap2[degree] += v
		} else if strings.Contains(secondPart[i], "x^") {
			v, _ := strconv.ParseFloat(strings.Split(secondPart[i], "x")[0], 64)
			degree, _ := strconv.ParseFloat(strings.Split(secondPart[i], "^")[1], 64)
			if v == 0 && len(strings.Split(secondPart[i], "x")[0]) == 1 {
				v = 1
			}
			degreeMap2[degree] += v
		} else if strings.Contains(secondPart[i], "x") {
			v, _ := strconv.ParseFloat(strings.Split(secondPart[i], "x")[0], 64)
			var degree float64 = 1
			if v == 0 && len(strings.Split(secondPart[i], "x")[0]) == 1 {
				v = 1
			}
			degreeMap2[degree] += v
		} else {
			v, _ := strconv.ParseFloat(secondPart[i], 64)
			var degree float64 = 0
			degreeMap2[degree] += v
		}
	}

	//
	// print readable
	//
	fmt.Printf("%sHuman readable:%s\n", string(colorP), string(colorReset))
	//print first part
	for key, element := range degreeMap {
		if key == 0 {
			fmt.Printf("%+g ", element)
		} else if key == 1 {
			fmt.Printf("%+gx ", element)
		} else {
			fmt.Printf("%+gx^%g ", element, key)
		}
	}
	fmt.Printf(" = ")
	//print second part
	for key, element := range degreeMap2 {
		if key == 0 {
			fmt.Printf("%+g ", element)
		} else if key == 1 {
			fmt.Printf("%+gx ", element)
		} else {
			fmt.Printf("%+gx^%g ", element, key)
		}
	}

	//print reduced form
	fmt.Printf("\n\n%sReduced form:%s\n", string(colorT), string(colorReset))
	for key, element := range degreeMap2 {
		degreeMap[key] -= element
		degreeMap2[key] = 0
	}
	for key, element := range degreeMap {
		if key == 0 && element != 0 {
			fmt.Printf("%+g ", element)
		} else if key == 1 && element != 0 {
			fmt.Printf("%+gx ", element)
		} else if element != 0 {
			fmt.Printf("%+gx^%g ", element, key)
		}
	}
	fmt.Printf(" = 0\n")

	var errorDegree []float64
	for key, element := range degreeMap {
		if key != 0 && key != 1 && key != 2 {
			if element != 0 {
				errorDegree = append(errorDegree, key)
			}
		}
	}
	if len(errorDegree) > 0 {
		fmt.Printf("\n%sError:%s Polynomial degree: %g\n", string(colorR), string(colorReset), errorDegree)
		fmt.Print("I can't solve polynomial degree outside of [0, 1, 2]. ")
		os.Exit(42)
	}

	return degreeMap[2], degreeMap[1], degreeMap[0]
}
