package main

import "fmt"

 
	//let'splay around with maps, in Golang

	var userKey int
	var userValue string

	var userKey1 string
	var userValue1 int

	//first creating an empty map

	myMap := make(map[int]string, 0)
	var maMap map[string]int

	//defining and initializing predifined "key-value pairs"
	mMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	//looping through our mMap
	for key, value := range mMap {
		fmt.Printf("Key>> %v: %v", key, value)
	}
	fmt.Println("#################################### ")
	fmt.Println("###########WE SHOULD NOW BE DELETING AN ELEMENT FROM THIS MAP ############### ")
	var toDelete string
	fmt.Println(" Enter key to delete")
	fmt.Scan(&toDelete)

	//we first checkif the key entered is in the map
	//using "temp, ok" operator for this

	temp, ok := mMap[toDelete]

	if ok == false {
		fmt.Printf("sorry %v doesn't exist in our map", toDelete)
	} else {
		delete(mMap, toDelete)
	}

	fmt.Println(mMap)
	fmt.Println("#################################### ")

	//adding data elements into our, myMap, map
	var userNam int
	fmt.Println("IT'S myMap: How many elements/keys Do you wanna add?: ")
	fmt.Scan(&userNam)

	for i := 0; i < userNam; i++ {
		fmt.Println("Add key to myMap")
		fmt.Scan(&userKey)
		fmt.Println("Enter its respective value: ")
		fmt.Scan(&userValue)

		//now we're matching the user's key to its respective value
		myMap[userKey] = userValue
	}

	//adding data elements to our, maMap, map
	var userNam1 int
	fmt.Println("IT'S maMap: How many elements/keys Do you wanna add?: ")
	fmt.Scan(&userNam1)

	for i := 0; i < userNam1; i++ {
		fmt.Println("Add key to myMap")
		fmt.Scan(&userKey1)
		fmt.Println("Enter its respective value: ")
		fmt.Scan(&userValue1)

		//now we're matching the user's key to its respective value
		maMap[userKey1] = userValue1
	}

	//we're printing out the elements of 'em maps?
	fmt.Println("myMap ", myMap)
	fmt.Println("#******************#")
	fmt.Println("maMap: ", maMap)


