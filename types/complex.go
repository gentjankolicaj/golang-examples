//go:build ignored
package main

import "fmt"

//Exported members start with uppercase & are visible outside package
//Unexported members start with lowercase & are not visible outside package

//complex64 & complex128

//Exported vars
var A complex64
var B complex128

//Unexported vars
var c complex64=complex(2,3)
var d complex128=complex(4,4)


var (
	f complex64=2+3i
	
)

func main(){
	fmt.Println(A,B,c,d,f)

	fmt.Println("d & f real part : ",real(d),real(f))
	fmt.Println("d & f img part : ",imag(d),imag(f))

}
