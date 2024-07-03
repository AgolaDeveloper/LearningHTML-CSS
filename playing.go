package main

import "fmt"

 
	//playing with slices

	//creating the 1st, empty, slice then prompting user to enter its data elements
	var userInput int
	mySlice := make([]int, 0)

	fmt.Println("Enter integer values; or E to stop/exit entering more values")

	for i := 0; i < 6; i++ {
		fmt.Println("add an integer value to our slice")
		fmt.Scan(&userInput)

		//JUST USED := AND = IN USING MAKE METHOD, to append elements into our slice... which is right way to do it?
		mySlice = append(mySlice, userInput)

	}

	fmt.Println("********************************************")
	for _, value := range mySlice {
		fmt.Printf("%v ", value)
	}
	fmt.Println(" ")
	//now we are creating 2nd slice by slicing the 1st slice
	fmt.Printf("length of mySlice: %v \n", len(mySlice))

	fmt.Println("")
	slice2 := mySlice[2:len(mySlice)]

	for _, value := range slice2 {
		fmt.Printf("%v ", value)
	}

	var isEqual bool
	isEqual = mySlice[len(mySlice)-1] == mySlice[5]

	fmt.Println("")
	fmt.Println(isEqual)


