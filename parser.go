package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parser() (float64, float64, float64) {
	flag.Parse()
	s1 := flag.Arg(0)

	//init equation
	if s1 != "" {
		//print formula
		s := fmt.Sprintf("%sFormula:%s %s\n", string(colorB), string(colorReset), s1)
		fmt.Println(s)
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

	//add symbol for first part
	if !(firstPart[0][0] == '-' || firstPart[0][0] == '+') {
		firstPart[0] = "+" + firstPart[0]
	}
	if !(secondPart[0][0] == '-' || secondPart[0][0] == '+') {
		secondPart[0] = "+" + secondPart[0]
	}

	//init abc
	a := 0.0
	b := 0.0
	c := 0.0
	a2 := 0.0
	b2 := 0.0
	c2 := 0.0
	degreeError := 0.0

	for i := range firstPart {
		if strings.Contains(strings.ToLower(firstPart[i]), "*x^2") {
			v, _ := strconv.ParseFloat(strings.Split(firstPart[i], "*")[0], 64)
			if v == 0 && len(strings.Split(firstPart[i], "x")[0]) == 1 {
				v = 1
			}
			a += v
		} else if strings.Contains(strings.ToLower(firstPart[i]), "*x^1") {
			v, _ := strconv.ParseFloat(strings.Split(firstPart[i], "*")[0], 64)
			if v == 0 && len(strings.Split(firstPart[i], "x")[0]) == 1 {
				v = 1
			}
			b += v
		} else if strings.Contains(strings.ToLower(firstPart[i]), "*x^0") {
			v, _ := strconv.ParseFloat(strings.Split(firstPart[i], "*")[0], 64)
			c += v
		} else if strings.Contains(strings.ToLower(firstPart[i]), "x^2") {
			v, _ := strconv.ParseFloat(strings.Split(firstPart[i], "x")[0], 64)
			if v == 0 && len(strings.Split(firstPart[i], "x")[0]) == 1 {
				v = 1
			}
			a += v
		} else if strings.Contains(strings.ToLower(firstPart[i]), "x^") {
			v, _ := strconv.ParseFloat(strings.Split(firstPart[i], "^")[1], 64)
			if v == 0 && len(strings.Split(firstPart[i], "x")[0]) == 1 {
				v = 1
			}
			degreeError = v
		} else if strings.Contains(strings.ToLower(firstPart[i]), "x") {
			v, _ := strconv.ParseFloat(strings.Split(firstPart[i], "x")[0], 64)
			if v == 0 && len(strings.Split(firstPart[i], "x")[0]) == 1 {
				v = 1
			}
			b += v
		} else {
			v, _ := strconv.ParseFloat(firstPart[i], 64)
			c += v
		}
	}
	for i := range secondPart {
		if strings.Contains(strings.ToLower(secondPart[i]), "*x^2") {
			v, _ := strconv.ParseFloat(strings.Split(secondPart[i], "*")[0], 64)
			if v == 0 && len(strings.Split(firstPart[i], "x")[0]) == 1 {
				v = 1
			}
			a2 += v
		} else if strings.Contains(strings.ToLower(secondPart[i]), "*x^1") {
			v, _ := strconv.ParseFloat(strings.Split(secondPart[i], "*")[0], 64)
			if v == 0 && len(strings.Split(firstPart[i], "x")[0]) == 1 {
				v = 1
			}
			b2 += v
		} else if strings.Contains(strings.ToLower(secondPart[i]), "*x^0") {
			v, _ := strconv.ParseFloat(strings.Split(secondPart[i], "*")[0], 64)
			c2 += v
		} else if strings.Contains(strings.ToLower(secondPart[i]), "x^2") {
			v, _ := strconv.ParseFloat(strings.Split(secondPart[i], "x")[0], 64)
			if v == 0 && len(strings.Split(firstPart[i], "x")[0]) == 1 {
				v = 1
			}
			a2 += v
		} else if strings.Contains(strings.ToLower(secondPart[i]), "x^") {
			v, _ := strconv.ParseFloat(strings.Split(secondPart[i], "^")[1], 64)
			if v == 0 && len(strings.Split(firstPart[i], "x")[0]) == 1 {
				v = 1
			}
			degreeError = v
		} else if strings.Contains(strings.ToLower(secondPart[i]), "x") {
			v, _ := strconv.ParseFloat(strings.Split(secondPart[i], "x")[0], 64)
			if v == 0 && len(strings.Split(firstPart[i], "x")[0]) == 1 {
				v = 1
			}
			b2 += v
		} else {
			v, _ := strconv.ParseFloat(secondPart[i], 64)
			c2 += v
		}
	}

	if degreeError > 0 {
		fmt.Printf("%sError:%s Polynomial degree: %g\n", string(colorR), string(colorReset), degreeError)
		fmt.Println("The polynomial degree is stricly greater than 2, I can't solve.")
		os.Exit(2)
	}
	//
	// print readable
	//
	fmt.Printf("%sHuman readable:%s\n", string(colorP), string(colorReset))
	//print first part
	if a != 0 {
		if a == 1 {
			fmt.Printf("+x^2")
		} else if a == -1 {
			fmt.Printf("-x^2")
		} else {
			fmt.Printf("%+gx^2", a)
		}
	}
	if b != 0 {
		if b == 1 {
			fmt.Printf(" +x")
		} else if b == -1 {
			fmt.Printf(" -x")
		} else {
			fmt.Printf(" %+gx", b)
		}
	}
	if c != 0 {
		fmt.Printf(" %+g", c)
	}
	fmt.Printf(" = ")
	//print second part
	if c2 == 0 && b2 == 0 && a2 == 0 {
		fmt.Printf("0")
	} else {
		if a2 != 0 {
			if a2 == 1 {
				fmt.Printf("+x^2")
			} else if a2 == -1 {
				fmt.Printf("-x^2")
			} else {
				fmt.Printf("%+gx^2", a2)
			}
		}
		if b2 != 0 {
			if b2 == 1 {
				fmt.Printf(" +x")
			} else if b2 == -1 {
				fmt.Printf(" -x")
			} else {
				fmt.Printf(" %+gx", b2)
			}
		}
		if c2 != 0 {
			fmt.Printf(" %+g", c2)
		}
	}

	//get reduced form
	a = a + -a2
	b = b + -b2
	c = c + -c2

	//print reduced form
	fmt.Printf("\n\n%sReduced form:%s\n", string(colorT), string(colorReset))
	if c == 0 && b == 0 && a == 0 {
		fmt.Println("0 = 0")
	} else {
		if a != 0 {
			if a == 1 {
				fmt.Printf("+x^2")
			} else if a == -1 {
				fmt.Printf("-x^2")
			} else {
				fmt.Printf("%+gx^2", a)
			}
		}
		if b != 0 {
			if b == 1 {
				fmt.Printf(" +x")
			} else if b == -1 {
				fmt.Printf(" -x")
			} else {
				fmt.Printf(" %+gx", b)
			}
		}
		if c != 0 {
			fmt.Printf(" %+g", c)
		}
		fmt.Println(" = 0")
	}
	return a, b, c
}
