package main

import "fmt"

type Student struct {
	Name     string
	Age      int
	Course   int
	AvgGrade float64
}

func NewStudent(name string, age int, course int, avgGrade float64) Student {
	return Student{
		Name:     name,
		Age:      age,
		Course:   course,
		AvgGrade: avgGrade,
	}
}

func (s Student) PrintInfo() {
	fmt.Printf("Имя: %s\nВозраст: %d\nКурс: %d\nСредний балл: %.2f\n", s.Name, s.Age, s.Course, s.AvgGrade)
}

func (s *Student) UpdateAvgGrade(newGrade float64) {
	s.AvgGrade = newGrade
}

func main() {
	student := NewStudent("Александр", 20, 2, 4.0)

	student.PrintInfo()

	student.UpdateAvgGrade(4.5)
	fmt.Println("После обновления среднего балла")
	student.PrintInfo()
}
