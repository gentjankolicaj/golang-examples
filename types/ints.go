//go:build ignored

package main

import "fmt"

//Exported members start with uppercase,and are accessible outside their package
//Non-exported members start with lowercase and are not accesible outside their package

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