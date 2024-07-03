package main

import "fmt"

type Studentss struct {
	studentName    string
	studentAdmNo   int
	officialParent string
	class          string
	classteacher   string
}

func structBasic() {
	fmt.Println("Struct Basics")
	var stdntName string
	var stdntAdm int
	var parent string
	var clss string
	var clssTchr string

	var student1 = Studentss{
		studentName:    stdntName,
		studentAdmNo:   stdntAdm,
		officialParent: parent,
		class:          clss,
		classteacher:   clssTchr,
	}
	fmt.Printf("Welcome, kindly Enter the following details (if you're a student)\n Your Name\n Your Admission Number\n Your Parent's Name\n Your Class\n Your Class Teacher\n")
	fmt.Println("Name?:")
	fmt.Scan(&stdntName)

	fmt.Println("Admission Number?:")
	fmt.Scan(&stdntAdm)

	fmt.Println("Parent's Name?:")
	fmt.Scan(&parent)

	fmt.Println("Your Class?:")
	fmt.Scan(&clss)

	fmt.Println("Class Teacher?:")
	fmt.Scan(&clssTchr)

	for {
		var yourChoice string
		fmt.Printf("%v would you like to go ahead and see your Details? Enter: \n 1. No \n 2. Yes\n ", student1.studentName)
		fmt.Scan(&yourChoice)
		if yourChoice == "2" || yourChoice == "Yes" || yourChoice == "yes" || yourChoice == "YES" {
			fmt.Printf("Name: %v", student1.studentName)
			fmt.Printf("Admission No.: %v", student1.studentAdmNo)
			fmt.Printf("Parent/Guardian: %v", student1.officialParent)
			fmt.Printf("Your Class: %v", student1.class)
			fmt.Printf("Your ClassTr: %v", student1.classteacher)

			break

		} else if yourChoice == "1" || yourChoice == "No" || yourChoice == "no" || yourChoice == "NO" {
			fmt.Printf("You just Exited %v: BYE...", student1.studentName)

			break
		} else {
			fmt.Printf("Acha Upuuzi %v: YOU ENTERED INVALID CHOICE\n Retry", student1.studentName)

			continue
		}

	}

}
