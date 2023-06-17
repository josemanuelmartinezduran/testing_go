package main

import (
	"testing"
)

func Test_esPrimo(t *testing.T) {
	/* esNumeroPrimo, _, _, _ := esPrimo(true, "", 0, 2)
	if esNumeroPrimo {
		t.Error("El resultado debería ser falso ya que 0 no es primo")
	}
	esNumeroPrimo, _, _, _ = esPrimo(true, "", -3, 2)
	if esNumeroPrimo {
		t.Error("El resultado debería ser falso ya que -3 no es primo")
	}
	esNumeroPrimo, _, _, _ = esPrimo(true, "", 1, 2)
	if esNumeroPrimo {
		t.Error("El resultado debería ser falso ya que 1 no es primo")
	}
	esNumeroPrimo, _, _, _ = esPrimo(true, "", 2, 2)
	if !esNumeroPrimo {
		t.Error("El resultado debería ser verdadero ya que el 2 es primo")
	}
	esNumeroPrimo, _, _, _ = esPrimo(true, "", 4, 2)
	if esNumeroPrimo {
		t.Error("El resultado debería ser falso ya que 4 no es primo")
	}
	esNumeroPrimo, _, _, _ = esPrimo(true, "", 7, 2)
	if !esNumeroPrimo {
		t.Error("El resultado debería ser verdadero ya que 7 si es primo")
	}

	esNumeroPrimo, _, _, _ = esPrimo(true, "", 11, 2)
	if !esNumeroPrimo {
		t.Error("El resultado debería ser verdadero ya que 11 si es primo")
	}

	esNumeroPrimo, _, _, _ = esPrimo(true, "", 97, 2)
	if !esNumeroPrimo {
		t.Error("El resultado debería ser verdadero ya que 97 si es primo")
	}
	*/

	primeTests := []struct {
		nombre        string
		numeroProbado int
		esNumeroPrimo bool
		mensaje       string
	}{
		{"Cero no es primo", 0, false, "El numero 0 no es primo por definición"},
		{"Uno no es primo", 1, false, "El numero 1 no es primo por definición"},
		{"Dos es primo", 2, true, "El numero 2 es primo"},
		{"Tres es primo", 3, false, "El numero 0 no es primo por definición"},
		{"Cuatro no es primo", 4, false, "El numero 0 no es primo por definición"},
		{"Cinco es primo", 5, false, "El numero 0 no es primo por definición"},
		{"Seis no es primo", 6, false, "El numero 0 no es primo por definición"},
		{"Siete no es primo", 7, false, "El numero 0 no es primo por definición"},
	}

	for _, prueba := range primeTests {
		res, mensaje := isPrime(prueba.numeroProbado)
		if res != prueba.esNumeroPrimo || prueba.mensaje != mensaje {
			t.Errorf("La prueba %s fallo esperaba %s y dio %s", prueba.nombre, prueba.mensaje, mensaje)
		}
	}
}
