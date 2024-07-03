//here is our HR's page that allows him to add and remove the HOTEL EMPLOYEES

package main

import "fmt"

//this is the function to validate our HR's identity
func welcome(hrpas string) {
	const HR_PASSWORD = "Hotel@HR1"
	//count := 0

	//fmt.Println("Login with either PASSWORD or USERNAME")

	//retry:

	if hrpas == HR_PASSWORD {
		fmt.Println("Welcome HR Manager")
		fmt.Println("YOUR RESPONSIBILITIES...:::: \n 1. ADDING NEW EMPLOYEES TO THE HOTEL SYSTEM\n")
		fmt.Println("2. REMOVING/DELETING FIRED EMPLOYEES OUTTA THE SYSTEM ")

	} else {

		fmt.Println("Oops...!\nWrong Password... \n You'll be redirected in 2 seconds to Retry>> ")
		return

		//count++
		//goto retry
	}

}

//ANOTHER Function
//...this function is for adding NEWLY RECRUITED EMPLOYEES...
//strictly meant for STRUCTS

func addEmployees(add int) {

	if add == 1 {

		//creating our own complex/composite newEmployee data type (-STRUCT)
		//every field we need our data type to have is all listed/created innit with their respective types.
		type newEmployee struct {
			firstName   string
			lastName    string
			age         int
			staffNumber string
			//addElement()
			mySlice []int
		}

		func(add *newEmployee) {
			employee.mySlice = append(employee.mySlice, add)
		}

		var fName string
		var lName string
		var ag int
		var sNumber string

		//getting/reading our struct's items/fields from user/HR
		fmt.Println("ENTER Employee's First Name: \n>> ")
		fmt.Scan(&fName)

		fmt.Println("ENTER Employee's Last Name: \n>> ")
		fmt.Scan(&lName)

		fmt.Println("ENTER Employee's Age: \n>> ")
		fmt.Scan(&ag)

		fmt.Println("ENTER Employee's Staff Number: \n>> ")
		fmt.Scan(&sNumber)

		//creating a variable,employee, of type newEmployee... to help us with assigning data values to the STRUCT's items.

		var employee = newEmployee{
			firstName:   fName,
			lastName:    lName,
			age:         ag,
			staffNumber: sNumber,
		}
		employees := make([]newEmployee, 0)
		employees = append(employees, employee)

		fmt.Println("Here are the details you just Entered for the employee:", employee)
		return
	} else if add == 2 {
		fmt.Println("You'll be directed inna few;to DELETE AN EMPLOYEE from the HOTEL DB")
	} else {
		fmt.Printf("Ooops!... In valid Choice...\n Retry\n You'll be redirected shortly")
	}

}


	//However,we must check/validate the HR's pass
	var role int
	var hrPass string
	fmt.Printf("ENTER PASSWORD TO LOG IN:\n >> ")
	fmt.Scan(&hrPass)

	welcome(hrPass)

	fmt.Printf("Enter 1 to ADD and, \n 2 to REMOVE Any Employee(s)")
	fmt.Scan(&role)

	addEmployees(role)


