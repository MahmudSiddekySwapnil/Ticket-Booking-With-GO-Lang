package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	for {
		var firstName, lastName, email string
		var userTickets uint

		fmt.Printf("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Printf("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Printf("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Printf("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		// ✅ Validation + logic
		if firstName == "" || lastName == "" {
			fmt.Println("Name cannot be empty")

		} else if !strings.Contains(email, "@") {
			fmt.Println("Invalid email address")

		} else if userTickets == 0 {
			fmt.Println("You must book at least 1 ticket")

		} else if userTickets <= remainingTickets {

			remainingTickets -= userTickets
			fullName := firstName + " " + lastName
			bookings = append(bookings, fullName)

			fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", fullName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Come back next year.\n")
				break
			}

		} else if userTickets > remainingTickets {

			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)

		} else {
			fmt.Println("Something went wrong, please try again")
		}
	}
}