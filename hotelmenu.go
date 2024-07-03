// We wanna create a Hotel menu application
//there'd be an ADMINISTRATIVE SECTION
//      1. Where admin shallbe updating themenu on available dishes
//      2. and deleting the dishes printed on the menu, but somehow out of stock in the kitchen/pantry

package main

import "fmt"

 	//ADMINISTRATIVE SECTION
	/*FEATURES

	Admin/Chief Chef UPDATING MENU
			A) adding new DISHES to the MENU
			B) Removing Dishes on MENU that are temporarily UNAVAILABLE
	*/

	//A SECTION FOR OUR CHEF TO ADD NEW DISHES INTO THE MENU

	const chefPass = "mainChef01"
	var user string
	menu := map[string]int{
		"Ugali":    100,
		"BeefStew": 200,
		"Rice":     150,
	}

	for {
		fmt.Println("ARE YOU OUR MAIN CHEF or CUSTOMER?:")
		fmt.Println(" 1. CHEF\n 2. CUSTOMER")
		fmt.Scan(&user)
		passError := 0

		if user == "1" || user == "CHEF" {
			var pass string
			fmt.Println("USERNAME: Chef@main")

			passError++

			fmt.Println("ENTER Password:")
			fmt.Scan(&pass)

			//if you enter wrong password 3 times,we log you out...from CHEF PANEL
			if passError == 3 {
				fmt.Println("Maximum Attempts reached")
				fmt.Println("Sure You're our Chef? CounterCheck")
				fmt.Println("*******************************************************************")
				fmt.Println("AUTOMATICALLY LOGGED OU TOF THE HOTEL PORTAL; RETRY AFTER 2 seconds")

				break

			} else if pass == chefPass {
				fmt.Println("WELCOME CHEF...")
				fmt.Println("Here's what we have in the menu so far:")

				for key, value := range menu {
					fmt.Println(key, ": ", value)
				}

				for {
					////what do you wanna do?
					var chefAct string
					fmt.Println("$$+++++++++++++++++++++$$")
					fmt.Printf(" 1. ADD Item\n 2. DELETE Item: \n")
					fmt.Scan(&chefAct)

					if chefAct == "1" {
						var addItem int
						var itemName string
						var itemPrice int

						fmt.Println("HOW MANY items do you wanna ADD?")
						fmt.Scan(&addItem)

						for i := 0; i < addItem; i++ {
							fmt.Println("Enter Item Name: >> ")
							fmt.Scan(&itemName)

							fmt.Println("\nAdd Item Price: >> ")
							fmt.Scan(&itemPrice)

							menu[itemName] = itemPrice
						}

						fmt.Println("$$+++++++++++++++++++++$$")
						fmt.Println("UPDATED MENU:")
						for key, value := range menu {
							fmt.Println(key, ": ", value)

						}

						fmt.Println("$$+++++++++++++++++++++$$")
						break

					} else if chefAct == "2" {
						var removeItem int
						var itemNam string
						//var itemPrice string

						for {

							fmt.Println("$$+++++++++++++++++++++$$")
							fmt.Println("MENU:")

							for key, value := range menu {
								fmt.Println(key, ": ", value)
							}
							fmt.Println("$$+++++++++++++++++++++$$")

							fmt.Println("HOW MANY items do you wanna DELETE?")
							fmt.Scan(&removeItem)

							if removeItem < len(menu) {

								for i := 0; i < removeItem; i++ {

									fmt.Println("Enter Item Name: >> ")
									fmt.Scan(&itemNam)

									//First checking if Item2Delete is available in the menu
									//can't delete sth that's unavailable

									_, ok = menu[itemNam]
									if ok == false {
										fmt.Printf("No item named %v in the HOTEL menu!\n Kindly Re-", itemNam)

									} else {
										delete(menu, itemNam)

									}

								}

								break

							} else {
								fmt.Println("INVALID CHOICE. Items to DELETE, more than ITEMS ON HOTEL MENU...")
								fmt.Printf("RECONFIRM: ")

							}
						}

					} else {

						fmt.Println("INVALID CHOICE: Retry...")

					}
				}

			} else {

				if passError == 2 {
					fmt.Println("Ooops... \nWrong Pass... Retry For the Last Attempt")
					fmt.Printf("\n")
					continue
				}

				fmt.Println("Ooops... \nWrong Pass... Retry")
				fmt.Printf("Re-")

			}

			break

		} else if user == "2" || user == "CUSTOMER" {

			//If user is Customer:
			//we first show our UPDATED HOTEL MENU...

			fmt.Println("$$+++++++++++++++++++++$$")
			fmt.Println("HERE'S THE HOTEL MENU:")
			for key, value := range menu {
				fmt.Println(key, ": ", value)

			}

			fmt.Println("$$+++++++++++++++++++++$$")

			//...then ask them to make an order
			//thereafter we'll keep billing every dish they order

			bill := 0
			fmt.Printf("What Would You Like to Eat?\n Make your Order: \n")

			for {
			cont:
				var custOrder string

				fmt.Scan(&custOrder)

				//then check whether whatever customer's ordering is available/updated on the HOTEL MENU

				//we use temp, ok to confirm this
				_, ok := menu[custOrder]

				if ok == false {
					fmt.Printf("Sorry: %v Unavailable on our Menu ", custOrder)

					fmt.Printf("Ensure Everything you are odering is on this Menu: \n")
					for key, value := range menu {
						fmt.Println(key, ": ", value)
					}

					fmt.Printf("\n Re-Order:\n >> ")

					continue

				} else {
					bill += menu[custOrder]
					var more string
					fmt.Println("Anything Else: 1 to Order more, 2 to exit")
					fmt.Scan(&more)

					if more == "1" {
						for key, value := range menu {
							fmt.Println(key, ": ", value)
						}

						fmt.Println("Make Another Order: ")
						goto cont
						//continue
					} else {
						fmt.Printf("Total Bill: %v", bill)
						break
					}

				}

			}

			break

		} else {
			fmt.Printf("INVALID CHOICE:\n Retry...")

		}
	}

