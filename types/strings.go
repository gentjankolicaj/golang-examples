//go:build ignored
package main

import(
	"fmt"
)


//Exported members start with uppercase & are visible outside package
//Unexported members start with lowercase & are not visible outside package

//Exported vars
var A string
var B string="Hello world"

//Unexported vars
var c string
var d string="Exported var d"


var (
	byteSlice=[]byte{'G','o','l','a','n','g'}
	f string=string(byteSlice)

)

func main(){
	fmt.Println(A,B,c,d)

	fmt.Println("byteSlice : ",byteSlice)
	fmt.Println(f)
}