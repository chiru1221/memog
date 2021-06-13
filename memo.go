// +build ignore
package main

import (
	"net/http"
	"html/template"
	"log"
)

type Todo struct {
	Task string
	Date string
	Deadline string
}
var todo = Todo{}

var templates = template.Must(template.ParseFiles("index.html"))

func index(w http.ResponseWriter, r *http.Request){
	err := templates.ExecuteTemplate(w, "index.html", todo)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func save(w http.ResponseWriter, r *http.Request){
	todo.Task = r.FormValue("task")
	todo.Date = r.FormValue("date")
	todo.Deadline = r.FormValue("deadline")
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/save", save)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
