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
	}
}