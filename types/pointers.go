//go:build ignored
package main


import(
	"fmt"
)

//Exported members start with uppercase & are visible outside package 
//Unexported members start with lowercase & are not visible outside package

/**

A pointer type denotes the set of all pointers to variables of a given type, called the base type of the pointer.
 The value of an uninitialized pointer is nil.


For an operand x of type T, the address operation &x generates a pointer of type *T to x.
The operand must be addressable, that is, either a variable, pointer indirection, 
or slice indexing operation; 
or a field selector of an addressable struct operand; 
or an array indexing operation of an addressable array.
As an exception to the addressability requirement, x may also be a (possibly parenthesized) composite literal.
If the evaluation of x would cause a run-time panic, then the evaluation of &x does too.

For an operand x of pointer type *T, the pointer indirection *x denotes the variable of type T pointed to by x.
If x is nil, an attempt to evaluate *x will cause a run-time panic.

*/

var iTest int=10
var fTest float64=3.14
var sTest string="Hello pointer"


//Exporte vars
var A *int
var B *int=&iTest
var C *float64
var D *float64=&fTest
var F *string
var G *string=&sTest


//Unexported vars
var a *int
var b *int=&iTest
var c *float64
var d *float64=&fTest
var f *string
var g *string=&sTest

func main(){
	fmt.Println("1- Pointer values : ",A,B,C,D,F,G)
	//fmt.Println("1- Actual values : ",*A,*B,*C,*D,*F,*G) Cause runtime panic because some pointers are nil => nil pointer dereference
	fmt.Println("1- Actual values : ",*B,*D,*G)

	fmt.Println("2- Pointer values : ",a,b,c,d,f,g)
	fmt.Println("2- Actual values : ",*b,*d,*g)

}

