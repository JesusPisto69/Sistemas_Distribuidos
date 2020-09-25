package main

import "fmt"

func main() {
	var area float32
	var side float32

	fmt.Println("\n\t***Cuadrado***\n")
	fmt.Println("Ingrese el valor del lado")
	fmt.Scan(&side)
	area = side * side
	fmt.Println("El area del cuadrado es: ", area)
}
