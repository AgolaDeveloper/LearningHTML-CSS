//PRACTISING THE CORE CONCEPTS I COVERED ON my WEEK1 OF LEARNING
/*1. Arrays
2. Slices
3. Maps
4. Structs
5. Functions
*/

package main

import (
	"fmt"
)


	//1. ARRAYS
	//ways of Declaring/Defining Arrays

	//DECLARATION WITHOUT INITIALIZATION
	//var array [5]int //BASIC declarATION without initialization

	//var array1 = [5]int{}   //ASSIGNMENT declaration without initialization
	//array2 := [2]int{}      //SHORT-HAND NOTATION declaration without intialization
	//var array3 [10]int
	//= [...]int{} //INFERENCE Declaration without Initialization

	//PLAYING WITH ARRAYS
	//ADDING 10 ELEMENTS INTO OUR EMPTY ARRAY, array3

	/*var userInput1 int
	for j := 0; j < len(array3); j++ {
		fmt.Printf("ENTER ELEMENT %v(into the array):\n ", j+1)
		fmt.Scan(&userInput1)

		array3[j] = userInput1

	}
	fmt.Println("*************************")

	//accessing the array and printing every element
	for index, value := range array3 {
		fmt.Println(index, ":", value)
	}
	fmt.Println(array3)

	fmt.Println("###############################################################")
	//SORTING ARRAYS IN DECREASING ORDER

	for _, value := range array3 {
		for _, val := range array3 {
			if value > val {
				fmt.Println(value)
			}
			break
		}
	}

	//DECLARATION WITH INITIALIZATION
	//var array4 = [4]int{1, 2, 3, 4}        //ASSIGNMENT declaration with INITIALIZATION
	// array5 = [...]int{3, 5, 6, 10, 32} //INFERENCE declaration with INITIALIZATION
	//array6 := [4]int{11, 22, 33, 44}       //SHORT-HAND NOTATION declaration bwith INITIALIZATION

	//2. SLICES
	//ways of Declaring/Defining Slices

	//var slice []int             //Basic declaration without INITIALIZATION
	//var slice1 = []int{}        //ASSIGNMENT declaration without INITIALIZATION
	slice3 := make([]int, 0, 0) //DECLARATION by MAKE() method

	//adding 10 elements into slice3
	var userInt int

	for i := 0; i < 10; i++ {
		fmt.Printf("ENTER AN INTEGER ELEMENT %v:\n ", i+1)
		fmt.Scan(&userInt)
		//using append() to add the userInput into the slice
		slice3 = append(slice3, userInt)
	}

	for index, value := range slice3 {
		fmt.Println(index, ": ", value)

	}
	fmt.Println("_________________________________")

	fmt.Println(slice3)

	//slice4 := array[1:4]        //DECLARATION by slicing an exissting array or slice

	//3. MAPS
	//ways of declaring a map
	//var mapp map[int]string      //BASIC declaration of map without Declaration
	//map1 := make(map[int]string) //Declaring map using MAKE() method
	//map2 := map[int]string

	//BASIC declaration with INITIALIZATION
	//var map2 map[int]string={
	//	1:	"Array",
	//	2: 	"Slices",
	//	3. 	"Maps",
	//	4. 	"Structs",
	//}
	fmt.Println("###############################################################")

	map3 := map[int]string{
		1: "Array",
		2: "Slices",
		3: "Maps",
		4: "Structs",
	}
	for key, value := range map3 {
		fmt.Println(key, ": ", value)

	}*/

	//Creating amapof 5 counties with their respective Cholera cases

	choleraCases := make(map[string]int, 5)
	totalCases := 0
	var county string
	var countyCases int
	for i := 0; i < 5; i++ {
		fmt.Println("Enter County & Respective Cases of Cholera Recorded")

		fmt.Println("ENTER COUNTY:")
		fmt.Scan(&county)

		fmt.Printf("%v's Cases recorded?:", county)
		fmt.Scan(&countyCases)

		choleraCases[county] = countyCases

	}
	//our map
	fmt.Println(choleraCases)
	//counties & respective cases
	for key, value := range choleraCases {
		fmt.Printf("%v county has %v cases\n", key, value)
		//summing up total number of cases
		totalCases += value
	}

	fmt.Printf("#######################\n TOTAL NUMBER OF CASES: %v\n", totalCases)

	fmt.Println("#######################")

	//Accessing Cholera Cases for a Specific County; upon searching
	for {
		var choice string
		fmt.Printf("You wanna Check Cholera Cases?:\n1. Yes\n2. No\n ")
		fmt.Scan(&choice)

		if choice == "1" || choice == "Yes" || choice == "YES" || choice == "yes" {

			var cnty string
			fmt.Println("Which County:")
			fmt.Scan(&cnty)

			//	for _, value := range choleraCases {

			//Comparing the user's entry-search if it matches any county we already have in our database
			//if there's a match, cases for the respective county get printed
			//in case the entry-search matches nothing in the database...
			//...we let user know invalid choice & and give a retry to make valid choice
			//seems SWITCH can be BETTER in this case compared to ELSE...IF
			//**********************************************************************************************************************
			/*if cnty == key {
				fmt.Printf("CHOLERA CASES RECORDED IN %v: %v", cnty, value)

			} else {
				fmt.Printf("INVALID CHOICE\n")

			}*/
			//**********************************************************************************************************************
			temp, ok := choleraCases[cnty]
			fmt.Println(temp, ok)
			if ok == false {
				fmt.Printf("%v's cases aren't submitted yet", cnty)
			} else {

				fmt.Printf("CHOLERA CASES RECORDED IN %v: %v \n", cnty, choleraCases[cnty])
			}
			//}

			break

		} else if choice == "2" || choice == "No" || choice == "NO" || choice == "no" {
			break
		} else {
			fmt.Println("INVALID CHOICE: RETRY")
		}
	}

	

	delete(choleraCases, choleraCases[])

	//accessing and printing the elements in the map

	//STRUCTS
	//CREATING type struct

	/*type Person struct {
		name             string
		age              int
		status           string
		levelOfEducation string
	}

	//declaring our localvariables that would store values of our struct's items
	var nam string
	var agg int
	var stat string
	var lvlOfEd string

	//prompting users to enter the values to be stored at our declared variables...
	//... (before assigning them to the struct's items by defining an object/instance of Person struct)

	fmt.Println("ENTER NAME:")
	fmt.Scan(&nam)

	fmt.Println("ENTER AGE:")
	fmt.Scan(&agg)

	fmt.Println("ENTER STATUS:")
	fmt.Scan(&stat)

	fmt.Println("ENTER EDUCATION LEVEL:")
	fmt.Scan(&lvlOfEd)
	//defining a local variable of type Person

	var cousin3 = Person{
		name:             nam,
		age:              agg,
		status:           stat,
		levelOfEducation: lvlOfEd,
	}

	//accessing the properties/items/details of our object cousin3

	fmt.Println("YOUR AGE:", cousin3.age)
	fmt.Println("YOUR NAME:", cousin3.name)
	fmt.Println("YOUR STATUS:", cousin3.status)
	fmt.Println("YOUR LEVEL OF EDUCATION:", cousin3.levelOfEducation)
	fmt.Println("##########################################################")
	fmt.Println(cousin3)

	//SHORT-HAND declaration of variable type Person
	/*cousin4 := Person{
		name:             nam,
		age:              agg,
		status:           stat,
		levelOfEducation: lvlOfEd,
	}*/


