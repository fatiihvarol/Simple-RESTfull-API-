package main

import (
	"first_project/Controller"
	"first_project/Initialize"
	"first_project/Router" // Import the router package
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db, err := Initialize.InitializeDB()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	pc := Controller.NewProductController(db)

	// Use the router package to set up routes
	Router.SetupRoutes(r, pc)

	// Start the HTTP server
	http.Handle("/", r)
	fmt.Println("HTTP server started on :8080")
	http.ListenAndServe(":8080", nil)
}
