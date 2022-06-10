package news

import (
	"gorm.io/gorm"
	"net/http"
	"fmt"
	"encoding/json"

)

type Category struct {

	gorm.Model

	CategoryName string        `json:"category" gorm:"primarykey"`

	Articles []Article         `json:"articles" gorm:"foreignkey: GUID"`

}

type Article  struct {
  
	gorm.Model
	GUID string           `json:"guid"`          
	Title string          `json:"title"`
	Description string    `json:"description"`
	Url string            `json:"url"`
	ImageUrl string       `json:"image_url"`
	PostDate string       `json:"posted_date"`
	Tags string           `json:"tags"`
	Author string         `json:"author"`
  
  }


type Handler struct {
    DB *gorm.DB
}

func New(db *gorm.DB) Handler {
    return Handler{db}
}

func(h Handler) CreateNews( w http.ResponseWriter , r *http.Request){

	defer r.Body.Close()
	
    data := ReadXml()

	var category Category

	var article Article
	var articles []Article


    for _,x := range data.Channel.Item {
		article.GUID = x.GuId
		article.Title = x.Title
		article.Description = x.Description
		article.Url=x.Link
		article.PostDate = x.PublishDate
		articles = append(articles , article)


	}


	category = Category {
		CategoryName : data.Channel.Title,
		Articles : articles,
	}



    // body ,err := ioutil.ReadAll(r.Body)
	// if err != nil {
	//  log.Print(err)
	// }
	// // h.DB.AutoMigrate(&category)

    // json.Unmarshal(body ,&categories)


	result := h.DB.Create(&article) ;

	if  result.Error != nil {
        fmt.Println(result.Error)
    }


    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(category)
    
  }

  func(h Handler) GetCategories ( w http.ResponseWriter , r *http.Request){

	var categories []Article

    // Append to the Books table
    if result := h.DB.Find(&categories); result.Error != nil {
        fmt.Println(result.Error)
    }

    // Send a 201 created response
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(categories)

}


