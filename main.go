package main

import (
	"net/http"

	"API-0.1/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/users", controllers.GetUsers)                          //Devuelve todos los usuarios
	r.HandleFunc("/api/user/{id}", controllers.GetUser).Methods("GET", "PUT") //Devuelve un usuario en especifico
	r.HandleFunc("/api/user", controllers.UserPost)                           //Crea un usuario
	r.HandleFunc("/api/user/update/{id}", controllers.UpdateUser)             //Actualiza un usuario especifico
	r.HandleFunc("/api/user/delete/{id}", controllers.DeleteUser)             //Elimina un usuario especifico

	http.ListenAndServe(":3000", r)
}
