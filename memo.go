package main

import (
	"net/http"
	"html/template"
	"log"
	"strconv"
	"github.com/gorilla/mux"
	"example.com/memog/utils"
)

var templates = template.Must(template.ParseFiles("index.html", "login.html"))

func login_check(w http.ResponseWriter, r *http.Request){
	userName, _ := utils.GetUser(r)
	if userName == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func index(w http.ResponseWriter, r *http.Request){
	login_check(w, r)
	_, user_id := utils.GetUser(r)
	// read database
	todolist := utils.ReadDB(user_id)

	err := templates.ExecuteTemplate(w, "index.html", todolist)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func save(w http.ResponseWriter, r *http.Request){
	login_check(w, r)

	// insert form values into structure
	var todo = utils.Todo{}
	todo.Task = r.FormValue("task")
	todo.Date = r.FormValue("date")
	todo.Deadline = r.FormValue("deadline")
	_, todo.User = utils.GetUser(r)
	
	var err int
	if (r.FormValue("id") != "") {
		todo.Id, _ = strconv.Atoi(r.FormValue("id"))
		// update database
		err = utils.UpdateDB(todo)
	}else{
		// insert database
		err = utils.InsertDB(todo)
	}

	if err == 1{
		log.Fatal("insert or update db error")
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func delete(w http.ResponseWriter, r *http.Request){
	login_check(w, r)

	var todo = utils.Todo{}
	todo.Id, _ = strconv.Atoi(r.FormValue("id"))
	
	err := utils.DeleteDB(todo)
	if err == 1{
		log.Fatal("delete db error")
	}

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
	router.HandleFunc("/save", save).Methods("POST")
	router.HandleFunc("/delete", delete).Methods("POST")

	router.HandleFunc("/login", login)
	router.HandleFunc("/login_internal", utils.LoginHandler).Methods("POST")
	router.HandleFunc("/logout", utils.LogoutHandler).Methods("POST")
	router.HandleFunc("/regist", utils.RegistHandler).Methods("POST")

	http.Handle("/", router)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
