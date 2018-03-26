// More web basics.
package main 

// Imports.
import ("fmt"
		"net/http")


// Index handler.
/*
Can pass tags just by writing them into the fprint is we want.
*/
func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey There!</h1>")
}
// Main function.
func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}