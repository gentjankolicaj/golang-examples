package main

import "fmt"

//Empty struct => no fields
type A struct{}

//With values
type B struct {
	field0 bool
	field1 uint8
	field2 int
	field3 float64
	field4 complex64
	field5 string
}

//With pointers
type C struct {
	field0 *bool
	field1 *uint8
	field2 *int
	field3 *float64
	field4 *complex64
	field5 *string
}

//With fields of type struct
type D struct {
	field0 A
	field1 B
	field2 C
}

//With pointers to struct type
type E struct {
	field0 *A
	field1 *B
	field2 *C
}

//With different fields
type F struct {
	field0 int
	field1 *int
	field2 D
	field3 E
}

func main() {
	var a = A{}
	fmt.Println("a ", a)

	var b = B{field0: true, field1: 12}
	fmt.Println("b ", b)

	var val0, val1 = true, "Hello pointer"
	var c = C{field0: &val0, field5: &val1}
	fmt.Println("c ", c)

	var d = D{field0: A{}, field1: B{field0: true, field1: 2}, field2: C{}}
	fmt.Println("d ", d)

	var e = E{field0: &a, field1: &b, field2: &c}
	fmt.Println("e ", e)

	//Type can be ommited when we provide value on declaration
	var val2 = 100
	var f = F{field0: 10, field1: &val2, field2: d, field3: e}
	fmt.Println("f ", f)

}
