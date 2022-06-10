package users

import (
	"html/template"
	"net/http"
)

type User struct {

   Name string          `json :"name"`
   UserName string      `json: "username"`
   Password string      `json:"password"`
   Email string          `json:"email"`


}

func Login(w http.ResponseWriter , r *http.Request){

   temp,_ := template.ParseFiles("templates/login.html")

   temp.Execute(w  ,nil)


}