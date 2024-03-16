package main

import (
	"accounts"
	"db"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	// Connect to the database
	db.Connect()
	db.Db.AutoMigrate(&accounts.Users{})

	//connect to rabbitmq
	// amqp.ConnectAMQP()

	// initialize the router
	router := chi.NewRouter()

	// middlware section
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	}))

	// define the routes
	router.Get("/api/v1/users/read", accounts.GetUsersController)
	router.Post("/api/v1/users/signup", accounts.SignUpController)
	router.Post("/api/v1/users/login", accounts.LoginController)

	// init server logs
	fmt.Println("Server Listening on port :8080")
	fmt.Println("To Quit, click CTRL + C")
	http.ListenAndServe(":8080", router)
}
