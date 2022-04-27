package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Profession struct {
	ProfessionNumber string
	ProfessionTitle  string
	ProfessionDesc   string
	InDebtTo         string
}

func main() {
	rand.Seed(time.Now().UnixNano())

	num1 := d3()
	num2 := d4()
	fmt.Println(num1)
	fmt.Println(num2)
}
