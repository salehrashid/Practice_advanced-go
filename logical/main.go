package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"golang-advanced/constants"
	"html/template"
	"net/http"
	"strings"
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

	if addUser(username[0], password[0]) {
		tmplte := template.Must(template.ParseFiles("template/index.html"))
		if err := tmplte.Execute(writer, nil); err != nil {
			panic(err)
		}
	} else {
		tmplte := template.Must(template.ParseFiles("template/error.html"))
		if err := tmplte.Execute(writer, nil); err != nil {
			panic(err)
		}
	}
}

func loginUser(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		panic(err)
	}

	username := request.Form["nama"]     //diambil dari attribute html name
	password := request.Form["password"] //diambil dari attribute html name

	if checkUser(username[0], password[0]) {
		tmplte := template.Must(template.ParseFiles("template/index.html"))
		if err := tmplte.Execute(writer, nil); err != nil {
			panic(err)
		}
	} else {
		tmplte := template.Must(template.ParseFiles("template/error.html"))
		if err := tmplte.Execute(writer, nil); err != nil {
			panic(err)
		}
	}
}

func addUser(username string, password string) bool {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		constants.Host, constants.Port, constants.User, constants.Password, constants.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//whitespace validation
	validUsername := strings.TrimSpace(username)
	validPassword := strings.TrimSpace(password)
	if validUsername == "" || validPassword == "" {
		return false
	}

	insertQuery := `INSERT INTO users(name, password) VALUES ($1,$2)`

	add, err := db.Query(insertQuery, username, password)
	if err != nil {
		panic(err)
	}
	defer add.Close()
	return true
}

func checkUser(username string, password string) bool {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		constants.Host, constants.Port, constants.User, constants.Password, constants.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	var exists bool
	var query string
	checkUserQuery := `SELECT EXISTS(SELECT name FROM users WHERE name='%s' AND password='%s')`

	query = fmt.Sprintf(checkUserQuery, username, password)
	row := db.QueryRow(query).Scan(&exists)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	fmt.Println(row)
	defer db.Close()
	return exists
}
