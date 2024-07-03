package main

import "fmt"


	//printing 5*5 asteric pattern
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("*")
			fmt.Printf("")
		}
		fmt.Println("")

	}
	fmt.Printf("\n####################\n SECOND PATTERN\n")

	//printing 5 by 5 trriangle
	for i := 0; i < 5; i++ {

		for j := 4; j >= i; j-- {
			//if i >= j {
			fmt.Printf("*")
			fmt.Printf(" ")
			//} else {
			//continue
			//}
		}
		fmt.Println("")
	}

	/*fmt.Println("#################")

	for i := 0; i < 5; i++ {

		for j := 0; j < 5; j++ {
			if i <= j {
				fmt.Printf("*")
				//fmt.Printf(" ")
			} else {
				continue
			}
		}
		fmt.Println("")

	}
	fmt.Println("#################")

	for i := 0; i < 5; i++ {

		for j := 5; j >= 0; j-- {
			if i <= j {
				fmt.Printf("*")
				//fmt.Printf(" ")
			} else {
				continue
			}
		}
		fmt.Println("")

	}*/


