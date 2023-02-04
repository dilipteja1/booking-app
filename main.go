package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint= 50
	fmt.Printf("conferenceTickets is %T,remainingTickets is %T, conferenceName is %T\n",conferenceTickets,remainingTickets,conferenceName)
	fmt.Println("welcome to", conferenceName, "booking application")
	fmt.Printf("we have total of %v tickets and %v are still available.\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")
	
	//slice with fixed size
	//var bookings [50]string

	//slice which is a dynamic way of slice.like vectors in cpp
	var bookings []string

	var userFirstName string
	var userLastName string
	var userMailID string
	var userTickets uint

	//ask user for their 
	//here for scan function.. we need to provide the pointer location of the variable 
	//instead of the variable name.
	fmt.Println("Enter the user First name : ")
	fmt.Scan(&userFirstName)

	fmt.Println("Enter the user Last name : ")
	fmt.Scan(&userLastName)

	fmt.Println("Enter the user mail address : ")
	fmt.Scan(&userMailID)

	fmt.Println("Enter the number of tickets to be booked : ")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings,userFirstName +" "+ userLastName)

	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v.\n",userFirstName,userLastName,userTickets,userMailID)
	fmt.Printf("remaining tickets are %v\n",remainingTickets)

	fmt.Printf("the whole slice is : %v\n",bookings)

}	
