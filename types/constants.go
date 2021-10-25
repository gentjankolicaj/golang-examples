//go:build ignored
package main

import(
	"fmt"
)

//Exported members start with uppercase & are visible outside package
//Unexported members start with lowercase & are not visible outside package

//Constants values must be assigned and can't be changed.

//Exported constants
const(
 A int=101
 B float64=12.3
 C bool=true
 D complex128=complex(3.14,3.14)
)


//Unexported constants
const(
	a int=102
	b float64=102
	c string="Hello constants"
	d =true
	
)

const X int=10

func main(){
	fmt.Println("Exported constants ",A,B,C,D)
	fmt.Println("Unexported constants ",a,b,c,d)
}