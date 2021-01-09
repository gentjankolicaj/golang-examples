package main

import "fmt"

var a1, b1 = 10, float64(10)

//Pointers to a1 & b1
var a1Ptr, b1Ptr = &a1, &b1

var (
	c = true
	d = "Hello world"
)

//Pointers to c & d
var (
	cPtr = &c
	dPtr = &d
)

func main() {

	fmt.Println("Values : ", a1, b1, c, d)
	fmt.Println("Ptrs : ", a1Ptr, b1Ptr, cPtr, dPtr, "\n")

	//Local declaration
	var e, f = 10, float64(13)
	var ePtr, fPtr = &e, &f

	var (
		g = true
		h = "Hello local variable"
	)

	var (
		gPtr = &g
		hPtr = &h
	)

	fmt.Println("Values : ", e, f, g, h)
	fmt.Println("Ptrs : ", ePtr, fPtr, gPtr, hPtr, "\n")

	//Declaration (*T where T is type ) and assignment (&var of type T)
	var myPtr *int
	myPtr = &e //because e is of type int
	fmt.Println("myPtr address & value : ", myPtr, *myPtr)

}
