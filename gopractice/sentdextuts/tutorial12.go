// Looping.
package main 

// Imports.
import ("fmt")

// Main function.
func main() {
	i:=0 // Initialise i.
	for i=0; i<10; i++ {
		// While i < 10 add 1 to i.
		fmt.Println(i)
	}
	// Same as above
	i=0
	for i<10 {
		fmt.Println(i)
		i++ // Or i+=1
	}
	/*
	Infinite loop.
	for {
		fmt.Println("hi")
	}
	*/

	// With if statement.
	x:=5
	for {
		fmt.Println("Do stuff", x)
		x+=3
		if x > 25 {
			break // Get out of this specific loop.
		}
	}
	// Another way to do above.
	for x:=5; x<25; x+=3{
		fmt.Println("Do stuff", x)
	}

}