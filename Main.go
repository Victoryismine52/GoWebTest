package main

import (
	"html/template"
	"net/http"
)

type SplashPage struct { //Used to pass Title and question to main page
	Title    string
	Question string
}

var MembersList []Member // creates variable for storing all members in array (slice)

type Member struct {
	FirstName string
	LastName  string
}

func useraddhandler(w http.ResponseWriter, r *http.Request) {
	p := SplashPage{Title: "Add User", Question: "Please Enter your first and last name"}
	t, _ := template.ParseFiles("adduser.html")
	t.Execute(w, p)
}

func vueform(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("VueForm.html")
	t.Execute(w, nil)
}

func listsave(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024)
	if err != nil {

	}
	member := Member{
		FirstName: r.FormValue("firstname"),
		LastName:  r.FormValue("lastname"),
	}

	MembersList = append(MembersList, member)
	t, _ := template.ParseFiles("ListRange.html")
	t.Execute(w, MembersList)
}

func byehandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("bye.html")
	t.Execute(w, nil)
}

func splashhandler(w http.ResponseWriter, r *http.Request) {
	p := SplashPage{Title: "Welcome!", Question: "Would you like to become a member?"}
	t, _ := template.ParseFiles("Index.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", splashhandler)
	http.HandleFunc("/useradd", useraddhandler)
	http.HandleFunc("/ListRange", listsave)
	http.HandleFunc("/bye", byehandler)
	http.HandleFunc("/vueform", vueform)
	http.ListenAndServe(":8080", nil)
}
