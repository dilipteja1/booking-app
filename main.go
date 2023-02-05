package main

import (
	"fmt"
	"booking-app/helper"
	

)
var confName = "Go conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	FirstName string
	LastName string
	Email string
	numberOfTickets uint
}

func main() {
	
	greetUsers()

	//infinite for loop
	for {

		userFirstName, userLastName, userMailID, userTickets := getUserInputs()
		//validating input

		isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(userFirstName, userLastName, userMailID, userTickets,remainingTickets)

		if isValidName && isValidEmail && isValidTicket {
			bookTicket(userTickets,userFirstName,userLastName,userMailID)

			
			var firstNames = getFirstNames()
			fmt.Printf("all our bookings : %v\n", firstNames)
			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				//end program
				fmt.Println("our conference tickets are sold out .please visit next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name is not valid")
			}
			if !isValidEmail {
				fmt.Println("Invalid email address entered")
			}
			if !isValidTicket {
				fmt.Println("invalid ticket number entered")
			}
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("conferenceTickets is %T,remainingTickets is %T, confName is %T\n", conferenceTickets, remainingTickets, confName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	var firstNames []string

	// '_' this is a blank identifier, used when a variable has to be initialized but we dont use it.
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.FirstName)
	}

	return firstNames
}


func getUserInputs() (string, string, string, uint) {
	//array with fixed size
	//var bookings [50]string

	//slice which is a dynamic way of slice.like vectors in cpp
	//ask user for their
	//here for scan function.. we need to provide the pointer location of the variable
	//instead of the variable name.
	var userFirstName string
	var userLastName string
	var userMailID string
	var userTickets uint

	fmt.Println("Enter the user First name : ")
	fmt.Scan(&userFirstName)

	fmt.Println("Enter the user Last name : ")
	fmt.Scan(&userLastName)

	fmt.Println("Enter the user mail address : ")
	fmt.Scan(&userMailID)

	fmt.Println("Enter the number of tickets to be booked : ")
	fmt.Scan(&userTickets)

	return userFirstName, userLastName, userMailID, userTickets
}

func bookTicket(userTickets uint, userFirstName string, userLastName string, userMailID string){
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		FirstName : userFirstName,
		LastName : userLastName,
		Email : userMailID,
		numberOfTickets : userTickets,
	}
	

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings : %v\v",bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v.\n", userFirstName, userLastName, userTickets, userMailID)
	fmt.Printf("remaining tickets are %v\n", remainingTickets)
}
