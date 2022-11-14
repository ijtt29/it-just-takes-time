package main

import "fmt"

type Animal interface {
	Cry()
}

type Dog struct{}

func (d *Dog) Cry() {
	fmt.Println("멍멍")
}

type Cat struct{}

func (c *Cat) Cry() {
	fmt.Println("야옹")
}

func main() {
	var animal Animal
	dog := &Dog{}
	cat := &Cat{}

	animal = dog
	animal.Cry()

	animal = cat
	animal.Cry()

}
