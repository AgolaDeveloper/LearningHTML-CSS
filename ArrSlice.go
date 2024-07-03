package main

import (
	"fmt"
)

func arrayz() {
	//Declaring an array

	//var arr[] int
	var array = [5]string{"Peter", "John", "Zoe", "Luke", "Sue"}
	var array1 = []int{}
	var array11 []int //declaring slice, the basic way.

	const sNam string = "Agola"
	slice1 := make([]int, 4) //declaring slice,using make method.
	fmt.Printf("The length of this size is %v and its capacity is %v", len(slice1), cap(slice1))

	array2 := [2]float64{2.1, 5.6}
	fmt.Println(array2)
	var array3 = []int{5, 6, 9, 10, 4}
	fmt.Println("***\nlength of array3 is ", len(array3))
	fmt.Printf("\n \n")

	array1 = append(array1, 11)
	array3 = append(array3, 11)
	slice := array[1:5] // declaring slice

	fmt.Println("length of my slice is", len(slice), " and it's capacity is", cap(slice))

	fmt.Println("length of array is", len(array))
	fmt.Println("length of array1 is", len(array1))
	fmt.Println("length of array11 is", len(array11))
	fmt.Println("length of array2 is", len(array2))
	fmt.Println("New length of array3 is", len(array3))
	fmt.Println("Welcome to the", sNam, "'s FAMILY")

	array11 = append(array11, array3...)
	fmt.Println("New length of array11 is", len(array11))

	for j := 0; j < len(array); j++ {
		fmt.Println("Child", j+1, ":", array[j], sNam)
	}

	for i := 0; i < len(array3); i++ {
		fmt.Println(array3[i])
	}

	array33 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(array33)
	fmt.Println("length of array33 is", len(array33), "and its capacity is", cap(array33))

}
