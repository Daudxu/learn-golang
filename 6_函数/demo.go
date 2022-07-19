package main

import "fmt"

func demo1(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("err" + op)
	}

}

func demo2()(a int, b string) {
	c := 1
	d := "2"

    return c,d
}
func main() {
	sa := demo1(1, 2, "+")
	a, b := demo2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(sa)
}