package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning Closures in Go!")

	ticketPurchaseRequest := 30
	availableTickets := 75
	ticketsRefundRequest := 3

	fmt.Println("Available Tickets: ", availableTickets)
	soldTickets := ticketsSold(availableTickets)

	fmt.Println(ticketPurchaseRequest , " were requested to be purchased")

	for i := 0; i < ticketPurchaseRequest; i++ {
		availableTickets = soldTickets()
	}

	fmt.Println("Available Tickets: ", availableTickets)
	
	fmt.Println(ticketsRefundRequest , " were requested to be refunded")

	refundTicket := ticketRefund(availableTickets)

	for i := 0; i < ticketsRefundRequest; i++ {
		availableTickets = refundTicket()
	}

	fmt.Println("Available Tickets: ", availableTickets)


	fmt.Println("Summary!")
	fmt.Println("What i learned!")
	fmt.Println("I learned closures in Go. I realised how effectively we can create reusable functions that can maintain state across multiple invocations. This is particularly useful for scenarios like ticket sales and refunds, where we need to keep track of the number of available tickets.")


}

func ticketsSold(tickets int) func() int {
	return func() int {
		tickets--
		return tickets
	}
}

func ticketRefund(tickets int) func() int {
	return func() int {
		tickets++
		return tickets
	}
}