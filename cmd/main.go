package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"template/internal/controller"
	"template/internal/data"
	"template/internal/service"
)

func main() {
	// Open database connection
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/template")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create repository and service instances
	userRepository := data.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	// Create controller instance
	userController := controller.NewUserController(userService)

	// Create router and register routes
	router := mux.NewRouter()
	router.HandleFunc("/user", userController.CreateUser).Methods("POST")
	router.HandleFunc("/user", userController.GetUser).Methods("GET")

	// Start server
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
