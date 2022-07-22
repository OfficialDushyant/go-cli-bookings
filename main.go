package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v  tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTicket uint

		// ask the user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTicket)

		remainingTickets = remainingTickets - userTicket
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTicket, email)
		fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of the bookings are: %v.\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is fully booked. Come back next year")
			break
		}
	}

}
