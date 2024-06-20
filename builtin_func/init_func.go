//go:build ignored
package main

import(
	"fmt"
)


func init(){
	fmt.Println("Before main-func is invoked, init-func is invoked.")
}


func main(){
	fmt.Println("Main-func invoked")
}