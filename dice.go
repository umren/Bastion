package main

import (
	"math/rand"
)

func d3() int {
	min := 1
	max := 3

	return rand.Intn(max-min+1) + min
}

func d4() int {
	min := 1
	max := 4

	return rand.Intn(max-min+1) + min
}

func d6() int {
	min := 1
	max := 6

	return rand.Intn(max-min+1) + min
}

func d8() int {
	min := 1
	max := 8

	return rand.Intn(max-min+1) + min
}

func d10() int {
	min := 1
	max := 10

	return rand.Intn(max-min+1) + min
}

func d12() int {
	min := 1
	max := 12

	return rand.Intn(max-min+1) + min
}

func d20() int {
	min := 1
	max := 20

	return rand.Intn(max-min+1) + min
}

func d100() int {
	min := 1
	max := 100

	return rand.Intn(max-min+1) + min
}

func dANY(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
