// we're advancing concepts of goroutines
// our channels shall hold structs
// we'll have 3 gourotines interconnected with 2 channels
package main

import "fmt"

//defining our composite type; Employee
type Employee struct {
	name  string
	bonus map[string]int
}

//defining the struct's setters
func (e *Employee) setName(name string) {
	e.name = name
}
func (e *Employee) setBonus(bonus map[string]int) {
	e.bonus = bonus
}

//defining the struct's GETTER methods

func (e Employee) getName() string {
	return e.name
}
func (e Employee) getBonus() map[string]int {
	return e.bonus
}

//first goroutine writes employee details into the channel, chennel 1
//it takes 2 parameters: object/instance of Employee struct and channel one

func (e Employee) empG(ch1 chan Employee) {
	//employeeNam := e.name
	//employeeBon := e.bonus
	ch1 <- e
}

//this method reads data from channel one
//prints last year's bonus for employee; jan-dec
//plus writes data to the 2nd channel
//thus, take two parameters; for reading from 1st and writing to 2nd data

func printBonus(ch1 chan Employee, ch2 chan Employee) {
	//first reads data from channel1
	emp1 := <-ch1

	//print the bonus
	fmt.Println("################################")
	fmt.Printf("Here are Bonus for %v: Jan-Dec\n", emp1.name)
	for key, value := range emp1.bonus {
		fmt.Printf("%v: %v\n", key, value)
	}
	fmt.Println("################################")

	//it then writes the data to another channel, channel2
	ch2 <- emp1
}

//the last method that receives one parameter, chan2
//it reads data from chan2 and...
//finds avarage bonus salary for the whole year
//and prints bonuses that are only >5000
func avaBonus(ch2 chan Employee) {
	//first, reads data from channel2
	emp2 := <-ch2

	//uses that data to find average bonus
	tot := 0
	count := 0
	for _, value := range emp2.bonus {
		tot += value
		count++
	}

	averageBonus := tot / count
	fmt.Println("AVARAGE BONUS (for the previous year): ", averageBonus)

	//finds bonuses above 5000
	tempMap := make(map[string]int)

	for key, value := range emp2.bonus {
		if value > 5000 {
			//append the pairs to another temp map
			tempMap[key] = value
		}
	}

	//prints if there's bonuses above 5000

	if len(tempMap) > 0 {
		//then print its values(obviously they are all above 5000)
		for key, value := range tempMap {
			fmt.Printf("%v: %v\n", key, value)
		}
	} else {
		fmt.Printf("NO BONUSES ABOVE 5000 for Employee %v\n", emp2.name)
	}
}

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func main() {
	//defining local instance of Employee and its items
	emp := Employee{}

	var name string
	bonus := make(map[string]int)

	//we read our local values befor setting var emp (of type Employee) with items

	fmt.Println("Enter Employee's Name: ")
	fmt.Scan(&name)

	fmt.Printf("WE ARE ABOUT TO ENTER %'%s Bonus Salary (JAN-DEC)\n", name)

	var bon int
	fmt.Println("Enter Bonus for JAN: ")
	for i := 0; i < 12; i++ {
		fmt.Println(" ")

		fmt.Println("ENTER HERE: \n>>>  \n")
		fmt.Scan(&bon)

		fmt.Println(" ")
		//switch mapping bonus entered with respective months
		//...in relation to the matching index number
		switch i {
		case 0:
			bonus["JANUARY"] = bon

			//letting users know month they just entered
			fmt.Println("You Just Entered Bonus for JANUARY")
			fmt.Println("Enter Bonus for FEB: ")

		case 1:
			bonus["FEBRUARY"] = bon

			//letting users know month they just entered
			fmt.Println("You Just Entered Bonus for FEB")
			fmt.Println("Enter Bonus for MARCH: ")

		case 2:
			bonus["MARCH"] = bon

			//letting users know month they just entered
			fmt.Println("You Just Entered Bonus for MARCH")
			fmt.Println("Enter Bonus for APRIL: ")

		case 3:
			bonus["APRIL"] = bon

			//letting users know month they just entered
			fmt.Println("You Just Entered Bonus for APRIL")
			fmt.Println("Enter Bonus for MAY: ")

		case 4:
			bonus["MAY"] = bon

			//letting users know month they just entered
			fmt.Println("You Just Entered Bonus for MAY")
			fmt.Println("Enter Bonus for JUNE: ")

		case 5:
			bonus["JUNE"] = bon

			//letting users know month they just entered
			fmt.Println("You Just Entered Bonus for JUNE")
			fmt.Println("Enter Bonus for JULY: ")

		case 6:
			bonus["JULY"] = bon

			//letting users know month they just entered
			fmt.Println("You Just Entered Bonus for JULY")
			fmt.Println("Enter Bonus for AUGUST: ")

		case 7:
			bonus["AUGUST"] = bon

			//letting users know month they just entered
			fmt.Println("You Just Entered Bonus for AUGUST")
			fmt.Println("Enter Bonus for SEPTEMBER: ")

		case 8:
			bonus["SEPTEMBER"] = bon

			//letting users know month they just entered
			fmt.Println("You Just Entered Bonus for SEPTEMBER")
			fmt.Println("Enter Bonus for OCTOBER: ")

		case 9:
			bonus["OCTOBER"] = bon

			//letting users know month they just entered
			fmt.Println("You Just Entered Bonus for OCTOBER")
			fmt.Println("Enter Bonus for NOVEMBER: ")

		case 10:
			bonus["NOVEMBER"] = bon

			//letting users know month they just entered
			fmt.Println("You Just Entered Bonus for NOVEMEBR")
			fmt.Println("Enter Bonus for DEC: ")

		case 11:
			bonus["DECEMBER"] = bon

			//letting users know month they just entered
			fmt.Println("You Just Entered Bonus for DEC")
			fmt.Println("Hurray.....BONUSES FOR ALL MONTHS ENTERED ")
			fmt.Println("+++++++++++++++++++++++++++++++++++++++++++")

		}
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++")

		fmt.Println("BONUSES RECORDED SO FAR: ")
		for key, value := range bonus {
			fmt.Printf("%v: %v\n", key, value)
		}

		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++")

	}
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	//setting the read local values/variables to our emp instance, asitems
	emp.setName(name)
	emp.setBonus(bonus)

	emp.getBonus()
	emp.getName()
	//+++++++++++++++++++++++++++++++++++++++++++++++++
	//let's create our two channels
	//...that our goroutines shall use for communication
	chan1 := make(chan Employee)
	chan2 := make(chan Employee)

	//calling our first goroutine
	//it'll write data to our first channel
	go emp.empG(chan1)

	//calling 2nd goroutine
	//it'll read data from 1st channel, print Bonuses
	//...and write to the 2nd channel

	go printBonus(chan1, chan2)

	//lastly...we're calling third goroutine
	//that'll read data from channel 2 and...
	// print average bonus+bonuses above 5000

	go avaBonus(chan2)
}
