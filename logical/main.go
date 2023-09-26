package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("udah jalan bang servernya http://localhost:8000")

	http.HandleFunc("/", root)

	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login-user", loginUser)
	http.HandleFunc("/signup-user", signupUser)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}

func root(writer http.ResponseWriter, _ *http.Request) {
	tmplte := template.Must(template.ParseFiles("template/root.html"))
	if err := tmplte.Execute(writer, nil); err != nil {
		panic(err)
	}
}

func login(writer http.ResponseWriter, _ *http.Request) {
	tmplte := template.Must(template.ParseFiles("template/login.html"))
	if err := tmplte.Execute(writer, nil); err != nil {
		panic(err)
	}
}

func signup(writer http.ResponseWriter, _ *http.Request) {
	tmplte := template.Must(template.ParseFiles("template/signup.html"))
	if err := tmplte.Execute(writer, nil); err != nil {
		panic(err)
	}
}

func signupUser(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		panic(err)
	}
	username := request.Form["nama"]     //diambil dari attribute html name
	password := request.Form["password"] //diambil dari attribute html name
	fmt.Println(username, " ", password)
	tmplte := template.Must(template.ParseFiles("template/index.html"))
	if err := tmplte.Execute(writer, nil); err != nil {
		panic(err)
	}
}

func loginUser(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		panic(err)
	}
	username := request.Form["nama"]     //diambil dari attribute html name
	password := request.Form["password"] //diambil dari attribute html name
	fmt.Println(username, " ", password)
	tmplte := template.Must(template.ParseFiles("template/index.html"))
	if err := tmplte.Execute(writer, nil); err != nil {
		panic(err)
	}
}
