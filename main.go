package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// Arrays var bookings = [50]string{}
// var bookings = make([]map[string]string, 0) // var bookings []string // create a slice

// package level variables
const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

// structures over maps - user mixed data types and you can set the properties and map is just an empty list
type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

var wg = sync.WaitGroup{}

func main() {

	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets)

	fmt.Printf("conferenceName is %T, remaining tickets is %T, conferenceTickets is %T\n", conferenceName, remainingTickets, conferenceTickets)

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {
		firstName, lastName, email, userTickets := getUserInput()
		isValidEmail, isValidName, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		//check if user tickets is greater than remainnder
		if isValidName && isValidEmail && isValidTickets {
			bookTicket(userTickets, firstName, lastName, email)
			// tells main function that it needs to wait for other threads to finish before exiting
			//wg.Add(1)
			//move this to another thread so code can continue
			go sendTicket(userTickets, firstName, lastName, email) //spins off a green thread and spins off the goroutines. it is an abstraction of actual thread

			firstNames := getFirstName()
			fmt.Printf("The first names of bookings are : %v\n", firstNames)

			noTicketsremaining := remainingTickets == 0
			if noTicketsremaining {
				//end program
				fmt.Println("Our conference is fully booked. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first or last name entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address entered doesn't contain @")
			}
			if !isValidTickets {
				fmt.Println("number of tickets is invalid")
			}
		}
		//wg.Wait()  just waits
	}

	// city := "London"
	// switch city {
	// 	case "New York":
	// 		//code
	// 	case "Singapore":
	// 		// code
	// 	case "London", "Berlin":
	// 		// code
	// 	case "Mexico City":
	// 		// code
	// 	default:
	// 		// when none are true
	// }
}

func greetUsers() {
	fmt.Printf("Welcome to the %v conference!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string {
	firstNames := []string{} // blank identifiers used to mark for variables that are not being used
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask the user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["tickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTickets,
	}

	// bookings = append(bookings, firstName+" "+lastName) // slice of string
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	//simulate generating a ticket and save formatting string
	// concurrency is cheap and easy in go sleep stops or blocks the thread
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#########################")
	fmt.Printf("sendinng ticket: %v\n to email address %v\n", ticket, email)
	fmt.Println("#########################")
	// wg.Done() removes the thread from the waitlist
}
