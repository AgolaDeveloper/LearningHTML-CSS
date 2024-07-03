package main

import (
	"fmt"
)

//this file has a function
//the function's parameter-less but,...
//statements revolve around arays
//...literally everything I've learn about Go Arrays

func arayz() {
	fmt.Println("Whatever I learnt in the Array session")
	/*DIFFERENT WAYS OF CREATING AN ARRAY
	Creation of arrays in Go lies in 2 categories:
		1. Basic Declaration, without initialization
		2. creation by Initialization (creating and assigning values at a go)*/

	//1. BASIC DECLARATION (WITHOUT INITIALIZATION)
	var array [11]int //BASIC Declaration, without initialization
	fmt.Println("array's elements:", array)
	//here, we've created an array of type integer..
	//.. and length 3 ; no initialization of elements yet

	//2. CREATION BY INITIALIZATION (ASSIGNMENT ELEMENTS/VALUES AT A GO)
	var array1 = [4]int{2, 4, 5, 10} //Declaration with INITIALIZATION, at a go
	fmt.Println("array1's elements:", array1)

	//here, our array is created and initialized(...
	//...assigned values/elements) at the same time
	//ARRAY TYPE: integer of length 4

	var array2 = [...]string{"Go", "Is", "Fun", "for real", "!"} //INFERRED INITIALIZATION
	fmt.Println("array2's elements:", array2)
	fmt.Println("array2's length:", len(array2))

	//Here, we create our array and assign values...
	//without specifying the length
	//the length will be infereed by the COMPILER...
	//..depending on the numbe rof values/elements assigned

	array3 := [5]string{"Yu", "Ye", "ZII", "Kolo", "Go"} //ShortHand-notation Declaration
	fmt.Printf("array3's 2nd and last element respectively: %v & %v", array3[1], array3[4])
	fmt.Println("")
	fmt.Println("array3's length:", len(array3))
	fmt.Println("array3's last element using len() method:", array3[len(array3)-1])

	//here, we use the ':=" operater to create our array
	//remember, shoerthand notation decalaration must..
	//..aoutomatically come with initialization
	//otherwise, it will throw an erro during compile time.
	var array4 = [8]int{0: 8, 5: 11} //SPECIFIC INITIALIZATION, using index
	fmt.Println(array4)
	fmt.Println("array4's elements:", array2)
	fmt.Println("array4's length:", len(array4))

	var array5 = [8]int{1, 3, 22, 5} //PARTIAL INITIALIZATION
	fmt.Println("array5's elements:", array5)
	fmt.Println("array5's length:", len(array5))
	fmt.Println("array5's last element:", array5[6])

	//HERE, we create and initialize the array with...
	//...a fewer elements (assigned) compared to the lenght indicated

	//WORKING WITH ARRAYS

	//Accessing Arrays
	sum := 0 //initializing variable 'sum' that'll sum up array's values during the iterations
	for i := 0; i < len(array); i++ {

		//if array[i] == 0 {
		array[i] = i + 2
		/* else {
			fmt.Println(array[i], "!= 0")
		}*/
		fmt.Println(array[i])

		sum += array[i] //variable sum is set to
	}
	//fmt.Println(array)
	fmt.Println("Sum of array's values/elements is", sum)

	fmt.Println("ORIGINAL ELEMENTS OF array4:", array4)
	tot := 0

	for j := 0; j < len(array4); j++ {
		if array4[j] != 0 { //checking element that's not 0 (already initialized) in array4...
			//continue
			array4[j] = len(array4) * 10 //...then CHANGES its VALUE to array4's length
		} else {
			array4[j] = j * len(array4) //else, the values with zero values (uninitialized) are assigned
			// values computed by multiplying the array4's length by respective index
		}
		tot += array4[j] //here, we get the total of new array4's elements

	}
	fmt.Println("NEW ELEMENTS OF array4:", array4)
	fmt.Println("Sum of New ELEMENTS OF array4:", tot)

	//Changing/Modifying Arrays

	//Copying one aray to the other (must be of same type&length though)
	array5 = array4 //this assigns elements of array4 to array5; deleting & replacing initial/original elements of array4
	fmt.Println("2nd NEW ELEMENTS OF array4:", array4)
	fmt.Println("NEW ELEMENTS OF array5:", array5)

}
