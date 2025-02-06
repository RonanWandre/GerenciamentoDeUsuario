package main

import (
	"fmt"
	"net/http"
	"user-management/config"
	"user-management/database"
	"user-management/routes"
)

func main() {
	fmt.Println("Iniciando servidor...")
	database.ConnectDatabase()
	routes.SetupRoutes()
	http.ListenAndServe(config.Port, nil)
}
