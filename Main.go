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

type Member struct { //member structure to store current item
	FirstName string
	LastName  string
}

//handler to adduser.html let's user add first and last name with navigation options to add or go back
func useraddhandler(w http.ResponseWriter, r *http.Request) {
	p := SplashPage{Title: "Add User", Question: "Please Enter your first and last name"}
	t, _ := template.ParseFiles("adduser.html")
	t.Execute(w, p)
}

//Vue version of entire app currently only displays form and outputs to console
func vueform(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("VueForm.html")
	t.Execute(w, nil)
}

//stores form information to member structure and then adds it to MembersList slice and passes it to ListRange.html for display of all usernames
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

//Fun way to say bye to user
func byehandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("bye.html")
	t.Execute(w, nil)
}

//Serves up home page with configurable Title and Question (want to randomize for fun if time)
func splashhandler(w http.ResponseWriter, r *http.Request) {
	p := SplashPage{Title: "Welcome!", Question: "Would you like to become a member?"}
	t, _ := template.ParseFiles("Index.html")
	t.Execute(w, p)
}

//Main function with all navigation points and current port (obviscated by Docker install to port 80)
func main() {
	http.HandleFunc("/", splashhandler)
	http.HandleFunc("/useradd", useraddhandler)
	http.HandleFunc("/ListRange", listsave)
	http.HandleFunc("/bye", byehandler)
	http.HandleFunc("/vueform", vueform)
	http.ListenAndServe(":8080", nil)
}
