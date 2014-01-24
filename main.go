package main
 
import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"regexp"
	"net/url"

)

type jsonData struct {
    Field1 string
    Field2 string
}

type JsonData struct {
    jsonData
}

// Implement json.Unmarshaller
func (d *JsonData) UnmarshalJSON(b []byte) error {
    return json.Unmarshal(b, &d.jsonData)
}

// Getter
func (d *JsonData) Field1() string {
    return d.jsonData.Field1
}


const lstfm= "http://www.last.fm/charts/artists/hyped/place/all"
func main() {
	fmt.Printf("Getting hyped artist list.\n\n")
	resp, err := http.Get(lstfm)
	if err != nil {
		fmt.Printf("Network error for url :  %s", lstfm)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	re, _ := regexp.Compile(`<span class="rankItem-title">(.*)</span>`)
	titles := re.FindAllStringSubmatch(string(body), -1)

	re2, _ := regexp.Compile(`href="(.*)"(.*)class="rankItem-blockLink"`)
	urls := re2.FindAllStringSubmatch(string(body), -1)

    for index,element := range titles {
    	name := element[1]

    	// get one video from youtube 

    	searchurl := "https://gdata.youtube.com/feeds/api/videos?caption&v=2&max-results=1&format=1&start-index=1&vq="+url.QueryEscape(name)+"&orderby=relevance&alt=json"
		resp, _ := http.Get(searchurl) 
		defer resp.Body.Close()
		searchresult, _ := ioutil.ReadAll(resp.Body)


		var f interface{}
		//  @todo check error
		json.Unmarshal(searchresult, &f);

		// @todo check null interfaces
		m, _ := f.(map[string]interface{})
		feed, _  := m["feed"].(map[string]interface{})
		entries, _  := feed["entry"].([]interface{})
		entry, _  := entries[0].(map[string]interface{})
		links, _  := entry["link"].([]interface {})
		link, _  := links[0].(map[string]interface{})
 
    	fmt.Printf("%v\t\t", name)
    	fmt.Printf("%+v\t", link["href"])

    	fmt.Printf("http://last.fm/%v \n", urls[index][1])
	}
}
