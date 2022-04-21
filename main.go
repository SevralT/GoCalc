package main

import (
	"fmt"
	"os"
	"strconv"
	oper "github.com/SevralT/GoCalc/oper"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Использование: a (+, -, *, /, %) b")
		os.Exit(1)
	}

	a := os.Args[1]
	op := os.Args[2]
	b := os.Args[3]
	var c float64

	a_new, err := strconv.ParseFloat(a, 8)
	b_new, err := strconv.ParseFloat(b, 8)

	if err != nil {
		fmt.Println(err)
	}

	switch op {
	case "+":
		c = oper.Add(a_new, b_new)
	case "-":
		c = oper.Minus(a_new, b_new)
	case "*":
		c = oper.Multi(a_new, b_new)
	case "/":
		c = oper.Div(a_new, b_new)
	}
	fmt.Println("Ваш результат:", c)
}
