package main

import (
	"GOOGLE-NEWS/src/news"
	"net/http"

	"github.com/gorilla/mux"
)

// var err error

// var db *gorm.DB



func HandleRequest() {


    DB := news.Initialize()
    h := news.New(DB)
    DB.AutoMigrate(&news.Article{})

  	r := mux.NewRouter().StrictSlash(true)
    r.HandleFunc("/add", h.CreateNews).Methods("POST")
  	r.HandleFunc("/get", h.GetCategories).Methods("GET")
    r.HandleFunc("/login",news.Login)
    r.HandleFunc("/home",news.Index)

    http.ListenAndServe(":8081",r)


}

func main(){


  HandleRequest()
     

}
