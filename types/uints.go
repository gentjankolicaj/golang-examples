//go:build ignored
package main

import "fmt"

//Exported members start with uppercase,and are accessible outside their package
//Non-exported members start with lowercase and are not accesible outside their package

//uint8       the set of all unsigned  8-bit integers (0 to 255)
//uint16      the set of all unsigned 16-bit integers (0 to 65535)
//uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
//uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

//Exported vars
var A uint
var B uint=10

//Un-exported vars
var c uint
var d uint

var(
	f uint=11
	g uint=100 //if negative values => overflow because uint accepts only positive values
)


func main(){
	fmt.Println(A,B,c,d,f,g)

	var a uint64=222
	var b uint64=222
	fmt.Println(a,b)
}