//go:build ignored
package main

import "fmt"

//Exported members start with uppercase & are visible outside package
//Unexported members start with lowercase & are not visible outside package but inside

//byte => true , false values

//Exported vars
var A bool
var B bool=true

//Un-exported vars
var c bool
var d bool=true

var(
	f=true
	g bool=false
)

func main(){
	fmt.Println(A,B,c,d,f,g)

	var a ,b , C bool= true,true,true
	fmt.Println(a,b,C)

}