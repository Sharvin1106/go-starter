package main

import "fmt"

func main() {

	// Strings use double quote
	// If you dont use a variable in go its an error
	//var nameOne string = "Mario"
	//// Self inference
	//var nameTwo = "luigi"
	//// Early initialization - will not case error if we are printing it
	//var nameThree string
	//fmt.Println(nameOne, nameTwo, nameThree)
	//
	//nameOne = "peach"
	//nameThree = "bowser"
	//
	//fmt.Println(nameOne, nameTwo, nameThree)
	//
	//// Shorthand for --> var nameFour string = "yoshi"
	//// Can't use the short hand notation outside a function
	//nameFour := "yoshi"
	//
	//fmt.Println(nameFour)
	//

	//ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	//Cant have a negative number
	var numThree uint8 = 255
}
