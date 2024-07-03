package main

import "fmt"

//TRYNNA play around with Structs
//...while grasping slices too
//Creating our own, complex, data type; STRUCT

 	type Workers struct {
		//these are the FIELDS that our struct has
		firstName   string
		lastName    string
		age         int
		staffNumber string
		nextOfKins  []string //one of them fields is a slice
	}

	//Once we have our complex data type (STRUCT) created...
	//...we're now free to create/declare and initialize COMPLEX variables of type struct
	//...plus local variable of basic data types that'll store values for teh STRUCT'S FIELDS

	//DECLARATION
	//we first declare variables that'll store values for our STRUCT'S FIELDS

	var fName string
	var lName string
	var agee int
	var staffNum string
	var nxtOfKin []string

	//declaring worker1 as variable of type Workers
	//assigns its fields with values retreived from our local variables too

	//Stilll, we can declare a local variable of type struct, while initializing at the same time

	/*var worker2 = Workers{
		firstName:   fName,
		lastName:    lName,
		age:         agee,
		staffNumber: staffNum,
		nextOfKins:  nxtOfKin,
	}*/

	//now lets take user inputs from our 2 employees

	fmt.Println("Enter Your Personal Information Here")

	fmt.Println("FIRST NAME:")
	fmt.Scan(&fName)

	fmt.Println("LAST NAME:")
	fmt.Scan(&lName)

	fmt.Println("YOUR AGE:")
	fmt.Scan(&agee)

	fmt.Println("STAFF NUMBER:")
	fmt.Scan(&staffNum)

	fmt.Println("")

	var noKins int
	//var userChoice string
	fmt.Println("How Many Next of Kins would you love to add?:")
	fmt.Scan(&noKins)

	for i := 0; i < noKins; i++ {
		var nxtOfKinName string
		fmt.Println("Add Kin (First Name only):")
		fmt.Scan(&nxtOfKinName)

		nxtOfKin = append(nxtOfKin, nxtOfKinName)
		fmt.Scan(&nxtOfKin)

	}

	fmt.Println("YOUR NEXT-OF-KINS SO FAR", nxtOfKin)
	fmt.Println("")

	var userChoice1 string
	fmt.Printf("Would you like to add more Kin(s),\n...in case there is someone you forgot to include?\n 1. Yes\n2. No \n")
	for {
		fmt.Scan(&userChoice1)

		if userChoice1 == "1" || userChoice1 == "Yes" || userChoice1 == "yes" || userChoice1 == "YES" {
			fmt.Printf("ALRIGHT %v %v, make it quick!\nSo ", fName, lName)

			var noKins1 int
			fmt.Println("How Many Next of Kins would you love to add?:")
			fmt.Scan(&noKins1)

			for j := 0; j < noKins1; j++ {
				var nxtOfKinName1 string
				fmt.Println("Add Kin (First Name only):")
				fmt.Scan(&nxtOfKinName1)

				nxtOfKin = append(nxtOfKin, nxtOfKinName1)
				//fmt.Scan(&nxtOfKin)
			}
			break

		} else if userChoice1 == "2" || userChoice1 == "No" || userChoice1 == "no" || userChoice1 == "NO" {
			break
		} else {
			fmt.Printf("ACHA UJINGA %v: %v is INVALID CHOICE...\n Enter 1 to continue adding more kins or,\n2 to exit/RETRY", fName, userChoice1)
			fmt.Println("\n 1. Yes\n2. No \n:")

		}
	}

	var worker1 = Workers{
		firstName:   fName,
		lastName:    lName,
		age:         agee,
		staffNumber: staffNum,
		nextOfKins:  nxtOfKin,
	}
	//fmt.Scan(&nxtOfKin)

	fmt.Println("HERE ARE YOUR INFORMATION")
	fmt.Println(worker1)

	/*fmt.Println(worker1.firstName)
	fmt.Println(worker1.lastName)
	fmt.Println(worker1.age)
	fmt.Println(worker1.staffNumber)
	fmt.Println(worker1.nextOfKins)*/


