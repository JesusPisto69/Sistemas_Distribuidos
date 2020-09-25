package main

import "fmt"

func main() {
	var celcius float32
	var fahrenheit float32

	fmt.Println("\t***Temperatura***\n")
	fmt.Println("Ingrese los grados Fahrenheit")
	fmt.Scan(&fahrenheit)
	celcius = (fahrenheit - 32) * 5 / 9
	fmt.Println(fahrenheit, "grados Fahrenheit son", celcius, "grafos Celcius")
}
