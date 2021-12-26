package main

import (
	"fmt"
	"reflect"
	"strconv"
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
	name := "Andrew Bell"
	course := "Getting started with Kubernetes"
	module := "4" // set as string
	clip := 2
	//courseComplete := false

	//show how variables are initialised to empty string or 0 (string or int)
	fmt.Println("Name an course are set to", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")

	//using reflect.typeof to show variable type
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))

	//test conversions with str & int
	//total := module + clip
	//fmt.Println("Module plus clip equals", total)

	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		fmt.Println("Module plus clip equals", total)

	}

	//using pointers
	// &<var name> to get pointer memory address
	// Adding a * before a pointer variable returns the value stored in the variable being ported to
	fmt.Println("Memory address of *course* variable is", &course)

	var ptr *string = &course
	fmt.Println("Pointing course variable at address,", ptr, "which holds this value,", *ptr)

}
