package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//func sayGreeting(n string) {
//	fmt.Printf("Good morning %v \n", n)
//}
//
//func sayBye(n string) {
//	fmt.Printf("Goodbye %v \n", n)
//}
//
//func cycleNames(n []string, f func(string)) {
//	for _, name := range n {
//		f(name)
//	}
//}
//
//func circleArea(r float64) float64 {
//	return math.Pi * r * r
//}
//
//func getInitials(n string) (string, string) {
//	s := strings.ToUpper(n)
//	names := strings.Split(s, " ")
//	var initials []string
//	for _, v := range names {
//		initials = append(initials, v[:1])
//	}
//	if len(initials) > 1 {
//		return initials[0], initials[1]
//	}
//
//	return initials[0], "_"
//}
// Function to demonstrate pointers
//func updateName(x *string) {
//	*x = "wedge"
//}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		// Getting an input from user is always string
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("item added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added - ", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you chose to save the bill", b.name)
	default:
		fmt.Println("that was not a valid option.....")
		promptOptions(b)
	}
}

func main() {

	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill)

	//name := "tifa"
	//
	//fmt.Println("Memory address of name is: ", &name)
	//
	//// Storing memory address
	//m := &name
	//fmt.Println("memory address:", m)
	//fmt.Println("value at memory address:", *m)
	//fmt.Println(name)
	//updateName(m)
	//fmt.Println(name)
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

	//greeting := "hello there friends"
	//
	//fmt.Println(strings.Contains(greeting, "hello!"))
	//fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	//fmt.Println(strings.ToUpper(greeting))
	//fmt.Println(strings.Index(greeting, "th"))
	//fmt.Println(strings.Split(greeting, " "))
	//fmt.Println("original string value =", greeting)
	//
	//ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	//
	//// Will alter the slice
	//sort.Ints(ages)
	//fmt.Println(ages)
	//
	//index := sort.SearchInts(ages, 45)
	//fmt.Println(index)

	//x := 0
	//// for - for and while loop
	//for x < 5 {
	//	fmt.Println("value of x is:", x)
	//	x++
	//}
	////for loop
	//for i := 0; i < 5; i++ {
	//	fmt.Println("value of i is:", i)
	//}
	//
	//names := [4]string{"yoshi", "mario", "peach", "bowser"}
	//
	//for i := 0; i < len(names); i++ {
	//	fmt.Println(names[i])
	//}
	//
	////Cycling through a slice
	//for index, name := range names {
	//	fmt.Printf("the position at index %v and value %v \n", index, name)
	//}
	//
	////For not wanting a value use underscore
	//for _, name := range names {
	//	fmt.Printf("value %v \n", name)
	//	// Doesn't alter the original value in the slice
	//	name = "new string"
	//}

	//age := 45
	//
	//fmt.Println(age <= 50)
	//fmt.Println(age >= 50)
	//fmt.Println(age == 45)
	//fmt.Println(age != 50)
	//
	//if age < 30 {
	//	fmt.Println("age is less than 30")
	//} else if age < 40 {
	//	fmt.Println("age is less than 40")
	//} else {
	//	fmt.Println("age is not less than 45")
	//}
	//
	//names := [4]string{"yoshi", "mario", "peach", "bowser"}
	//
	//for index, name := range names {
	//	if index == 1 {
	//		fmt.Println("Continuing at pos", index)
	//		continue
	//	}
	//	if index > 2 {
	//		fmt.Println("breaking at pos", index)
	//		break
	//	}
	//	fmt.Printf("the value at pos %v is %v \n", index, name)
	//}
	//sayGreeting("mario")
	//sayGreeting("luigi")
	//sayBye("mario")
	//
	//cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	//cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)
	//
	//a1 := circleArea(10.5)
	//a2 := circleArea(15)
	//
	//fmt.Println(a1, a2)
	//fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f \n", a1, a2)
	//
	//fn, sn := getInitials("sharvin Harchana")
	//
	//fmt.Println(fn, sn)

	//sayHello("mario")
	//
	//for _, v := range points {
	//	fmt.Println(v)
	//}
	//menu := map[string]float64{
	//	"soup":           4.99,
	//	"pie":            7.99,
	//	"salad":          6.99,
	//	"toffee pudding": 3.55,
	//}
	//
	//fmt.Println(menu)
	//fmt.Println(menu["pie"])
	//
	////Looping maps
	//for k, v := range menu {
	//	fmt.Println(k, "-", v)
	//}
	//
	////ints as key type
	//phonebook := map[int]string{
	//	267584967: "mario",
	//	984759373: "luigi",
	//	845775485: "peach",
	//}
	//
	//fmt.Println(phonebook)
	//fmt.Println(phonebook[267584967])
	//
	//phonebook[984759373] = "bowser"
	//
	//fmt.Println(phonebook)

	//Bill

}
