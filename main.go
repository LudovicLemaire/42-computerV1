package main

const (
	colorReset = "\033[0m"

	colorR = "\033[1;31m"
	colorG = "\033[1;32m"
	colorY = "\033[1;33m"
	colorB = "\033[1;34m"
	colorP = "\033[1;35m"
	colorT = "\033[1;36m"
)

func main() {
	a, b, c := parser()
	if a != 0.0 {
		degreeTwo(a, b, c)
	} else if b != 0 {
		degreeOne(b, c)
	} else {
		degreeZero(c)
	}
}
