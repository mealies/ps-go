package main

import (
	"fmt"
	"os"
)

// variables declared at package level (outside main)
// must use var
//declaring variables with no data
/*var (
	name   string
	course string
	module int
	clip   int
)
*/

// variables declared at package level (outside main) global in scope
//declaring variables with values
//uses type inference to work out variable type
/*var (
	name   = "Andrew Bell"
	course = "Getting started with Kubernetes"
	module = "4" // set as string
	clip   = 2
)
*/
func main() {

	//within functions, you can declare vars using shorthand
	// := means you dont need to use var keyword
	name := os.Getenv("USER")
	course := "Getting started with Kubernetes"

	fmt.Println("\nHi", name, "your current course is", course)
	//to change the course var in main from the function
	// we need to pass it as a pointer and change the original
	//memory location, or the update func will change a copy only
	updateCourse(&course)
	fmt.Println("You're currently watching", course)

}

//
func updateCourse(course *string) string {
	*course = "Getting started with Docker"
	fmt.Println("Updated course to", *course)
	return *course
}
