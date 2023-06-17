package main

import "fmt"

func main() {
	_, mensaje, _, _ := esPrimo(false, "", 4, 2)
	fmt.Println(mensaje)
}

func isPrime(numero int) (bool, string) {
	esNumeroPrimo, mensaje, _, _ := esPrimo(true, "", numero, 2)
	return esNumeroPrimo, mensaje
}

func esPrimo(esNumeroPrimo bool, mensaje string, numero, divisor int) (bool, string, int, int) {
	if numero < 2 {
		return false, fmt.Sprintf("El numero %d no es primo por definición", numero), numero, divisor
	} else if divisor >= ((numero / 2) + 1) {
		return true, fmt.Sprintf("El numero %d es primo", numero), numero, divisor
	} else if numero%divisor == 0 {
		return false, fmt.Sprintf("El numero %d no es primo es divisible entre %d", numero, divisor), numero, divisor
	} else {
		return esPrimo(esNumeroPrimo, mensaje, numero, divisor+1)
	}
}
