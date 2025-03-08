package main

import "fmt"

func main() {

	/*
	age := 19

	if age < 21 {
		fmt.Println("You can't drink")
	} else if age == 21 {
		fmt.Println("At least another year for drinking")
	} else {
		fmt.Println("You can drink")
	}
	*/


	/*
	age := 23
	sex := "male"

	if age >= 21 && sex == "male" {
		fmt.Println("You can get married")
	} else {
		fmt.Println("You can't get married")
	}
	*/

	/*
	age := 23
	sex := "male"

	if age >= 21 || sex == "male" {
		fmt.Println("You can get married")
	} else {
		fmt.Println("You can't get married")
	}
	*/


	/*
	isPretty := false

	if !isPretty {
		fmt.Println("You are Pretty")
	} else {
		fmt.Println("You are not Pretty")
	}
	*/


	a := 3

	switch a {
	case 1:
		fmt.Println("A is 1")
	case 2, 3:	
		fmt.Println("A is 2 or 3")
	
	default:
		fmt.Println("A is not 1, 2 or 3")
	}	



}

// You have to put conditions in the same line where the previous condition closing bracket ends in go.