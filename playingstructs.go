// we wanna play around with structs
// today, we're creating a slice of structs
// the structs shall have several objects
package main

import "fmt"

//creating a function that'll welcome vistors to our booking app

func welkam() {
	//this is more of a home page
	fmt.Println("WELCOME TO TENE BOOKING APPLICATION")

}

//we first create our STRUCT first
//attendees booking for the concert tickets shall have create these:
//1. first name
//2. second name
//3. age
//3. number of tickets
//4. email address
//5. phone number
//6. address
type Attendee struct {
	firstName       string
	lastName        string
	age             int
	numberOfTickets int
	email           string
}


	welkam()
	var cont int

	fmt.Println("Press 1 to continue to our booking page ")
	fmt.Scan(&cont)
	if cont == 1 {
		goto contin
	} else {
		fmt.Println("Ooops... A'ight. BYE!")
		return
	}

	//we declare/create local variables that'll hold values to the struct's items
contin:

	var fName string
	var lName string
	var ag int
	var numOfTickets int
	var ema string

	//capacity of our concert is 5
	//so we'll only have 5 tickets available for booking

	//const
	totalTickets := 5
	//var ticketsBooked int
	var remainingTickets int

	//we'll have to know the tickets remaining after every booking
	//to help with avoiding buffer overflow; knowing whether there's possibility of further bookings

	//remainingTickets = totalTickets - numOfTickets

	//creating an empty slice that'll be holding details of all the attendees, who've booked
	attending := make([]Attendee, 0)

	//here's where booking starts... more inna a loop
	//however, we'll allow for booking untill remaining tickets is
	remainingTickets = totalTickets - numOfTickets
	count := 0

	for {
		count++

		fmt.Println("Go AHEAD AND BOOK A PLACE; ENTER ALL THESE: ")
		fmt.Println("NUMBER OF TICKETS")
		fmt.Scan(&numOfTickets)
		//checking if number of tickets booked is more than the capacity
		if numOfTickets > remainingTickets {
			fmt.Printf("Cannot book %v tickets: Our Capacity is %v \n", numOfTickets, remainingTickets)
			fmt.Printf(" Max number of tickets you can book is %v\n", remainingTickets)
			continue
		} else if numOfTickets == 0 {
			fmt.Println("********* INVALID NUMBER OF TICKETS. Can't book 0 tickets ********************")
			continue
		} else {
			goto add
		}
	add:

		fmt.Println("FIRST NAME: ")
		fmt.Scan(&fName)
		fmt.Println("LAST NAME: ")
		fmt.Scan(&lName)
		fmt.Println("AGE: ")
		fmt.Scan(&ag)

		if ag < 18 {
			fmt.Println("This Concert is only for adults: age 18 years and above")
			fmt.Printf("%v year-olds aren't allowed to attend\n", ag)
			continue
		}
		fmt.Println("EMAIL: ")
		fmt.Scan(&ema)

		//now we create an object of our struct
		Attende := Attendee{
			firstName:       fName,
			lastName:        lName,
			age:             ag,
			numberOfTickets: numOfTickets,
			email:           ema,
		}
		remainingTickets = totalTickets - numOfTickets

		attending = append(attending, Attende)

		fmt.Println("BOOKED TICKETS: ", numOfTickets)
		fmt.Println("REMAINING TICKETS: ", remainingTickets)

		totalTickets = remainingTickets
		if remainingTickets == 0 {
			fmt.Println("THE CONCERT IS FULLY BOOKED... There's always a next time: SEE YOU!")
			break
		}
	}
	fmt.Println("*****************************************")

	fmt.Printf("%v people booked for the concert so far\n Here are their details: \n ", count)

	for index, value := range attending {
		index++
		fmt.Printf("ATTENDEE %v: %v \n", index, value.firstName)
		fmt.Println("TICKETS: ", value.numberOfTickets)
		fmt.Println("LAST NAME: ", value.lastName)
		fmt.Println("EMAIL: ", value.email)
		fmt.Println("AGE: ", value.age)
		fmt.Println("*****************************************")
		fmt.Println()

	}

}
