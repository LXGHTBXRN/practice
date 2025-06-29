package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthYear int
	AvgGrade  float64
}

func (s Student) Age() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

func (s Student) Status() string {
	switch {
	case s.AvgGrade >= 4.5:
		return "Отличник"
	case s.AvgGrade >= 3.5:
		return "Хорошист"
	default:
		return "Троечник"
	}
}

func main() {
	student := Student{
		Name:      "Андрей",
		BirthYear: 2004,
		AvgGrade:  4.4,
	}

	fmt.Printf("Студент: %s\n", student.Name)
	fmt.Printf("Возраст: %d\n", student.Age())
	fmt.Printf("Статус: %s\n", student.Status())
}
