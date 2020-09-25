package main

import "fmt"

func main() {
	var PI float32 = 3.1416
	var area float32
	var radius float32

	fmt.Println("\t***Circulo***\n")
	fmt.Println("Ingrese el radio del circulo")
	fmt.Scanf("%g", &radius)
	area = (radius * radius) * PI
	fmt.Println("Area del circulo: ", area)
}
