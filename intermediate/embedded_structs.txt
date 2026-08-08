package main

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	person
	employeeID string
	experience  int
	title	string
	salary	 float64
}

func main() {
	fmt.Println("Experimenting with embedded structs in Go")

	seniorSoftwareEngineer := Employee{
		person: person{name: "John Doe", age: 30},
		employeeID: "E12345",
		salary: 120000.00,
		experience: 8,
		title: "Senior Software Engineer",
	}

	seniorSoftwareEngineer.introduce()
}

func (p person) introduce() {
	fmt.Printf("Hi, my name is %s and I am %d years old.\n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("Hi, my name is %s, I am %d years old, and I work as a %s with employee ID %s.\n", e.name, e.age, e.title, e.employeeID)
}