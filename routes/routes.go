package routes

import (
	"net/http"
	"user-management/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/users", controllers.GetUsers)
	http.HandleFunc("/users/create", controllers.CreateUser)
}
