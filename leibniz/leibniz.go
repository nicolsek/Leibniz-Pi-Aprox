package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := getArgs()
	PiApprox := approx(n)
	fmt.Printf("Pi estimated using %v itterations: %v\n", n, PiApprox)
}

func approx(n int) float64 {
	val := float64(1)
	odd := false
	for i := 3; i < n+3; i += 2 {
		odd = !odd

		switch odd {
		case false:
			val += float64(1) / float64(i)
		case true:
			val -= float64(1) / float64(i)
		}

	}

	return 4 * val
}

func getArgs() int {
	args := os.Args[1:]
	val, _ := strconv.Atoi(args[0])
	return val
}
