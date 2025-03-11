package main

import "fmt"

/*
func main(){

	// Print Welcome Message
	fmt.Println("Welcome to the Application")

	// User Input
	var name string
	fmt.Println("Enter your Name: ")
	fmt.Scanln(&name)

	var num1 int
	var num2 int

	fmt.Println("Enter First Number: ")
	fmt.Scanln(&num1)

	fmt.Println("Enter Second Number: ")
	fmt.Scanln(&num2)

	// Display Results
	fmt.Println("Hello! ", name)
	fmt.Println("Total is: ", num1 + num2)

	fmt.Println("Thank you for using the Application")
	fmt.Println("Goodbye!")
}
*/



func printWelcomeMessage(){
	fmt.Println("Welcome to the Application")
}

func getUserName() string {
	var name string
	fmt.Println("Enter your Name: ")
	fmt.Scanln(&name)
	return name
}

func getTwoNumbers() (int, int){
	var num1 int
	var num2 int

	fmt.Println("Enter First Number: ")
	fmt.Scanln(&num1)

	fmt.Println("Enter Second Number: ")
	fmt.Scanln(&num2)

	return num1, num2
}

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func display(name string, sum int) {
	fmt.Println("Hello! ", name)
	fmt.Println("Total is: ", sum)
}

func printGoodbyeMessage() {
	fmt.Println("Thank you for using the Application")
	fmt.Println("Goodbye!")
}



func main(){

	printWelcomeMessage()

	name := getUserName()

	num1 , num2 := getTwoNumbers()

	sum := add(num1, num2)

	display(name, sum)

	printGoodbyeMessage()

}