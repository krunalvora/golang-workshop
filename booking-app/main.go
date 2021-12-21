package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// package level variables
const conferenceTickets = 50

var conferenceName = "Go Conference" // is the same as conferenceName := "Go Conference" but := cannot be used on package level
var remainingTickets uint = 50       // optionally, we can mention the type of the variable after the name like unsigned int
var bookings = make([]UserData, 0)   // list of structs

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{} // Need this to make sure main thread waits for all other threads to complete

func main() {
	greetUsers()
	for { // infinite loop same as for true {}
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidTicketNumber && isValidEmail && isValidName {

			bookTickets(userTickets, firstName, lastName, email)

			wg.Add(1)                                               // Incrementing the wait group thread counter
			go sendTickets(userTickets, firstName, lastName, email) // Adding the go keyword runs this func in a new non blocking thread and deletes the thread once completed

			firstNames := getFirstNames()
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			var noTicketsRemaining bool = remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is full. Come back next year.")
				break // exit loop
			}
		} else {
			if !isValidName {
				fmt.Println("First or last name entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address does not contain @.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets entered is invalid.")
			}
		}
		fmt.Printf("We have a total of %v tickets and %v are still remaining.\n", conferenceTickets, remainingTickets)
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}           // slice as opposed to fixed length array => firstNames: [50]string
	for _, booking := range bookings { // for each: first variable is index but if not using it, use _ which is the blank identifier
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets to book:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you '%v %v' for booking %v tickets.\n", firstName, lastName, userTickets)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("Sending ticket: \n %v \nto email address: %v\n", ticket, email)
	fmt.Println("##############")
	wg.Done() // Decrease the wait group thread counter
}
