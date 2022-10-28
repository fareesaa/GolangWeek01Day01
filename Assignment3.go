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

	fmt.Printf("a = %d is type of %T\n", a, a)
	fmt.Printf("b = %d is type of %T\n", b, b)
	fmt.Printf("c = %d is type of %T\n", c, c)
	fmt.Printf("d = %d is type of %T\n", d, d)
	fmt.Printf("e = %f is type of %T\n", e, e)
	fmt.Printf("f = %f is type of %T\n", f, f)
	fmt.Printf("g = %s is type of %T\n", g, g)
}
