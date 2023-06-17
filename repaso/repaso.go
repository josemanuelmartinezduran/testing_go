package main

/* Hola mundo
func hola(){
	fmt.Println("hello world")
}
func main() {
	go hola()
} */

/*
func hastaDiez(numero int) int {
	if numero > 10 {
		return 0
	} else {
		fmt.Println("Numero ", numero)
		return hastaDiez(numero + 1)
	}
}*/

/*  Valores
func main() {

	fmt.Println("Hola" + "Mundo")
	fmt.Println("1+1 =", 1+1)

	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
*/
/* Variables */

/* func main() {

	f, _ := "Mi dato", 3
	fmt.Println(f)

	a, b := true, false
	c := !(a != b)
	fmt.Println("Resultado", c)
} */

/* If

func main() {
	valor := 8
	if valor%2 == 0 {
		fmt.Printf("El número %d es par", valor)
	} else {
		fmt.Printf("El numero %d es impar", valor)
	}
} */

/* For */
/*
func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	ké := 1
	for {
		fmt.Println(ké)
		if ké >= 3 {
			break
		}
		ké++
	}

	for n := 0; n <= 5; n++ {
		if n%2 != 0 {
			continue
		}
		fmt.Println(n)
	}
}*/

/* Slices */
/* func main() {

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s))

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	s[0] = "0"
	s[1] = "1"
	s[2] = "2"
	fmt.Println("set:", s)
	fmt.Println("get:", s[len(s)-1])
	fmt.Println("len:", len(s))

	s = append(s, "3")
	s = append(s, "4", "5", "6", "7")
	if "a" > "b" {
		fmt.Println("A es mayor que B")
	} else {
		fmt.Println("B es mayor que A")
	}
	fmt.Println("apd:", s)
	c := make([]string, len(s))
	nc := s
	copy(c, s)
	c[1] = "Hola"
	fmt.Println("El slice s tiene", s)
	fmt.Println("El slice copiado tiene", c)
	fmt.Println("El slice no copiado tiene", nc)
	fmt.Println("cpy:", c)
	l := s[2:5]
	fmt.Println("sl 2:5:", l)
	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)
	l = s[:len(s)-1]
	fmt.Println("Con len -2", l)
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
} */

/* Mapas */

/* func main() {

	m := make(map[string]int)

	m["k1"] = 10
	m["k2"] = 50
	m["k3"] = 7

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	i, v3 := m["k3"]
	if v3 {
		fmt.Println("Valor de k3")
		fmt.Println(v3)
		fmt.Println(i)
	}

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	zero, prs := m["k2"]
	fmt.Println("prs:", prs, zero)

	n := map[string]int{
		"foo":           1,
		"otro elemento": 4,
		"bar":           2}

	fmt.Println("map:", n)

} */

/* Range */

/* func main() {

	nums := []int{2, 3, 4, 5, 6, 7, 8, 10}
	sum := 0
	for index, num := range nums {
		if index > 7 {
			break
		}
		fmt.Printf("El indice %d tiene el valor %d \n\r", index, num)
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println(k, "->", v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
} */

/* Funciones */
/*
func suma(a, b, c int) int {
	return a + b + c
}

func sumaExtraña(a, b, c int) (bool, int) {
	return a > b, a + b + c
}

func sumaTodo(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}*/

/* func main() {
	a := suma(1, 2, 3)
	fmt.Println("La suma es", a)

	b, c := sumaExtraña(2, 3, 4)
	fmt.Println("El resultado es", b, c)

	fmt.Println("El total es:", sumaTodo(2, 3, 4, 5, 6, 7, 8, 45, 955))

} */

/* Apuntadores */
/*
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	fmt.Println("la direccion de memoria de i es:", iptr)
	*iptr = 0
}*/

/*
func main() {
	i := 1
	fmt.Println("El valor de i es", i)
	fmt.Println("i esta almacenada en la posicion", &i)
	j := &i
	fmt.Println("El valor de j es", *j)
	fmt.Println("J esta almacenada en la posicion", j)
	*j = 50
	fmt.Println("El valor de i es", i)
	fmt.Println("i esta almacenada en la posicion", &i)
	fmt.Println("valor inicial:", i)
	zeroval(i)
	fmt.Println("valor no modificado", i)
	zeroptr(&i)
	fmt.Print("valor modificado", i)
	fmt.Println("Que direccion tiene i", &i)
}*/

/* Structs */

