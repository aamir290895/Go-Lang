package main

import (

	"net/http"
    "github.com/gorilla/mux"
	"gorm.io/gorm"

	"GOOGLE-NEWS/src/news" 
    "gorm.io/driver/postgres"


)

var err error

var db *gorm.DB




func Initialize() *gorm.DB{

  dsn := "host=localhost user=postgres password=aamir@2908 dbname=pg port=5432 sslmode=disable"
  
  db, err = gorm.Open(postgres.Open(dsn) ,&gorm.Config{})

  if err != nil {
    panic ("Failed to connect database")
  }

  return db

}


func HandleRequest() {

    DB := Initialize()
    h := news.New(DB)

  	r := mux.NewRouter()
    r.HandleFunc("/add", h.CreateNews).Methods("POST")
  	r.HandleFunc("/get", h.GetCategories).Methods("GET")

   
    http.ListenAndServe(":8081",r)
}

func main(){


  Initialize()
  HandleRequest()
     

}