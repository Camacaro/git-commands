package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Directions struct {
	North, South, East, West bool
}

type Animal struct {
	Name string
}

type Cat struct {
	Animal
}
type Dog struct {
	Animal
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Esto es un comentario en hotfix")
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func (d Directions) String() string {
	return fmt.Sprintf("North: %v, South: %v, East: %v, West: %v", d.North, d.South, d.East, d.West)
}
