package main

import "fmt"

func main2() {

	var a uint8 = 5
	var b uint16 = 6

	var c int8 = -5
	var d int16 = -6

	var e float32 = -5.5
	var f float64 = -6.6

	var g = "foo"

	fmt.Printf("a is type %T\n", a)
	fmt.Printf("b is type %T\n", b)
	fmt.Printf("c is type %T\n", c)
	fmt.Printf("d is type %T\n", d)
	fmt.Printf("e is type %T\n", e)
	fmt.Printf("f is type %T\n", f)
	fmt.Printf("g is type %T\n", g)
}
