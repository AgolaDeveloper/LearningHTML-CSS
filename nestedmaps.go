package main

import "fmt"

type NestedMap struct {
	firstName string
	lastName  string
	grades    map[string]map[string]int
}

// setter methods SECTIONS
func (n *NestedMap) setFirstName(firstName string) {
	n.firstName = firstName

}

func (n NestedMap) getFirsName() string {
	return n.firstName
}
func (n *NestedMap) setLastName(lastName string) {
	n.lastName = lastName

}

func (n NestedMap) getLastName() string {
	return n.lastName
}

func (n *NestedMap) setGrades(grads map[string]map[string]int) {
	n.grades = grads
}
func (n NestedMap) getGrades() map[string]map[string]int {
	return n.grades
}

func main() {
	//creating/declaring an empty instance of our struct, NestedMap
	//var student NestedMap
	student := NestedMap{}

	//declaring local variables for holding values to the struct's items
	var firstName string
	var lastName string

	//creating our nested map; both parent and child maps
	grades1 := make(map[string]map[string]int)
	grades2 := make(map[string]int)

	//creating local variables for storing/holding key-value pairs of our maps
	var key string
	//var value string
	var k string
	var val int

	//prompting for first and last names
	fmt.Println("ENTER FIRST NAME OF STUDENT: ")
	fmt.Scan(&firstName)

	fmt.Println("ENTER LAST NAME OF STUDENT: ")
	fmt.Scan(&lastName)

	//we wanna allow recording for upto 3 subjects max

	for i := 0; i < 3; i++ {
		//we first enter the key of our outter map
		fmt.Println("ENTER COURSE CODE: ")
		fmt.Scan(&key)
		for j := 0; j < 1; j++ {

			//we're moving to its value; our inner map
			//we begin with prompting for its, inner, key
			fmt.Println("ENTER COURSE NAME: ")
			fmt.Scan(&k)

			//then, value of the inner map

			fmt.Println("ENTER COURSE SCORE: ")
			fmt.Scan(&val)

			//what follows is assigning inner value to inner key
			grades2[k] = val
		}
		//remember the entire inner map is value to key of our outter map ...
		//...now we assign the entire inner map to our outer map's key, as its respective value
		//...this happens for every iteration in this loop-block
		grades1[key] = grades2
	}

	//WE ARE SETTING OUR ITEMS
	student.setFirstName(firstName)
	student.setLastName(lastName)

	//... and pass our locally-populated map to the struct's setter method
	student.setGrades(grades1)

	//WE ARE GETTING BACK OUR OBJECT'S VALUES/ITEMS
	fmt.Println(student.getFirsName())
	fmt.Println(student.getLastName())
	fmt.Println(student.getGrades())
}
