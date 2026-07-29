package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Welcome to the Go Loops Program")
	var current_time int = time.Now().Second()

	if( current_time % 2 == 0) {
		printNumbers()
	}else if (current_time % 3 == 0) {
		nestedLoop()
	}else if (current_time % 5 == 0) {
		guessTheNumberGame()
	}else if (current_time % 7 == 0) {
		fruitCheck("Mango")
	}else {
		printReportCard()
	}

}

func fruitCheck(fruitName string) {
	switch fruitName {
		case "Apple":
			fmt.Println("You have selected Apple")
		case "Banana":
			fmt.Println("You have selected Banana")
		case "Mango":
			fmt.Println("You have selected Mango")
		default:
			fmt.Println("You have selected an unknown fruit")
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

func guessTheNumberGame () {
fmt.Println("Lets play the guessing number game")

	randomNumber := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randomNumber)

	targetNumber := random.Intn(100) + 1

	var guess int

	for {
		fmt.Print("Enter your guess (1-100): ")
		fmt.Scanln(&guess)

		if guess < targetNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > targetNumber {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("Congratulations! You guessed the correct number:", targetNumber)
			break
		}
	}
}