package main

import (
	"fmt"
	"sort"
	"strings"
)

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
	//var ageOne int = 20
	//var ageTwo = 30
	//ageThree := 40

	//fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	//var numOne int8 = 25
	//var numTwo int8 = -128
	////Cant have a negative number
	//var numThree uint8 = 255

	//age := 35
	//name := "shaun"
	//
	//fmt.Print("hello, ")
	//fmt.Print("world! \n")
	//fmt.Print("new line \n")
	//
	//fmt.Println("hello ninjas!")
	//fmt.Println("goodbye ninjas!")
	//fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted strings)
	// v --> variable specifier
	// q --> number specifier
	// T --> for type specifier
	// f --> floating points
	//fmt.Printf("my age is %v and my name is %v \n", age, name)
	//fmt.Printf("my age is %q and my name is %q \n", age, name)
	//fmt.Printf("age is of type %T \n", age)
	//fmt.Printf("you scored %0.1f points! \n", 225.55)

	// Sprintf (save formatted strings)
	//var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	//fmt.Println("the saved string is:", str)

	// syntax --> var variable_name [length]data_type = [length]data_type{values}
	//var ages [3]int = [3]int{20,25,30}
	//short hand
	//var ages = [3]int{20, 25, 30}
	//
	//names := [4]string{"yoshi", "mario", "peach", "bowser"}
	//names[1] = "luigi"
	//fmt.Println(ages, len(ages))
	//fmt.Println(names, len(names))
	//
	//// slices (use arrays under the hood), can append items
	//var scores = []int{100, 50, 60}
	//scores[2] = 25
	//scores = append(scores, 85)
	//
	//fmt.Println(scores, len(scores))
	//
	//// slice
	//// Does not include 3
	//rangeOne := names[1:3]
	//rangeTwo := names[2:]
	//rangeThree := names[:3]
	//
	//fmt.Println(rangeOne, rangeTwo, rangeThree)
	//
	//rangeOne = append(rangeOne, "koopa")
	//fmt.Println(rangeOne)

	greeting := "hello there friends"

	fmt.Println(strings.Contains(greeting, "hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "th"))
	fmt.Println(strings.Split(greeting, " "))
	fmt.Println("original string value =", greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	// Will alter the slice
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 45)
	fmt.Println(index)
}
