package main

import (
	"fmt"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
		// panic("unsupported operation: " + op)
	}
}

// func div(a, b int) (q, r int) {
// 	q = a / b
// 	r = a % b
// }

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 4)

	fmt.Println(q, r)
}
