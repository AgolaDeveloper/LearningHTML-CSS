// WE CREATING A STRUCT WITH A 'SLICE OF STRUCTS' AS one of its items
package main

import "fmt"

//creating struct... BANK with employee information
//since our bank has many branches we'd be creating several instances of type BANK...
//...the instances represent every bank branch we got.
//and what details should we have in every Bank branch?...
//. 1. Branch Name
// 2. Branch Manager
//. 3. CustomersInfo

//FIRST, We create customerInfo, a struct that'll store info for every customer that register with the respective branch.

type customerInfo struct {
	firstName      string
	lastName       string
	accountNumber  string
	accountBalance uint
	accountPin     string
}

//METHODS FOR customerInfo STRUCT...
//SETTER Methods

func (c *customerInfo) setFirstName(fNam string) {
	c.firstName = fNam

}
func (c *customerInfo) setLastName(lNam string) {
	c.lastName = lNam
}
func (c *customerInfo) setAccntNumber(accNum string) {
	c.accountNumber = accNum
}
func (c *customerInfo) setAccntBalance(accBal uint) {
	c.accountBalance = accBal
}
func (c *customerInfo) setAccntPin(accPin string) {
	c.accountPin = accPin
}

//GETTER Methods
func (c customerInfo) getFirstName() string {
	return c.firstName

}
func (c *customerInfo) getLastName() string {
	return c.lastName
}
func (c *customerInfo) getAccntNumber() string {
	return c.accountNumber
}
func (c *customerInfo) getAccntBalance() uint {
	return c.accountBalance
}
func (c *customerInfo) getAccntPin() string {
	return c.accountPin
}

//NOW CREATING THE Bank struct, that'll hold every info/detail for our bank [rather bank branches]

type Bank struct {
	branchName    string
	branchManager string
	customers     []customerInfo
}

//METHODS FOR Bank STRUCT
//SETTER Methods...
func (b *Bank) setBranchName(branch string) {
	b.branchName = branch
}
func (b *Bank) setBranchManager(manager string) {
	b.branchName = manager
}
func (b *Bank) setCustomers(customer []customerInfo) {
	b.customers = customer
}

//GETTER Methods...
func (b Bank) getBranchName() string {
	return b.branchName
}
func (b Bank) getBranchManager() string {
	return b.branchName
}
func (b Bank) getCustomers() []customerInfo {

	for index, value := range b.customers {
		fmt.Println(index, ": ", value)
	}
	return b.customers
}

//this method omitts the customers' sensitive info & and gives the basic one
func (b Bank) displayCustomers() {
	for _, value := range b.customers {
		fmt.Printf("\n#############\n %v: \n Account No.: %v\n Balance: %v\n", value.lastName, value.accountNumber, value.accountBalance)
	}
}

//ADDING A FEW METHODS THAT'LL HELP WITH MODIFYING SOME CUSTOMER DETAILS

//THIS helps with changing a customer's account pin
func (c *customerInfo) changePin(accNum, newPin string) {

	//we'll receive 2 arguments from user in charge of changing customer pin
	//... user's account number and new pin
	//... we'll range through the slice of customers in the Bank instance...
	//... comparing their account numbers with the one passed
	//... in case there's a match, we go a head and change the respective object's pin to what's passed

	//accountExist := 0
	//for _, value := range b.customers {

	oldPin := c.getAccntPin()
	x := c.getAccntNumber()

	if x == accNum {

		//accountExist++

		c.setAccntPin(newPin)

		fmt.Println("Hurray!... Pin changed successfully")
		fmt.Println("******************")
		fmt.Printf(" Old Pin: %v \nNew Pin: %v \n", oldPin, c.getAccntNumber())

	} else {
		fmt.Println("Ooops.... Account Doesn't Exist")

	}

}

func main() {
	//first creating an instance for our customers
	customer := customerInfo{}

	//declaring local variables that'll hold customer's items/values
	var firstName string
	var lastName string
	var accountNumber string
	var accountBalance uint
	var accountPin string

	//again, creating an instance of our bank[branches]
	branch := Bank{}

	//declaring localvariables that'll hold bank's items/values
	var branchName string
	var branchManager string
	customers := make([]customerInfo, 0)

	//we're getting started with adding all branches we have in the country, into our database
	for {
		var brnch string
		var addCust string
		fmt.Printf("You wanna add another Branch?\n(y/n):>> \n")
		fmt.Scan(&brnch)
		if brnch == "y" {
			fmt.Println("Enter Name of the Branch")
			fmt.Scan(&branchName)

			fmt.Println("Enter Name of the Branch Manager")
			fmt.Scan(&branchManager)

			//setting branch details, customers exclusive
			branch.setBranchManager(branchManager)
			branch.setBranchName(branchName)

		anyCus:
			fmt.Printf("Any Customers to register? \n (y/n):\n")
			fmt.Scan(&addCust)
			if addCust == "y" {
				goto registerCustomers
			} else if addCust == "n" {
				break
			} else {
				goto anyCus
			}

		} else if brnch == "n" {
			break

		} else {
			fmt.Println("INVALID CHOICE!...Retry>>")
		}
	}

registerCustomers:
	//here's where we register customers
	for {

		fmt.Println("Kindly Register customer")

		var regCus int
		fmt.Println("Enter 1 to Register a Customer")
		fmt.Scan(&regCus)

		if regCus == 1 {
			//Enter Customers' details
			fmt.Println("Enter Customer First Name")
			fmt.Scan(&firstName)

			fmt.Println("Enter Customer Last Name")
			fmt.Scan(&lastName)

			fmt.Println("Enter Customer's Account Number")
			fmt.Scan(&accountNumber)

			fmt.Println("Enter Customer Account Balance")
			fmt.Scan(&accountBalance)

			fmt.Println("Enter Customer PIN")
			fmt.Scan(&accountPin)

			//setting the customer details
			customer.setAccntBalance(accountBalance)
			customer.setAccntNumber(accountNumber)
			customer.setAccntPin(accountPin)
			customer.setFirstName(firstName)
			customer.setLastName(lastName)

			//we add, customer details, instance of customerInfo to the Bank
			//...bank has slice of structs, as an item, to hold/store/keep all objects of customerInfo (details of registered Customers)
			//...so we append our object customer to the Bank's local slice-customers

			customers = append(customers, customer)

			//then we set the Bank's slice-method
			branch.setCustomers(customers)

		} else {
			fmt.Printf("HERE ARE INFO AT YOUR %v BRANCH: \n", branch.getBranchName())
			fmt.Println(branch.getBranchManager())
			branch.getCustomers()

			break

		}
	}

	//printing the information wealready have at our Bank branch
	fmt.Println("DETAILS OF OUR BANK BRANCH:")

	fmt.Printf("\n BRANCH NAME: ", branch.getBranchName())

	fmt.Printf("\n BRACH Manager: ", branch.getBranchManager())

	fmt.Printf("\n BRACH CUSTOMERS: \n")
	//we're only printing last names and account Numbers of the customers available
	branch.displayCustomers()

	fmt.Println("$$$$$$$$$$$$")
	fmt.Println(" Let's Experiment with changing a user's Account pin")

	var accNum string
	fmt.Println("Enter Account Number: ")
	fmt.Scan(&accNum)

	var newPin string
	fmt.Println("What's Your New Pin?: ")
	fmt.Scan(&newPin)

	customer.changePin(accNum, newPin)

	branch.getCustomers()

}
