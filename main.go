package main

import (
	"fmt"
	"strings"
)

func main(){

	conferenceName :="Go Conference"
	var remainingTickets uint= 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	for{
		var firstName , lastName, email string
		var userTickets uint
		fmt.Printf("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Printf("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Printf("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Printf("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets || userTickets ==0 {
			fmt.Printf("Invalid number of tickets, remaining tickets are %v\n", remainingTickets)
		}

		remainingTickets -= userTickets
		fullName := firstName + " " + lastName
		bookings = append(bookings, fullName)

		fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", fullName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings{
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Printf("Our conference is booked out. Come back next year.\n")
			break
		}


	}

}