// Function
/*
package main

import "fmt"

func add(num1 int, num2 int){
	sum := num1 + num2

	fmt.Println(sum)
}

func main(){
	a := 10
	b := 20

	add(a,b)
	add(5,7)
}
*/

// Function with Return Value and Type
/*
package main

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2

	return sum
}

func main(){
	a := 5
	b := 19

	sum := add(a,b)

	fmt.Println(sum)
}
*/

// Function with 2 Return

/*
package main

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2

	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	
	mul := num1 * num2

	return sum, mul
}

func main(){
	a := 5
	b := 19

	p, q := getNumbers(a,b)

	fmt.Println(p, q)
}
*/

// Functions without Input or Output and with String Input


package main

import "fmt"


func printSomething(){
	fmt.Println("Education must be free.")
}

func sayHello(name string){
	fmt.Println("Welcome to my Github. Make sure you hire me. Thank You, ", name)
}

func main(){
	printSomething()
	sayHello("Recruiter")
}