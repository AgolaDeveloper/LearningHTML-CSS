package main

import "fmt"

func loopSlice() {

	fmt.Printf("WELCOME: You wanna PLAY WITH SLICES? We gotchu\n\n Get Started with first slice, slice1!:\n")

	var slice []int
	var userInt int
	sliceSum := 0

	for len(slice) < 10 {
		//var slice []int
		//var userInt int
		fmt.Println("Kindly add an integer element to the Slice:")
		fmt.Scan(&userInt)
		slice = append(slice, userInt)

		fmt.Printf("\n Our User Input is %v \n", userInt)
		fmt.Println("Now, here is our Slice with the current elements updated:", slice)

		/*if len(slice) >= 5 {
			// we're endin our loop
			fmt.Printf(" we're now playing above length 5 of our slice:\n")
			//break
		}*/

		fmt.Println("length:", len(slice))
		fmt.Println("Capacity:", cap(slice))

		//now summing up the elements we already have in our slice

		sliceSum += userInt //this adds every user input to the previous value of our sliceSum

	}

	fmt.Printf("The sum of elements in our slice is %v \n\n", sliceSum) //printing the sum of elements

	//******************************************************************************************************************//

	//SECOND SLICE, slice1

	for {
		var userChoice string
		fmt.Println("DO YOU WANNA CONTINUE TO 2nd SLICE? (Y/N):")
		fmt.Scan(&userChoice)

		if userChoice == "Y" || userChoice == "y" {

			//APPENDING ANOTHER SLICE TO parent slice
			//we are asking user to add elements into second slice, slice1
			//we only need 15 here, so we must control how many elements our user inputs
			var slice1 []int
			var userInt1 int
			slice1Sum := 0

			fmt.Printf("\nWELCOME TO OUR SECOND SLICE, slice1. Now ")

			for len(slice1) < 15 {
				fmt.Println("Add element to the slice:")
				fmt.Scan(&userInt1)
				slice1 = append(slice1, userInt1)

				fmt.Printf("\n Our User Input is %v \n", userInt)
				fmt.Println("Slice1:", slice1)

				/*if len(slice) >= 5 {
					// we're endin our loop
					fmt.Printf(" we're now playing above length 5 of our slice:\n")
					//break
				}*/

				fmt.Println("length:", len(slice1))
				fmt.Println("Capacity:", cap(slice1))

				slice1Sum += userInt1

			}
			fmt.Printf("The sum of elements in our slice1 is %v \n\n", slice1Sum) //printing the sum of elements
			slice = append(slice, slice1...)                                      //appending CHILD SLICE, slice1 to our PARENT SLICE, slice

			fmt.Printf("After appending slice1 to first slice, here's what we have: %v \n", slice)
			fmt.Println("LENGTH:", len(slice))
			fmt.Println("CAPACITY:", cap(slice))
			fmt.Printf("ALL ELEMENTS IN UPDATED SLICE:")
			/*for i := 0; i < len(slice); i++ {
				fmt.Println(slice[i])
			}*/

			/*var index int
			for index, userInt = range slice {
				fmt.Println(index, ":", userInt)

			}*/

			for _, userInt = range slice {
				fmt.Println(userInt)

			}
			break

		} else if userChoice == "N" || userChoice == "n" {
			fmt.Println("ALRIGHT. IT WAS NICE HAVING YOU; BYE")
			break

		} else {
			if userChoice == "1" || userChoice == "a" {

				fmt.Printf("INVALID CHOICE...RETRY!\n Enter:\n Y or y (to continue) \nN or n (to Exit)\n") //How do we get the user back to retry...

			} else {

				fmt.Printf("INVALID CHOICE...RETRY!\n Enter:\n Y or y (to continue) \nN or n (to Exit)\n") //How do we get the user back to retry...
				//... entering a choice once their initial choice prove invalid
				continue
			}

		}

		fmt.Printf("%v INVALID CHOICE...\n", userChoice) //How do we get the user back to retry...

	}
}
