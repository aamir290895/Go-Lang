package main

import (

	"net/http"
  "github.com/gorilla/mux"
	"gorm.io/gorm"

	"gorm/src/users" 
  "gorm.io/driver/postgres"


)

var err error

var db *gorm.DB




func Initialize() *gorm.DB{

  dsn := "host=localhost user=postgres password=aamir@2908 dbname=crud port=5432 sslmode=disable"
  
  db, err = gorm.Open(postgres.Open(dsn) ,&gorm.Config{})

  if err != nil {
    panic ("Failed to connect database")
  }

  db.AutoMigrate(&users.Employee{})
  return db

}


func HandleRequest() {

    DB := Initialize()
    h := users.New(DB)

  	r := mux.NewRouter()
    r.HandleFunc("/get", h.GetUsers).Methods("GET")
    r.HandleFunc("/add", h.PostUser).Methods("POST")
    r.HandleFunc("/update", h.PutUser).Methods("PUT")
    r.HandleFunc("/delete/{id}", h.DeleteUser).Methods("DELETE")
    http.ListenAndServe(":8081",r)
}

func main(){


  Initialize()
  HandleRequest()
     

}