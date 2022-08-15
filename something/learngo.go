package something

import (
	"fmt"
	"strings"
)

func multi(a int, b int) int {
	return a * b
}

func lendAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(word ...string) {
	fmt.Println(word)
}

func array() {

	names := []string{"Seo", "Kim", "Dal"}
	names = append(names, "Ukang")
	fmt.Print(names)

}

func maps() {
	nico := map[string]string{
		"name": "nico", "age": "12"}
	fmt.Println(nico)

	for key, value := range nico {
		fmt.Println(key, value)
	}
}

type person struct {
	name    string
	age     int
	favFood []string
}

func structs() {
	favFood := []string{"Kimchi", "Ramen"}
	nico := person{
		"nico",
		19,
		favFood,
	}
	seo := person{name: "Seo", age: 18, favFood: favFood}
	fmt.Println(nico, seo)

}

func main() {
	structs()

	maps()

	array()
	a := 2
	b := &a
	a = 10
	a = 77777777

	fmt.Println(a, *b)
	*b = 100
	fmt.Println(a, *b)

}
