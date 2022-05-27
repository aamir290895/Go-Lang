package main

import (
 "net/http"   
 "github.com/jinzhu/gorm"
 "gorm.io/driver/sqlite"

)

var db *gorm.DB
var err error

type User struct{

   gorm.Model	
   Name string
   email string

}

func main(){
   db,err = gorm.Open(sqlite.Open("sqlite3","test.db"))
   if err != nil {
      panic("failed to connect database")
    }


}