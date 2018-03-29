// Parsing xml.
// Instantiate package.
package main 

// Import libraries.
import ("fmt"
		"net/http"
		"io/ioutil"
		"encoding/xml")



// Define structure of xml document.
type SiteMapIndex struct {
	// An array of location type.
	Locations []Location `xml:"sitemap"` // Locations HAS to be capitalised.
}

type Location struct {
	Loc string `xml:"loc"`
}
/*
 [X]type == array // Fixed size.
 	ie [5 5]int // 5 by 5 int array.
 []type == slice 
*/

// String method for transformation.
func (L Location) String() string {
	return fmt.Sprintf(L.Loc)

}
// Main function.
func main() {
	// Request.
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SiteMapIndex
	xml.Unmarshal(bytes, &s) // Unmarshall into memory address of s.
	fmt.Println(s.Locations)
	
}