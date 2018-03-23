// Types in Go.
// Define package
package main

// Import libraries.
import ("fmt")


// Function to add together 2 variables.
/* In go you have to make sure the types of these variables
are defined before you add them together, something that you
don't need to do in Python.  */
func add(x, y float64) float64 {
	return x + y
}

// Function that returns multiple things. You must specify every return type, even if they're the same.
func multiple(a, b string) (string, string) {
	return a, b

}
// Define a constant value.
const z int = 5

// Main function.
func main() {
	// Define a variable.
	//var num1, num2 float64 = 5.6, 9.5

	// Automatically give variables a type. If you want types to be explicit it is better to use the above method.
	num1, num2 := 5.6, 9.5
	fmt.Println(add(num1, num2))

	w1, w2 := "Hey", "There"
	fmt.Println(multiple(w1, w2))

	// Convert a variable to another type.
	var a int = 62
	var b float64 = float64(a)
}