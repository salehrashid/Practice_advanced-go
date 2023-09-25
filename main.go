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

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}

func root(writer http.ResponseWriter, request *http.Request) {
	tmplte := template.Must(template.ParseFiles("template/root.html"))
	if err := tmplte.Execute(writer, nil); err != nil {
		panic(err)
	}
}

func login(writer http.ResponseWriter, request *http.Request) {
	tmplte := template.Must(template.ParseFiles("template/login.html"))
	if err := tmplte.Execute(writer, nil); err != nil {
		panic(err)
	}
}

func signup(writer http.ResponseWriter, request *http.Request) {
	tmplte := template.Must(template.ParseFiles("template/signup.html"))
	if err := tmplte.Execute(writer, nil); err != nil {
		panic(err)
	}
}
