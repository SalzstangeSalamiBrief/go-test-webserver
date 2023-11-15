package main

import (
	"go-http/config"
	"go-http/database"
	"go-http/router"
	"go-http/useCases"
	"net/http"
)

const ADDR = "localhost:3000"

func main() {
	prepareDatabase()
	addRoutesToApp()
	http.HandleFunc("/", router.HandleRouting)
	http.ListenAndServe(ADDR, nil)
}

func prepareDatabase() {
	c := config.ReadConfig()
	database.ConnectToDb(c["db_username"], c["db_password"], c["db_name"])
	database.CreateSchema()
	database.SeedData()
}

func addRoutesToApp() {
	router.Routes.AddRoute(useCases.CreateUserUseCase)
	router.Routes.AddRoute(useCases.GetUserUseCase)
	router.Routes.AddRoute(useCases.GetUsersUseCase)
	router.Routes.AddRoute(useCases.EditUserUseCase)
	router.Routes.AddRoute(useCases.DeleteUserUseCase)
	router.Routes.AddRoute(useCases.DownloadFileUseCase)
	router.Routes.AddRoute(useCases.UploadFileUseCase)
}
