package main

import "fmt"


	//creating a slice, myArray using make method
	myArray := make([]int, 0)
	j := 0
	var val int

	//this block of code allows user to input inegers into our myArray slice

	for k := 0; k < 10; k++ {
		fmt.Printf("Enter int vaalue into the array: \n")
		fmt.Scan(&val)
		myArray = append(myArray, val)

	}

	//this block of code prints our slice, myArray so that we see values we got
	fmt.Println(myArray)

	//this block of code checks if there are numbers in the list that are duplicates
	for index, value := range myArray {
		for i, v := range myArray {
			if index < i && value == v {
				j++
				fmt.Printf("Value %v  at index %v is duplicate at index %v\n", value, index, i)

			}
		}

	}
	//noDuplicate:
	fmt.Printf("we had %v duplicates \n", j)


