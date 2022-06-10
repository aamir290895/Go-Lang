package news

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"net/http"
    "log"
    // "encoding/json"
)


type JNews struct {

	Channel struct {

		Title string  `json:"title"`
		Description string `json:"description"`
		Link string `json:"link"`

		Item []struct {
			Title string        `json:"title"`
			Description string  `json:"description"`
			GuId string         `json:"guid"`
			Link string         `json:"link"`
			PublishDate string  `json:"pubDate"`

		}  `json:"item"`
	} `json:"channel"`


} 

func ReadJSON() *RSS{

	resp, err := http.Get("https://timesofindia.indiatimes.com/rssfeedstopstories.cms")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

    rss := &RSS{}
	

	xml.Unmarshal(body, rss)

	if (err != nil){
		fmt.Println("Error in fetching data")
	}
    
	return rss
    
}
