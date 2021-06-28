// +build ignore
package main

import (
	"net/http"
	"html/template"
	"log"
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	// "github.com/gorilla/securecookie"
	"example.com/memog/utils"
)

type Todo struct {
	Task string
	Date string
	Deadline string
}

// var todolist = []Todo

var templates = template.Must(template.ParseFiles("index.html", "login.html"))

func login_check(w http.ResponseWriter, r *http.Request){
	userName := utils.GetUserName(r)
	if userName == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func index(w http.ResponseWriter, r *http.Request){
	login_check(w, r)
	//
	// read database
	//
	var todo = Todo{"todo", "date", "deal"}
	var todolist = []Todo{todo}
	err := templates.ExecuteTemplate(w, "index.html", todolist)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func save(w http.ResponseWriter, r *http.Request){
	var todo = Todo{}
	todo.Task = r.FormValue("task")
	todo.Date = r.FormValue("date")
	todo.Deadline = r.FormValue("deadline")
	//
	// save database
	//

	http.Redirect(w, r, "/", http.StatusFound)
}

func login(w http.ResponseWriter, r *http.Request){
	err := templates.ExecuteTemplate(w, "login.html", nil)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var router = mux.NewRouter()

func main() {

	router.HandleFunc("/", index)
	router.HandleFunc("/save", save)

	router.HandleFunc("/login", login)
	router.HandleFunc("/login_internal", utils.LoginHandler).Methods("POST")
	router.HandleFunc("/logout", utils.LogoutHandler).Methods("POST")

	http.Handle("/", router)
	
	// http.HandleFunc("/", index)
	// http.HandleFunc("/save", save)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
