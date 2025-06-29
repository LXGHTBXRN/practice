package main

import "fmt"

type Engine struct {
	Power int
	Type  string
}

type Car struct {
	Brand  string
	Model  string
	Year   int
	Engine Engine
}

func main() {
	car := Car{
		Brand: "BMW",
		Model: "M3",
		Year:  2019,
		Engine: Engine{
			Power: 350,
			Type:  "Бензин",
		},
	}
	fmt.Printf("Марка: %s\nМодель: %s\nГод выпуска: %d\n", car.Brand, car.Model, car.Year)
	fmt.Printf("Двигатель: %d л.с., тип: %s\n", car.Engine.Power, car.Engine.Type)
}
