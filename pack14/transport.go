package main

import "fmt"

type Transport interface {
	Move()
	Stop()
}

type Car struct {
	Brand string
}

func (c Car) Move() {
	fmt.Printf("%s едет\n", c.Brand)
}

func (c Car) Stop() {
	fmt.Printf("%s остановился\n", c.Brand)
}

type Bicycle struct {
	Brand string
}

func (b Bicycle) Move() {
	fmt.Printf("Он едет на велосипеде %s\n", b.Brand)
}

func (b Bicycle) Stop() {
	fmt.Printf("%s остановился\n", b.Brand)
}

func main() {
	var t Transport

	t = Car{Brand: "Toyota"}
	t.Move()
	t.Stop()

	t = Bicycle{Brand: "Schwinn"}
	t.Move()
	t.Stop()
}
