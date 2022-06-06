package users

import (

	"net/http"
	"gorm.io/gorm"
	"fmt"
	"encoding/json"
    "log"
    "io/ioutil"
    "github.com/gorilla/mux"
    "strconv"


)


type Handler struct {
    DB *gorm.DB
}

func New(db *gorm.DB) Handler {
    return Handler{db}
}

type Employee struct {
	ID int         `json:"id" gorm:"primaryKey"`
	Name string    `json:"name"`
	Age int        `json:"age"`
  }
  

func(h Handler) GetUsers (w http.ResponseWriter , r *http.Request){
    var employees []Employee

    if result := h.DB.Find(&employees); result.Error != nil {
        fmt.Println(result.Error)
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(employees)
}


func(h Handler) PostUser (w http.ResponseWriter , r *http.Request){
	defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var employee Employee
    json.Unmarshal(body, &employee)

    // Append to the Books table
    if result := h.DB.Create(&employee); result.Error != nil {
        fmt.Println(result.Error)
    }

    // Send a 201 created response
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")

}


func(h Handler) PutUser (w http.ResponseWriter , r *http.Request){
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    // Read request body
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var employeeOld Employee
    json.Unmarshal(body, &employeeOld)

    var employeeNew Employee

    if result := h.DB.First(&employeeOld, id); result.Error != nil {
        fmt.Println(result.Error)
    }

    employeeNew.ID = employeeOld.ID
    employeeNew.Name = employeeOld.Name
    employeeNew.Age = employeeOld.Age

    h.DB.Save(&employeeNew)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Updated")

}


func(h Handler) DeleteUser (w http.ResponseWriter , r *http.Request){
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])


    var employee Employee

    if result := h.DB.First(&employee, id); result.Error != nil {
        fmt.Println(result.Error)
    }

    // Delete that book
    h.DB.Delete(&employee)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Deleted")
}



