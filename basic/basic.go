package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d,%q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Println(
		cmplx.Exp(1i*math.Pi)+1,
		cmplx.Pow(math.E, 1i*math.Pi)+1)
	fmt.Printf("%.3f,%.3f\n",
		cmplx.Exp(1i*math.Pi)+1,
		cmplx.Pow(math.E, 1i*math.Pi)+1)
}

func main() {
	fmt.Println("Hello ChenYao")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)

	euler()
}
