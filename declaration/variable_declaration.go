package main

import "fmt"

var a, b = 9, 10

var (
	A, B = 11, 12
)

func main() {
	fmt.Println(a, b, A, B)

	var c, d = 13, 14
	var e, f = 17, 18

	var (
		C = 15
		D = 16

		E = 19
		F = 20
	)

	fmt.Println(c, d, C, D)
	fmt.Println(e, f, E, F)
}
