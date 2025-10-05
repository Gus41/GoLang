package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) sayHello() {
	fmt.Println("Olá! Meu nome é", p.name, "e tenho", p.age, "anos.")
}

func main() {
	var person Person = Person{name: "Gustavo", age: 22}
	person2 := Person{name: "Teste", age: 21}
	person.sayHello()
	person2.sayHello()

	var a, b int = 1, 2
	var c int = a + b
	fmt.Print(a)
	fmt.Print(c)

	f := "stringValue" // string type
	fmt.Print(f)

	//const
	const constantValue = "constant"

}
