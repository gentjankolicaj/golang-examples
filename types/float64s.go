//go:build ignored
package main

import "fmt"

//Exported members start with uppercase,and are accessible outside their package
//Non-exported members start with lowercase and are not accesible outside their package


//float32     the set of all IEEE-754 32-bit floating-point numbers
//float64     the set of all IEEE-754 64-bit floating-point numbers


//Exported vars
var A float64
var B float64=12.333

//Un-exported vars
var c float64
var d float64=12.33

var(
	f float64=22.3

)
func main(){
	fmt.Println(A,B,c,d,f)

	var a,b float64=12,33
	fmt.Println(a,b)

}