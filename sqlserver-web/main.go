package main

import (
	"log"
	"net/http"
	"./controllers"
)

func main() {
	
	http.Handle("/", http.FileServer(http.Dir("./public")))

	log.Println("Server started on: http://localhost:8084")
	http.HandleFunc("/index", controllers.Index)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)

	//HANDLERS PARA SESIONES
	http.HandleFunc("/home", controllers.HomePage)
	http.HandleFunc("/signup", controllers.SignupPage)
	http.HandleFunc("/login", controllers.LoginPage)
	http.ListenAndServe(":8084", nil)
}
