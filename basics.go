package main // Every Go program must have a package

import "fmt" // import fmt package which is short for Format

func main() {

	fmt.Println("Hello World!") // This is print function works.

	a := 10
	b := "Apurba Halder"


	fmt.Println(a, b)

	var x int = 1903 // var keyword is used to declare a variable
	fmt.Println(x)

	y := "This is a string" // we don't have to declare type of variable
	
	y = "This is a new string" 
	/* if we want to put a new value in an already declared variable we have to remove the colon before = */

	fmt.Println(y)


	const p = 100
	//p = 100 // we can't change the value of constant variable
}

/*
	Data Types
	1. int
	2. float32
	3. string
	4. bool
*/
