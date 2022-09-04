package main

import "fmt"

func main(){
	var conferenceName = "Golang confernece"
	const conferenceTickets  = 50
	var bookings []string

	var firstName , lastName string
    
    fmt.Println("welcome")
	fmt.Println("please enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("please enter your last name")
	fmt.Scan(&lastName)

    bookings = append(bookings,firstName + " " + lastName) 

	fmt.Printf("%v  welcome to %v ",bookings[0],conferenceName)

 

    fmt.Printf("conferenceName: %v\n", conferenceName)
	fmt.Printf("bookings: %v\n", bookings)

}
