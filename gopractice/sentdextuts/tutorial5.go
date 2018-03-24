// Simple web application.
package  main 

// Import libraries.
import ("fmt"
		"net/http")


// Index handler function.
func index_handler(w http.ResponseWriter, r *http.Request) {
	// w = writer, puts information onto the page, r = request, reads through http request (from memory).
	fmt.Fprintf(w, "Whoa, Go is neat!") // Formats whatever you specified, based on the writer and outputs what you ask.
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	// w = writer, puts information onto the page, r = request, reads through http request (from memory).
	fmt.Fprintf(w, "This is made in Go.") // Formats whatever you specified, based on the writer and outputs what you ask.
}


func main() {
	// Handler.
	http.HandleFunc("/", index_handler) // Home page.
	// To add a new page.
	http.HandleFunc("/about/", about_handler)
	// Server.
	http.ListenAndServe(":8000", nil) // Port, and equivalent of pythons None.
}