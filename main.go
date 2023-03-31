package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceName = "A2G Store"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Println("welcome to our conference", conferenceName)
	fmt.Printf("we have a total of %v tickets and this many %v are stil available\n", conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("add your first name")
		fmt.Scan(&firstName)

		fmt.Println("add your last name")
		fmt.Scan(&lastName)

		fmt.Println("add your email")
		fmt.Scan(&email)

		fmt.Println("number of tickets")
		fmt.Scan(&userTickets)

		
		remainingTickets = remainingTickets - userTickets
		
		var fullName = firstName + " " + lastName
		bookings = append(bookings, fullName)


		 firstNames := []string{}


		 //not clear yet
		 for _, booking := range bookings {
			var names = strings.Fields(booking)
		    firstNames = append(firstNames, names[0])
		 }


		fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)
		fmt.Printf("%v\n", firstNames)
	}
}

