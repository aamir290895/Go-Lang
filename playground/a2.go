package main


import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"net/http"
    "log"
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


func main(){
   
	rss := ReadXml()

	fmt.Println(rss.Channel.Title)
	fmt.Println(rss.Channel.Description)


}

