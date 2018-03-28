// Accessing the internet finding news sources.
// Instantiate package.
package main 

// Import libraries.
import ("fmt"
		"net/http"
		"io/ioutil")

// Main function.
func main() {
	// Request.
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println(string_body)
	resp.Body.Close()
	
}