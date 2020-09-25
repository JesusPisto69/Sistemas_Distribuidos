package main

import "fmt"

func main() {
	var area float32
	var base float32
	var height float32

	fmt.Println("\n\t***Triangulo***\n")
	fmt.Println("Ingrese el valor de la base")
	fmt.Scan(&base)
	fmt.Println("Ingrese la altura")
	fmt.Scan(&height)
	area = (base * height) / 2
	fmt.Println("El area del triangulo es: ", area)
}
