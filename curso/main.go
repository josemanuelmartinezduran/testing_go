package main

import "fmt"

func main() {
	fmt.Println("Hola mundo")
}

func esPrimo(n int) (bool, string) {
	if n == 0 {
		return false, "0 No es primo"
	}
	if n < 0 {
		return false, fmt.Sprintf("%d No es primo, es negativo", n)
	}
	for i=2 ; i<range()
}
