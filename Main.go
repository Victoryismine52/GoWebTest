package main

import (
	"html/template"
	"net/http"
)

type SplashPage struct {
	Title    string
	Question string
}

var MembersList []Member // Attempted to do through an array but commenting out for now

type UserPage struct {
	Users        string
	Current_User string
}

type UserPageList struct {
	MembersList  []Member
	Current_User string
}

type Member struct {
	FirstName string
	LastName  string
}

var members = ""

//var MemberList = []Member{} // Attempted to do through an array but commenting out for now

func useraddhandler(w http.ResponseWriter, r *http.Request) {
	p := SplashPage{Title: "Add User", Question: "Please Enter your first and last name"}
	t, _ := template.ParseFiles("adduser.html")
	t.Execute(w, p)
}

func vueform(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("VueForm.html")
	t.Execute(w, nil)
}

func save(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024)
	if err != nil {

	}
	member := Member{
		FirstName: r.FormValue("firstname"),
		LastName:  r.FormValue("lastname"),
	}

	MembersList = append(MembersList, member)
	members = members + member.FirstName + " " + member.LastName + ",\n"
	//fmt.Fprintf(w, "Hello, %s %s Welcome to the club \n", member.FirstName, member.LastName)
	//fmt.Fprintf(w, "Your part of an exclusive club including \n%s", members)
	//fmt.Printf("%#v", "Your part of an exclusive club including %s", members)
	p := UserPage{Users: members, Current_User: member.FirstName}
	t, _ := template.ParseFiles("UserList.html")
	t.Execute(w, p)
	/*
		f, err := os.OpenFile("members.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer f.Close()

		b, err := json.Marshal(MembersList)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		f.Write(b)
		f.Close()
		p := UserPage{Users: members, Current_User: member.FirstName}
		t, _ := template.ParseFiles("UserList.html")
		t.Execute(w, p)
	*/
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
	p := UserPageList{MembersList: MembersList, Current_User: member.FirstName + " " + member.LastName}
	//fmt.Fprintf(w, "Hello, %s %s Welcome to the club \n", member.FirstName, member.LastName)
	//fmt.Fprintf(w, "Your part of an exclusive club including \n%s", members)
	//fmt.Printf("%#v", "Your part of an exclusive club including %s", members)
	t, _ := template.ParseFiles("ListRange.html")
	t.Execute(w, p)
	/*
		f, err := os.OpenFile("members.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer f.Close()

		b, err := json.Marshal(MembersList)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		f.Write(b)
		f.Close()
		p := UserPage{Users: members, Current_User: member.FirstName}
		t, _ := template.ParseFiles("UserList.html")
		t.Execute(w, p)
	*/
}

func byehandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("bye.html")
	t.Execute(w, nil)
}

func splashhandler(w http.ResponseWriter, r *http.Request) {
	p := SplashPage{Title: "Welcome!", Question: "Would you like to become a member?"}
	t, _ := template.ParseFiles("Splash2.html")
	t.Execute(w, p)
}

func userlisthandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("UserList.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", splashhandler)
	http.HandleFunc("/useradd", useraddhandler)
	http.HandleFunc("/save", save)
	http.HandleFunc("/ListRange", listsave)
	http.HandleFunc("/bye", byehandler)
	http.HandleFunc("/vueform", vueform)
	http.HandleFunc("/UserList", userlisthandler)
	http.ListenAndServe(":8080", nil)
}
