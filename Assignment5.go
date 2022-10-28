package main

import "fmt"

func assignment5() {
	var a = 2
	var b = 3
	var c = 5

	fmt.Println("Pertambahan")
	fmt.Printf("%d + %d = %d", c, c, c+c)
	fmt.Println("")
	fmt.Println("")

	fmt.Println("Pengurangan")
	fmt.Printf("%d - %d = %d", c, c, c-c)
	fmt.Println("")
	fmt.Println("")

	fmt.Println("Perkalian")
	fmt.Printf("%d * %d = %d", c, c, c*c)
	fmt.Println("")
	fmt.Println("")

	fmt.Println("Pembagian")
	fmt.Printf("%d / %d = %d", c, c, c/c)
	fmt.Println("")
	fmt.Println("")

	fmt.Println("Perhitungan")
	fmt.Printf("%d + %d * %d = %d", a, b, c, a+(b*c))
}
