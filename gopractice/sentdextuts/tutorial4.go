// Pointers.
// Define package.
package main

// Import libraries.
import "fmt"

// Main function.
func main() {
	x := 15
	a := &x // & allows us to 'point' to something, which is the memory address.
	fmt.Println(a) // Prints the memory address.
	fmt.Println(*a) // Reads from the memory address.
}