/* type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

type perro struct {
	numeroPatas int
	nombre      string
}

func newPerro(nombre string) *perro {
	p := perro{4, nombre}
	return &p
}

func main() {
	solobino := newPerro("Solobino")
	fmt.Println(solobino)

	p := person{"Jose", 40}
	fmt.Println("La persona ", p.name, "tiene", p.age)

	j := person{name: "Lupita"}
	fmt.Println(j)

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
} */

/* Metodos */
/*
type rect struct {
	width, height int
}

func (q *rect) area() int {
	return q.width * q.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func calculaArea(r rect) int {
	return r.width * r.height
}

func main() {
	recta := rect{3, 5}
	fmt.Println("El area es ", recta.area(), "el perimetro es ", recta.perim())
	otro := rect{800, 1600}
	fmt.Println("El area es ", otro.area())
}*/
/*
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("El area es", g.area(), "y el perimetro es", g.perim())
}

func main() {

	sg := []geometry{rect{3, 4}, circle{5}}
	for _, val := range sg {
		measure(val)
	}
}
*/
/* Generics */

/* func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lista *List[T]) agrega(val T) {
	if lista.head == nil {
		lista.head = &element[T]{nil, val}
		lista.tail = lista.head
	} else {
		lista.head = &element[T]{lista.tail, val}
		lista.tail = lista.head
	}
}

func (lista *List[T]) imprime() {
	elemento := lista.head
	for {
		fmt.Println(elemento.val)
		if elemento.next == nil {
			break
		}
		elemento = elemento.next
	}
}

type miTipo[T any] struct {
	a int
	b float64
	c string
	d T
}

func main() {
	lista := List[int]{}
	lista.agrega(0)
	lista.agrega(1)
	lista.agrega(2)
	lista.agrega(3)
	lista.agrega(4)
	lista.agrega(5)
	lista.agrega(6)
	lista.agrega(7)
	lista.imprime()

	r := miTipo[string]{}
	r.a = 1
	r.b = 3.4
	r.c = "Dato"
	r.d = "ejemplo"
	fmt.Println(r)

	q := miTipo[int]{1, 3.4, "Dato", 3}
	fmt.Println(q)
} */

/* Mapas */
/* func main() {
	calificaciones := map[string]float64{}
	calificaciones["Jose"] = 8.2
	calificaciones["Javier"] = 9.1
	calificaciones["Sophia"] = 10.0
	fmt.Println(calificaciones["Jose"])
	fmt.Println(calificaciones)
}
*/
/* Recursivo */
/* func recurPrime(i int, isPrime bool, div int) (int, bool, int) {
	if i <= 1 {
		return i, false, div
	} else if i == 2 {
		return i, true, div
	} else if div < i/2 {
		return i, true, div
	} else {
		return recurPrime(i, isPrime, (div - 1))
	}
} */

/*
func impares(final int) {
	for i := 1; i <= final; i += 2 {
		fmt.Println("Numero impar ", i)
	}
}

func main() {
	impares(100)
} */

/* var m = map[int]string{1: "2", 2: "4", 4: "8"}

fmt.Println("keys:", MapKeys(m))

_ = MapKeys[int, string](m)

lst := List[int]{}
lst.Push(10)
lst.Push(13)
lst.Push(23)
fmt.Println("list:", lst.GetAll()) */

/* Goroutines */
/* func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Second)
	}
}

func main() {

	f("hola")

	go f("mundo")
	go f("mundo2")
	go f("mundo3")
	go f("mundo4")
	go f("mundo5")
	go f("mundo6")
	go f("mundo7")
	go f("mundo8")
	go f("mundo9")
	go f("mundo10")

	go func(msg string) {
		fmt.Println(msg)
	}("que rayos??")

	time.Sleep(time.Second * 2)
	fmt.Println("fin")
} */
/*
func x(messages chan string, valor string) {
	messages <- valor
}

func y(message chan string, mensajex string) {
	fmt.Println("Impriminedo desde y", mensajex)
	message <- "pong"
}

func main() {
	messages := make(chan string)
	go x(messages, "Ping")
	mensajeString := <-messages
	fmt.Println(mensajeString)
	mensajey := make(chan string)
	go y(mensajey, mensajeString)
	mensajeString = <-mensajey
	fmt.Println(mensajeString)
	go x(messages, "Pong")
	mensajeString = <-messages
	fmt.Println(mensajeString)
}
*/
