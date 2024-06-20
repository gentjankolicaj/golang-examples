//go:build ignored
package main

import(
	"fmt"
)

//Exported members start with uppercase & are visible outside package
//Unexported members start with lowercase & are not visible outside package but inside

//byte        alias for uint8
//where uint8 => the set of all unsigned  8-bit integers (0 to 255)

//Exported vars
var A byte
var B byte=127

//Unexported vars
var c byte=11
var d byte

var (
	f byte=0
	g byte=125
)


func main(){
	fmt.Println(A,B,c,d,f,g)

	var a,b byte=0,0
	fmt.Println(a,b)
}
