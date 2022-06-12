package news

import (
	"fmt"
	"html/template"
	"net/http"
	// "gorm.io/gorm"
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



func Signup (w http.ResponseWriter , r *http.Request){
   temp,_ := template.ParseFiles("templates/login.html")

   temp.Execute(w  ,nil)
}



func Index(w http.ResponseWriter , r *http.Request){
 
   db := Initialize()
   var articles []Article

   result := db.Find(&articles)
   if result.Error != nil {
      fmt.Println(result.Error)
  }

   var article Article
   for _,x := range articles{
         
      article = Article{

         Title : x.Title,
         Description : x.Description,
         ImageUrl : x.ImageUrl,
      }
      temp,_ := template.ParseFiles("templates/index.html")
      temp.Execute(w ,article)
   }

   fmt.Println(article)
   // temp,_ := template.ParseFiles("templates/index.html")
   // temp.Execute(w ,articles)
}

