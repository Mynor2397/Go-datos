package controllers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"../Database"
)


type persona struct {
	Id       string
	Name     string
	Location string
}

//VARIABLE PUBLICA PARA TOMAR LO QUE HAY EN CARPETA VIEWS
var tmpl = template.Must(template.ParseGlob("views/employee/*"))

//MOSTRAR EL FORMULARIO PAR INSERTAR 
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

//INSERTAR 
func Insert(w http.ResponseWriter, r *http.Request) {
	db := Database.Dbconn()

	if r.Method == "POST" {

		Id := r.FormValue("Id")
		Name := r.FormValue("Name")
		Location := r.FormValue("Location")

		tsql := fmt.Sprintf("INSERT INTO employee.dbo.employee (Id, Name, Location) VALUES('%s', '%s', '%s');", Id, Name, Location)
		_, err := db.Exec(tsql)
		if err != nil {
			fmt.Println("error en la consulta: " + err.Error())
		}
		log.Println(r.Form)
		log.Println("insert" + " " + Id + " " + Name + " " + Location)
	}
	defer db.Close()
	http.Redirect(w, r, "/index", 301)

}

//MUESTRA EL SELECT DE LOS USUARIOS REGISTRADOS
func Index(w http.ResponseWriter, r *http.Request) {
	
	db := Database.Dbconn()
	selDB, err := db.Query("SELECT * FROM employee.dbo.employee")
	if err != nil {
		panic(err.Error())
	}

	per := persona{}
	res := []persona{}

	for selDB.Next() {
		var id string
		var name string
		var location string

		err = selDB.Scan(&id, &name, &location)
		if err != nil {
			panic(err.Error())
		}
		per.Id = id
		per.Name = name
		per.Location = location

		res = append(res, per)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

//ACTUALIZA LOS DATOS DE LOS REGISTROS
func Update(w http.ResponseWriter, r *http.Request) {
	db := Database.Dbconn()
	if r.Method == "POST" {

		Id := r.FormValue("Id")
		Name := r.FormValue("Name")
		Location := r.FormValue("Location")

		tsql := fmt.Sprintf("UPDATE employee.dbo.employee SET Name='%s', Location='%s'  WHERE Id='%s';", Name, Location, Id)
		_, err := db.Exec(tsql)
		if err != nil {
			panic(err.Error())
		}
		log.Println("UPDATE: Id: " + Id + " | Name: " + Name + "  | Location: "+Location)
	}
	defer db.Close()
	http.Redirect(w, r, "/index", 301)
}

//MUESTRA EL REGISTRO EN EL NUEVO FORMULARIO SEGUN EL ID
func Edit(w http.ResponseWriter, r *http.Request) {
	db := Database.Dbconn()
	
	nId := r.URL.Query().Get("id")
	tsql := fmt.Sprintf("SELECT * FROM employee.dbo.employee where Id='%s';", nId)
	selDB, err := db.Query(tsql)
	if err != nil {
        panic(err.Error())
	}

	per :=persona{}
	for selDB.Next(){
		var (
			Id string
			Name string
			Location string
		)

		err =selDB.Scan(&Id, &Name, &Location)
		if err != nil {
			panic(err.Error())
		}
		per.Id=Id
		per.Name = Name
		per.Location = Location
	}

	log.Println(per)
	tmpl.ExecuteTemplate(w, "Show", per)

}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := Database.Dbconn()
	nId := r.URL.Query().Get("id")
	tsql := fmt.Sprintf("DELETE FROM Employee WHERE id='%s';", nId)
    delForm, err := db.Prepare(tsql)
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(nId)
    log.Println("DELETE")
    defer db.Close()
    http.Redirect(w, r, "/index", 301)
}


