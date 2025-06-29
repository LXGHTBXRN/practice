package main

import "fmt"

func addGrade(grades map[string]int, name string, grade int) {
	grades[name] = grade
}

func findGrade(grades map[string]int, name string) (int, bool) {
	grade, ok := grades[name]
	return grade, ok
}

func removeGrade(grades map[string]int, name string) {
	delete(grades, name)
}

func main() {
	grades := make(map[string]int)

	addGrade(grades, "Ivan", 4)
	addGrade(grades, "Vlad", 3)
	addGrade(grades, "Alex", 5)

	fmt.Println("Оценки студентов: ", grades)

	name := "Alex"
	if grade, found := findGrade(grades, name); found {
		fmt.Println("Оценка студента", name, grade)
	} else {
		fmt.Println("Студент не найден", name)
	}

	removeGrade(grades, "Ivan")
	fmt.Println(grades)
}
