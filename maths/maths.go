package maths

import (
	"fmt"
)

type Equation struct {
	Sum    string
	Result float64
}

func Add(x float64, y float64, channel chan Equation) {
	equ := Equation{
		Sum:    fmt.Sprintf("%.2f + %.2f", x, y),
		Result: x + y,
	}
	channel <- equ
}
