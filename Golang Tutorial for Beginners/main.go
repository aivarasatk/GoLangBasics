package main

import (
	"booking-app/common"
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName string
	lastName  string
	tickets   int
}

var packageLevelValue = "Package level"
var wg = sync.WaitGroup{}

func main() {

	routinesMain()

	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	common.GreetUsers()

	fmt.Printf("Welcome to %v booking application\r\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "and", remainingTickets, "are still remaioning")

	bookings := []string{}
	firstNames := []string{}
	bookings = append(bookings, "Nana")
	bookings = append(bookings, "bana")

	var userData = UserData{
		firstName: "Me",
		lastName:  "last me",
		tickets:   69,
	}

	userData.lastName = "new me"
	array, number := getUsers()

	fmt.Print(array, number)
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Enter first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter email:")
		fmt.Scan(&email)

		fmt.Println("Enter count:")
		fmt.Scan(&userTickets)

		remainingTickets -= userTickets

		firstNames = append(firstNames, firstName)
		for index, element := range firstNames {
			fmt.Println("Foreach", index, element)
		}
		fmt.Println(firstName, lastName, email, userTickets)
		fmt.Println("Remaining tickets", remainingTickets)

		endOfTickets := remainingTickets <= 0
		if endOfTickets {
			break
		}
	}

	wg.Add(10)
	for index := 0; index < 10; index++ {
		var unicode string = "German Umlauts - ä, ö, ü"
		fmt.Println(unicode, "žč ä, ö, ü Spinning user processing")

		go sendTicket("Me", uint(index), "email@e.com")
		fmt.Println("Starting new loop")
	}

	wg.Wait()
}

func getUsers() (books []string, numbers int) {
	return []string{}, 15
}

func sendTicket(firstName string, tickets uint, email string) {
	time.Sleep(time.Duration(tickets) * time.Second)
	var ticket = fmt.Sprint(tickets, "tiickets for", firstName)

	fmt.Printf("sending %v to %v", ticket, email)
	fmt.Println()

	wg.Done()
}

func routinesMain() {
	c := make(chan int) // Create a channel to pass ints
	for i := 0; i < 5; i++ {
		go cookingGopher(i, c) // Start a goroutine
		time.Sleep(time.Millisecond * 200)
	}

	for i := 0; i < 5; i++ {
		gopherID := <-c // Receive a value from a channel
		fmt.Println("gopher", gopherID, "finished the dish")
	} // All goroutines are finished at this point
}

/* Notice the channel as an argument */
func cookingGopher(id int, c chan int) {
	fmt.Println("gopher", id, "started cooking")
	time.Sleep(time.Second * 2)
	c <- id // Send a value back to main
}
