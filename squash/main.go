package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Android struct {
	Person
	Model string
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Main 1")
	fmt.Println("Develop 2")
}

func (p *Person) AddPerson() {
	fmt.Println("persona agregada")
	fmt.Println("Persona agregada", p.Name)
}

func (a *Android) AddAndroid() {
	fmt.Println("Android")
	fmt.Println("Android", a.Model)
}
