//here we're getting an introductory practice into struct methods/functions

package main

import (
	"fmt"
)

type Student struct {
	name   string
	age    int
	grades map[string]map[string]int
}

//every function of type Students must have a receiver argument...
//...specifying that they are ony functional on instances/objects of a given Struct type

// setter and getter functions for the objetct's name item
func (s *Student) setName(nam string) {
	s.name = nam

}

func (s Student) getName() string {

	//fmt.Println(s.name)
	return s.name
}

// just a Display method; utilizing a getter function (of the same construct though)
func (s Student) welkaam() {
	fmt.Printf("WELCOME to OUR StructMethod PLAYGROUND....HAVE FUN\n")
}

// setter and getter functions for the objetct's age item
func (s *Student) setAge(ag int) {
	s.age = ag
}
func (s Student) getAge() int {
	return s.age
}

// setter and getter functions for the objetct's grades item
func (s *Student) setMarks(grads map[string]int) {

	s.grades = grads
}
func (s Student) getMarks() {

	totalMarks := 0
	fmt.Println("MARKS")
	fmt.Println("\n+++++++++++++++++++\n")

	for key, value := range s.grades {

		totalMarks += value
		fmt.Printf("\n%v: %v\n", key, value)
	}
	fmt.Println("\n+++++++++++++++++++\n")
	fmt.Println("TOTAL MARKS: ", totalMarks)
	fmt.Println("\n+++++++++++++++++++\n")

	//return totalMarks

	//return s.grades
}

// this method returns the total marks scored by the student
/*func (s Student) getTotalMarks() int {

	totalMarks := 0
	count := 0

	for _, value := range s.grades {
		count++
		totalMarks += value
	}
	avarage := totalMarks / count

	fmt.Printf("TOTAL MARKS: %v\n", totalMarks)
	return avarage
}*/

// this struct method helps with getting the students grade
func (s Student) getGrade() {
	tot := 0
	for _, val := range s.grades {
		tot += val
	}

	studentAvarage := tot / len(s.grades)
	fmt.Println("GRADE: ")

	if studentAvarage >= 90 {
		//return ("A ")
		fmt.Printf("A (%v)\n", studentAvarage)
	} else if studentAvarage >= 80 && studentAvarage < 90 {
		//return ("B")
		fmt.Printf("B (%v)\n", studentAvarage)

	} else if studentAvarage >= 60 && studentAvarage < 80 {
		//return ("C")
		fmt.Printf("C (%v)\n", studentAvarage)

	} else if studentAvarage >= 40 && studentAvarage < 60 {
		//return ("D")
		fmt.Printf("D (%v)\n", studentAvarage)

	} else {
		//return ("E")
		fmt.Printf("E (%v)\n", studentAvarage)

	}
}

func main() {

	//declaring our custom variable,studen, of type Student struct
	studen := Student{} //is this nil or empty struct
	//and where are they applicable?

	studen.welkaam()

	//declaring local variables that'll store values of our structs' items before they're set
	var name string
	var age int
	grades := make(map[string]map[string]int)
	gradesValue := make(map[string]int)

	var numOfGrades int

	//we're setting the struct instance's items using our already-defined struct methods

	//Prompting user input on name
	//...and, setting it to the struct name item
	fmt.Println("ENTER YOUR NAME: ")
	fmt.Scanln(&name)
	studen.setName(name)

	//Prompting user input on age
	//...and, setting it to the struct age item

	fmt.Println("ENTER YOUR AGE: ")
	fmt.Scan(&age)
	studen.setAge(age)

	//Prompting user input on grades
	//...and, setting it to the struct grades item

	//here, we'll enter Subjects as keys and respective Grades as values
	fmt.Println("HOW MANY SUBJECTS DO YOU WANNA RECORD THEIR MARKS?: ")
	fmt.Scan(&numOfGrades)

	//then we use a loop to ensure that only the specified number of subjects/grades are entered/mapped
	var subject string
	var marks int
	var code string

	for i := 0; i < numOfGrades; i++ {
		fmt.Printf("\nEnter subject CODE: \n")
		fmt.Scan(&code)

		fmt.Printf("\nEnter Subject for code %v: \n")
		fmt.Scan(&subject)
		fmt.Printf("\nEnter MARKS for SUBJECT %v: \n", subject)
		fmt.Scan(&marks)
		gradesValue[subject] = marks

		//now we map the subject with its respective grade for every iteration
		//by calling setGrades method;since it's defined to do that
		//studen.setGrades(subject, grade)
		grades[code] = gradesValue
	}

	fmt.Println(grades)
	studen.setMarks(grades)

	//...our map is no longer empty
	//and we can now use its setter method to set it to the construct's grades item, of type map

	fmt.Println("############################")
	fmt.Println()
	//GETTERS' SECTION
	//we use getter functions to get display on every information/item for the student

	fmt.Printf("NAME: %v\n", studen.getName())

	fmt.Printf("AGE: %v\n", studen.getAge())

	studen.getMarks()
	studen.getGrade()

	//fmt.Printf("\n\n GRADE: %v \n", y)

}
