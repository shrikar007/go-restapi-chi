package main

import (
"encoding/json"
"log"
"net/http"

"github.com/gorilla/mux"
)

type Emp struct {
Id        string `json:"id,omitempty"`
FirstName string `json:"firstName,omitempty"`
LastName  string `json:"lastName,omitempty"`
Age       int    `json:"age,omitempty"`
}

var emps []Emp

func Getall(w http.ResponseWriter, r *http.Request) {
       json.NewEncoder(w).Encode(emps)
}

func Getbyid(w http.ResponseWriter, r *http.Request) {
       params := mux.Vars(r)
       for _, emp := range emps {
              if emp.Id == params["id"] {
                     json.NewEncoder(w).Encode(emp)
                     return
              }
       }
       json.NewEncoder(w).Encode(&Emp{})
}

func Create(w http.ResponseWriter, r *http.Request) {
       var emp Emp
       _ = json.NewDecoder(r.Body).Decode(&emp)
       emps = append(emps, emp)
       json.NewEncoder(w).Encode(emp)
}

func Delete(w http.ResponseWriter, r *http.Request) {
       params := mux.Vars(r)
       for index, emp := range emps {
              if emp.Id == params["id"] {
                     emps = append(emps[:index], emps[index+1:]...)
                     break
              }
              json.NewEncoder(w).Encode(emps)
       }
}

func main() {
       r := mux.NewRouter()
       emps = append(emps, Emp{"1", "shrikar", "vaitala", 22})
       emps = append(emps, Emp{"2", "pratham", "yemul", 23})

       r.HandleFunc("/emp", Getall).Methods("GET")
       r.HandleFunc("/emp/{id}", Getbyid).Methods("GET")
       r.HandleFunc("/emp", Create).Methods("POST")
       r.HandleFunc("/emp/{id}", Delete).Methods("DELETE")
       log.Fatal(http.ListenAndServe(":8082", r))
}