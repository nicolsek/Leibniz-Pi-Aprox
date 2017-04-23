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
	val := float64(3)
	odd := false

	for i := 2; i < n+2; i += 2 {
		switch odd {
		case false:
			val += float64(float64(4) / float64((i)*(i+1)*(i+2)))
		case true:
			val -= float64(float64(4) / float64((i)*(i+1)*(i+2)))
		}
		odd = !odd
	}

	return val
}

func getArgs() int {
	args := os.Args[1:]
	val, _ := strconv.Atoi(args[0])
	return val
}
