// Define package.
package main

// Import mulitiple libraries. When a slash is used it's a subpackage.
import ("fmt"
		"math"
		"math/rand")


// Define function.
/* In go if the first letter is capitalised, that function will be
exported by Go. If you don't capitalise the letter it's considered to
be an internal function. If you're using a package, all of the functions
in the package will always start with a capital letter. 

In go your main function is considered your main loop. It will always run
and does not need to be called like in Python. 

This does not apply to functions that are not called "main". for example,
func foo() would not automatically run, it needs to be called. */
func foo() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

// The main function is passed foo() as it doesn't automatically compile.
/*func main() {
	foo()
}*/

// Random numbers in Go. The output is always the same number unless the seed is changed.
func main() {
	fmt.Println("A number from 1 to 100", rand.Intn(100))
}

/* If you want to know more about a particular package in Go you can use
godoc packagename functionname in the commandline. */