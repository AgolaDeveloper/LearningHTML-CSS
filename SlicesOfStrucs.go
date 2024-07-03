package main

import "fmt"

type Students struct {
	firstName    string
	lastName     string
	class        string
	classTeacher string
	subjects     []string
}

 
	var firstName string
	var lastName string
	var class string
	var classTeacher string
	var subjects []string

	//CREATING A SLICE, myStudents, of type Students (structs), length 5
	//...this is a slice whose elements/values will be objects of type Studnets.
	myStudents := make([]Students, 0)

	//adding elements into our slice, would repetitive thus, loop is vital/necessary.
	//conditional statements shall help with breaking out of the loop...once some conditions are met...
	//...condition we'll depend on is decision/choice of user to either continue adding students details or exit

	fmt.Println("WELCOME TO OUR STUDENTS ADMISSION FORM")

	i := 0
	for {

		var myChoice string
		fmt.Printf("YOU READY TO ADD (Adding a Student's Details?)")
		fmt.Println("?\n 1. Yes\n 2. No")
		fmt.Scan(&myChoice)

		if myChoice == "1" || myChoice == "Yes" || myChoice == "yes" || myChoice == "YES" {

			i++

			fmt.Printf("OK. Go Ahead & ")
			fmt.Println("Enter the student's details here:")
			fmt.Println("FIRST NAME:")
			fmt.Scan(&firstName)

			fmt.Println("LAST NAME:")
			fmt.Scan(&lastName)

			fmt.Println("CLASS:")
			fmt.Scan(&class)

			fmt.Println("CLASS TEACHER:")
			fmt.Scan(&classTeacher)

			fmt.Println("SUBJECTS:")
			/*	//So funny, trying to ensure that the local slice subjects is nill b4 we append another studnet's subjects
				//let's see if it'll work ...LOL!clear

				var x int
				//var value1 string
				for x = range subjects {
					subjects[x] = ""

				}*/

			j := 0
			for {
				j++
				var mySubjects string
				fmt.Printf("ENTER SUBJECT %v:\n", j)
				fmt.Scan(&mySubjects)

				subjects = append(subjects, mySubjects)

				if j == 5 {
					fmt.Printf("Max is %v SUBJECTS", j)
					break
				}

			}

			//var student Students
			//var student =Students{ }
			student := Students{
				firstName:    firstName,
				lastName:     lastName,
				class:        class,
				classTeacher: classTeacher,
				subjects:     subjects,
			}

			myStudents = append(myStudents, student)

			fmt.Printf("You've Entered %v's Details Successfully\n", firstName)

		} else if myChoice == "2" || myChoice == "No" || myChoice == "no" || myChoice == "NO" {

			break
		} else {
			fmt.Println("INVALID CHOICE...RETRY")
			continue

		}

	}
	//trying to ensure that we don print details recorde (empty list) when there is nothing received from teh user

	if i == 0 {
		fmt.Printf("Ooops...\n")

		fmt.Printf(" %v Details Recorded \nBYE\n", i)

	} else {
		fmt.Printf("Details of %v students recorded\n", i)
		fmt.Println(myStudents)

		fmt.Println("#########################")

		var index int
		//var student Students
		for index = range myStudents {

			fmt.Println(myStudents[index])
			//fmt.Println(student)
		}
		fmt.Println("#########################")
	}


