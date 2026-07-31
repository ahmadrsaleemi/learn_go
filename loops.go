package main

import (
	"fmt"
	"maps"
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
	}else if (current_time % 11 == 0) {
		returnBalonDorWinner()
	}else if(current_time % 13 == 0){
		mapIntroduction()
	}else {
		printReportCard()
	}

}

func mapIntroduction() {
	myMap := make(map[string]int)
	myMap["streetname"] = 123
	myMap["zipcode"] = 45678
	myMap["housenumber"] = 42
	fmt.Println("Map: ", myMap)

	delete(myMap, "zipcode")
	fmt.Println("Map: ", myMap)

	// clear(myMap)
	// fmt.Println("Map: ", myMap)

	_, unknown := myMap["streetname"]
	fmt.Println("Is a value associated with: " , unknown)

	orderCancelled := map[int]string {1001: "John Doe", 1002: "Christopher"}

	paymentFailed := map[int]string {1001: "John Doe", 1002: "Christopher"}

	if maps.Equal(orderCancelled, paymentFailed) {
		fmt.Println("All orders whose paymentFailed are cancelled!")
	}

	fmt.Println("List of cancelled orders!")

	for orderId, customerName := range orderCancelled {
		fmt.Println("Order Id: " , orderId , " and customer name: " , customerName)
	}
}

func sliceIntroduction() {
	// var numbers [] int
	// var numbers1 = [] int {1, 2, 3, 4, 5}

	// numbers2 := [] int {6, 7, 8, 9, 10}

	// slice := make([] int, 5)

	arr := [5] int {11, 12, 13, 14, 15}

	slice1 := arr[1:4]

	fmt.Println("Slice 1: ", slice1)

	sliceCopy := make([] int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("Slice Copy: ", sliceCopy)

	for index, value := range slice1 {
		fmt.Println("Index: ", index, " Value: ", value)
	}

	twoD := make([][] int, 5)
	for i:=0; i<5; i++ {
		innerLen := i + 1
		twoD[i] = make([] int, innerLen)
		for j:=0; j<innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	

	for _, innerSlice := range twoD {
		fmt.Println(innerSlice)
	}

	sliced := slice1[1:3]
	fmt.Println("Sliced: ", sliced)
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

func returnBalonDorWinner() {
	var list [][]string = [][]string{{"Lionel Messi", "2021"}, {"Cristiano Ronaldo", "2017"}, {"Luka Modric", "2018"}}

	fmt.Println("List of Ballon d'Or Winners:")
	for _, winner := range list {
		fmt.Println("Player:", winner[0], ", Year:", winner[1])
	}
}