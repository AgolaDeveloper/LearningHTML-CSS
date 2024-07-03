package main

import (
	"fmt"
)

func mySlice() {

	var arr [7]int
	//arr[3] = 6
	for i := 0; i < len(arr); i++ {
		if (i < 3) && (i != 0) {
			arr[i] = i * len(arr)
		} else if i > 3 {
			arr[i] = i * i
		} else {
			continue
		}
	}

	fmt.Println("Array's Elements:", arr)

	//*****************************************************************************************************

	//3 WAYS OF CREATING SLICES
	var slice []int //BASIC Declaration of a SLICE
	fmt.Println("Slice Elements:", slice)

	var slice1 = arr[4:7] //Slicing an existing ARRAY
	fmt.Println("Slice1 Elements:", slice1)

	slice2 := make([]int, 9) //Using make() method
	fmt.Println("SLice2 Elements:", slice2)
	/*
		for j := 0; j < len(slice2); j++ {
			slice2 = append(slice2, j)
		}
		fmt.Println("NEW Slice2 Elements:", slice2)*/

	//WORKING WITH SLICES

	slice = append(slice, 98, 100, 3) //we're using append() method to add elements to our slice
	fmt.Println("NEW Slice Elements:", slice)

	fmt.Println("What's your name?")
	var nam string
	fmt.Scan(&nam)

	fmt.Printf("WELCOME Ms. %v, to our Truth & Dare Game \nRAMBANYA you look so Gorgeous", nam)
	fmt.Println(" ")

	fmt.Printf("SO %v, would you tell us (for the start) Who your Love is, (like boo?): \n", nam)

	var lov string
	fmt.Scan(&lov)

	fmt.Printf("Thanks Ms. %v, for letting us know that %v ema ng'othi (jamayii) \n", nam, lov)
	fmt.Printf("\nSani %v ma-jamayi-no ee chuori. Iwinnjo Ms. %v", lov, nam)
	fmt.Printf("\n\nUNFORTUNATELY, your %v postponed the game. \nIt's re-scheduled for next weekend \n \nLOVE U...Gud 9t \n\n", lov)

}
