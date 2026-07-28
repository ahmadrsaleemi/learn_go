package main

import "fmt"

var firstName string = "Ahmad"
var lastName string = "Raza"

type employee struct {
	firstName string
	lastName  string
	age       int
	grossSalary int
}

func main() {
	var courseName string = "Learn Go"
	fmt.Println("Course Name: ", courseName)
	// var age int = 31
	// count := 10
	lastName = "Saleemi"
	var fullName = concatenateString(firstName, lastName)
	printName(fullName)

	var days int = 2
	var hours_daily int = 2

	totalHours := days * hours_daily

	fmt.Println("Total Hours spent on the course: ", totalHours)

	const SALARYPERHOUR = 1000

	emp := employee{}
	emp.firstName = firstName
	emp.lastName = lastName
	emp.age = 31
	emp.grossSalary = totalHours * SALARYPERHOUR

	printEmployeeDetails(emp)



}

func concatenateString(firstName string, lastName string) string {
	var fullName string = firstName + " " + lastName
	return fullName
}

func printName(name string) {
	fmt.Println("My name is: ", name)
}

func printEmployeeDetails(employee employee) {
	details := "My name is " + concatenateString(employee.firstName, employee.lastName) + " and my age is " + fmt.Sprint(employee.age) + " and I have earned " + fmt.Sprint(employee.grossSalary)
	fmt.Println(details)
}