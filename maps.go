package main

import "fmt"

//WE WANNA PLAY with MAPS


	//let's create a MENU MAP

	menu := map[string]int{
		"ugaliMix":    70,
		"matumbo":     400,
		"rice":        100,
		"pilau Njeri": 350,
	}

	//let's print out what we have in our menu

	fmt.Println(menu)
	//here we are printing out the price of Matumbo

	var customerName string
	fmt.Printf("Welcome Customer to RasTene Dishes\n What's your First name?")
	fmt.Scan(&customerName)

	//Let's give the customer option to select what they need from the menu
	//...and print them the price

	//var customerChoice string
	fmt.Println("Here's our menu\n Kindly Choose:")

	fmt.Println("######################")
	fmt.Println(" 1. UgaliMix\n2. Matumbo\n 3. Rice\n 4. Pilau Njeri")
	/*fmt.Println("######################")

	fmt.Printf(" So, What would %v like to eat for lunch?", customerName)
	fmt.Scan(&customerChoice)

	switch customerChoice {
	case "1":
		fmt.Printf("ugaliMix is Ksh. %v", menu["ugaliMix"])
	case "2":
		fmt.Printf("matumbo is Ksh. %v", menu["matumbo"])
	case "3":
		fmt.Printf("Rice is Ksh. %v", menu["rice"])
	case "4":
		fmt.Printf("Pilau Njeri is Ksh. %v", menu["pilau Njeri"])
	default:
		fmt.Printf("INVALID CHOICE, MOTHERFUCKER!... Go enjoy your Obwolwo anyways")

	}*/

	//Let's loop through the menu and print every item there

	for i, v := range menu {
		fmt.Println(i, " is Ksh. ", v)

	}
	//for k := 0; k < len(menu); k++ {
	//fmt.Println(menu["k"])
	//}

