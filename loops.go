package main

import (
	"fmt"
	"time"
)

func main() {

	nestedLoop()
	var current_time int = time.Now().Second()

	if( current_time % 2 == 0) {
		printNumbers()
	} else {
		printReportCard()
	}

}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}


func printReportCard() {
		marks_obtained := [] int {85, 81, 69, 78, 91, 43}
	subjects := [] string {"Maths", "Science", "English", "History", "Geography", "Arts"}

	var total int = 0

	below_average := 0

	for index, value := range marks_obtained {
		if (value < 80) {
			below_average += 1
			continue
		}

		fmt.Println(subjects[index] + " : " + fmt.Sprint(value))
		total += value
	}

	fmt.Println("---------------------------------")
	fmt.Println("Total Marks Obtained : " + fmt.Sprint(total))

	total_subjects := len(subjects)

	var average_marks float64 = float64(total) / float64(total_subjects)
	fmt.Println("Average Marks Obtained : " + fmt.Sprint(average_marks))
	fmt.Println("Number of Subjects Below Average : " + fmt.Sprint(below_average))
}

func nestedLoop() {
	rows := 5

	for i:= 1; i<= rows; i++ {
		for j:= 1; j<= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}