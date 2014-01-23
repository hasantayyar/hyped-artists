package main
 
import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)
 
const url = "http://www.last.fm/charts/artists/hyped/place/all"
func main() {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to fetch %s", url)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		log.Printf("Failed to read the response body!")
	}
	fmt.Println(string(body))
}
