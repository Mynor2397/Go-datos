package controllers

import (
	"net/http"
	"fmt"
	"text/template"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
	"../Database"
	
)
var db *sql.DB
var err error
var tmpls = template.Must(template.ParseGlob("views/login/*"))

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpls.ExecuteTemplate(w, "login", nil)
}

func SignupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		tmpls.ExecuteTemplate(res, "signup", nil)
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var user string

	db :=Database.Dbconn()
	tsql := fmt.Sprintf("SELECT username FROM employee.dbo.users WHERE username='%s';", username)
	err := db.QueryRow(tsql).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(res, "Error del servidor, No se puede crear su cuenta.", 500)
			return
		}

		tsql:=fmt.Sprintf("INSERT INTO employee.dbo.users(username, password) VALUES('%s', '%s');", username, hashedPassword)
		_, err = db.Exec(tsql)
		if err != nil {
			http.Error(res, "Error del servidor, No se puede crear su cuenta.", 500)
			return
		}

		res.Write([]byte("Usuario registrado!"))
		return
	case err != nil:
		http.Error(res, "Error del servidor, No se puede crear su cuenta.", 500)
		return
	default:
		http.Redirect(res, req, "/home", 301)
	}
}

func LoginPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		tmpls.ExecuteTemplate(res, "login", nil)
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var databaseUsername string
	var databasePassword string

	db:=Database.Dbconn()

	tsql:=fmt.Sprintf("SELECT username, password FROM employee.dbo.users WHERE username='%s'", username)
	err := db.QueryRow(tsql).Scan(&databaseUsername, &databasePassword)

	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		http.Redirect(res, req, "/login", 301)		
		return
	}
	
	http.Redirect(res, req, "/index", 301)
	res.Write([]byte("Hola Bienvenido:  " + databaseUsername))


}

func message()(string, bool){
	return "hola mundo", true
}
