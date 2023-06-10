package main

import "fmt"

func main() {
	fmt.Printf("Hola mundo")
	_, msg := esPrimo(7)
	fmt.Printf(msg)
}

func esPrimo(n int) (bool, string) {
	if n < 2 {
		return false, fmt.Sprintf("El numero %d no es primo por definición", n)
	}
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("El número %d no es primo es divisible entre %d", n, i)
		}
	}
	return true, fmt.Sprintf("El numero %d es primo", n)
}
