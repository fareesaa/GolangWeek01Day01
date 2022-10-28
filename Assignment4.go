package main

import "fmt"

func assignment4() {
	var i = 5

	var j float64 = -6.6

	fmt.Printf("i is type of %T\n", i)
	fmt.Println("Value of i is: ", *(&i), " at address ", &i)

	fmt.Printf("j is type of %T\n", j)
	fmt.Println("Value of j is: ", *(&j), " at address ", &j)
}
