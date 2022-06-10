package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"GOOGLE-NEWS/src/news"
	"GOOGLE-NEWS/src/users"

	"gorm.io/driver/postgres"
)

var err error

var db *gorm.DB

func init(){

  Initialize()
}


func Initialize() *gorm.DB{

  dsn := "host=localhost user=postgres password=aamir@2908 dbname=crud port=5432"
  
  db, err = gorm.Open(postgres.Open(dsn) ,&gorm.Config{})

  if err != nil {
    panic ("Failed to connect database")
  }

  return db

}


func HandleRequest() {

    DB := Initialize()
    h := news.New(DB)
    DB.AutoMigrate(&news.Article{})

  	r := mux.NewRouter().StrictSlash(true)
    r.HandleFunc("/add", h.CreateNews).Methods("POST")
  	r.HandleFunc("/get", h.GetCategories).Methods("GET")
    r.HandleFunc("/login", users.Login)

    http.ListenAndServe(":8081",r)


}

func main(){


  HandleRequest()
     

}
