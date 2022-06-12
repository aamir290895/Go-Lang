package news

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "encoding/json"
)


type RSS struct {

	Channel struct {

		Title string  `xml:"title"`
		Description string `xml:"description"`
		Link string `xml:"link"`

		Item []struct {
			Title string        `xml:"title"`
			Description string  `xml:"description"`
			GuId string         `xml:"guid"`
			Link string         `xml:"link"`
			PublishDate string  `xml:"pubDate"`

		}  `xml:"item"`
	} `xml:"channel"`


} 

func ReadXml() *RSS{

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
