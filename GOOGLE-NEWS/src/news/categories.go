package news

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



type Category struct {

	gorm.Model

	CategoryName string        `json:"category" gorm:"primarykey"`

	Articles []Article         `json:"articles" gorm:"foreignkey: GUID"`

}

type Article  struct {
  
	gorm.Model
	GUID string           `json:"guid" gorm:"primaryKey"`          
	Title string          `json:"title"`
	Description string    `json:"description"`
	Url string            `json:"url"`
	ImageUrl string       `json:"image_url"`
	PostDate string       `json:"posted_date"`
	Tags string           `json:"tags"`
	Author string         `json:"author"`
  
  }

var db *gorm.DB

var err error

type Handler struct {
    DB *gorm.DB
}

func New(db *gorm.DB) Handler {
    return Handler{db}
}


func Initialize() *gorm.DB{


	dsn := "host=localhost user=postgres password=aamir@2908 dbname=crud port=5432"
	
	db, err = gorm.Open(postgres.Open(dsn) ,&gorm.Config{})
  
	if err != nil {
	  panic ("Failed to connect database")
	}
  
	return db
  
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
		article.ImageUrl = "data:image/png;base64,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"
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


