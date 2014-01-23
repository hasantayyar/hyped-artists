package main
 
import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"regexp"
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


	re, _ := regexp.Compile(`<span class="rankItem-title">(.*)</span>`)
    titles := re.FindAllStringSubmatch(string(body), -1)



	re2, _ := regexp.Compile(`href="(.*)"(.*)class="rankItem-blockLink"`)
    urls := re2.FindAllStringSubmatch(string(body), -1)

    for index,element := range titles {
    	fmt.Printf("%v\t\t", element[1])
    	fmt.Printf("http://last.fm/%v \n", urls[index][1])
	}
}
