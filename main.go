package main

import (
	"fmt"

)

func main() {
	var conferenceName = "Golang confernece"
	const conferenceTickets = 50
	// var bookings []string

	// var firstName, lastName string

	// for {

	// 	fmt.Println("welcome")
	// 	fmt.Println("please enter your first name")
	// 	fmt.Scan(&firstName)
	// 	fmt.Println("please enter your last name")
	// 	fmt.Scan(&lastName)

	// 	bookings = append(bookings, firstName+" "+lastName)
        
	// 	for index,booking := range bookings{
    //        var names = strings.Fields(booking)
	// 	    initial := names[0][0]
	
	// 	   fmt.Printf("names: %v\n", names)
	// 	   fmt.Printf("index: %v\n", index)
	// 	   fmt.Printf("initial: %v\n", initial)
	// 	}

	// 	fmt.Printf("%v  welcome to %v ", bookings[len(bookings) - 1], conferenceName)
    //     fmt.Printf("bookings: %v\n", bookings)

	// }
 initials := conferenceName[0]
   fmt.Printf("conferenceName: %c\n", conferenceName[0])

   if (initials  == 'G'){
     println("G matches",initials)
   }
}
