//go:build ignored
package main

import "fmt"

//Exported members start with uppercase,and are accessible outside their package
//Non-exported members start with lowercase and are not accesible outside their package


//int8        the set of all signed  8-bit integers (-128 to 127)
//int16       the set of all signed 16-bit integers (-32768 to 32767)
//int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
//int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

//Exported vars
var A int
var B int=10

//Non-Exported vars
var c int
var d int=1

var (
	f int=1
	g=11
)

func main(){
    var a int
	a=20

	fmt.Println("a ",a)

	fmt.Println(A,B,c,d,f,g)
	
}