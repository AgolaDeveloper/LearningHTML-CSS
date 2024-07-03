// Here, we cover every basic concept of arrays
package main

import "fmt"

//function to modify our originalArray
func modifyArray(arr [5]int, index, element int) [5]int {
	arr[index] = element
	return arr
}

 
	//DECLARATION OF our ORIGINAL ARRAY, of size 3
	originalArray := [5]int{}

	//var arr [2]int

	//populating our original array
	var element int
	for i := 0; i < len(originalArray); i++ {
		fmt.Println("Enter array-element [type integer]")
		fmt.Scan(&element)

		originalArray[i] = element
	}

	//COMMON OPERATIONS ON ARRAYS
	// 1. Copying an Array
	// 2. Modifying an Array
	// 3. Getting length of an Array
	// 4. Comparing Arrays
	// 5. Iterating an Array

	//1. MAKING A COPY OF OUR ORIGINAL ARRAY
	//...here, the copiedArray shall be equal to originalArray since,we'd have copied all the elements to it
	copiedArray := originalArray

	//2. MODIFYING AN ARRAY
	var in2M int
	fmt.Printf("Enter index to modify... Your index should be <= %v\n", len(originalArray)-1)
	fmt.Scan(&in2M)

	var el2M int
	fmt.Println("Enter a replacement element")
	fmt.Scan(&el2M)

	modifiedArray := modifyArray(originalArray, in2M, el2M)

	//3. GETTING LENGTH OF ARRAY
	lenOfOrigArray := len(originalArray)
	lenOfCopyArray := len(copiedArray)
	lenOfModArray := len(modifiedArray)

	//4. COMPARING ARRAYS
	isEqual := originalArray == copiedArray
	isEqual1 := originalArray == modifiedArray
	isEqual2 := copiedArray == modifiedArray

	// 5.ITERATING ARRAYS
	// 5a. using Traditional For Loop

	fmt.Println("Our Original Array:")

	for i := 0; i < len(originalArray); i++ {
		fmt.Printf("originalArray[%v]: %v\n", i, originalArray[i])
	}
	fmt.Println("Length: ", lenOfOrigArray)
	fmt.Println()

	// 5a. using Range Loop
	fmt.Println("Our Modified Array:")

	for index, value := range modifiedArray {
		fmt.Printf("modifiedArray[%v]: %v\n", index, value)
	}
	fmt.Println("Length: ", lenOfModArray)
	fmt.Println()

	fmt.Println("Our Copied Array:")

	for index, value := range copiedArray {
		fmt.Printf("modifiedArray[%v]: %v\n", index, value)
	}
	fmt.Println("Length: ", lenOfCopyArray)
	fmt.Println()

	fmt.Printf("Are our originalArray and copiedArray equal?\n >> %v \n", isEqual)
	fmt.Printf("Are our originalArray and modifiedArray equal?\n >> %v \n", isEqual1)
	fmt.Printf("Are our modifiedArray and copiedArray equal?\n >> %v \n", isEqual2)

	//MOVING TO SLICES;with originalArray as our backing Array
	//creating a slice; by slicing the originalArray
	mySlice := originalArray[2 : len(originalArray)-1]
	fmt.Printf("Our slice before appending more elements: %v\n", mySlice)

	//Common Operations on Slices
	//1. Appending elements into the slice
	var myEle int
	fmt.Printf("How many elements do you wanna append to our slice?:\n >>")
	fmt.Scan(&myEle)

	var ele int
	for i := 0; i < myEle; i++ {
		fmt.Printf("Enter element to Append:\n >>")
		fmt.Scan(&ele)
		mySlice = append(mySlice, ele)
	}

	fmt.Printf("Our slice after appending %v elements: \n %v >>\n", myEle, mySlice)

	fmt.Printf("ANOTHER OPERATION ON SLICESIS DELETING ELMENTS:\nLET'S DELETE SOME ELEMENT >>")
	fmt.Println()

	var ind int
	fmt.Println("At what index do you wanna delete an element?")
	fmt.Scan(&ind)

	fmt.Printf("Your slice before deleting element %v at index %v:\n %v\n", mySlice[ind], ind, mySlice)

	//deleting an element at index specified by the user
	mySlice = append(mySlice[:ind], mySlice[ind+1:]...)
	fmt.Println("Your slice after deleting an element: \n ", mySlice)


