// here we practice goroutines
// ... and channels

// Imma define a goroutine that writes data into our empty channel
// and main function gotta read form the channel
package main

import "fmt"

//defining a goroutine responsible for writing data into the channel'

func compute(a, b int, chan1 chan int) {
	//mehtod first do some computation
	sum := a + b

	//then writes/sends data (the computation) to our channel
	chan1 <- sum

}

//this method reads data from the channel
func readChannel(ch chan int) int {
	//it must have a variable to store data from the channel
	readChan := <-ch
	return readChan
}

func main() {
	//here we first create an empty channel
	channel1 := make(chan int)
	//then declare local variables for computation
	num1 := 10
	num2 := num1 * 2

	//now we pass data and channel to a goroutine
	//it's this goroutine that's responsible for writing these data to our channel

	go compute(num1, num2, channel1)

	//what'll receive/read the data from our channel?
	//we can have a function for that
	//or main function can do it...however,the main must
	//....have a variable that'd be ready to store the data read from this channel

	//readChannel := <-channel1 //this reads data from our channel

	//fmt.Println("Data read from our channel, by main function: ", readChannel)

	//let's experiment with a local function as receptor of our channel's data
	//go readChannel(channel1)

	fmt.Println("Data read from our channel, by a local method", readChannel(channel1))

}
