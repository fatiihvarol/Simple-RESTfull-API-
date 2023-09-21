package Router

import (
	"first_project/Controller"
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, pc *Controller.ProductController) {
	r.HandleFunc("/products", pc.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id:[0-9]+}", pc.GetProduct).Methods("GET")
	r.HandleFunc("/products/{id:[0-9]+}", pc.DeleteProduct).Methods("DELETE") // Add the DELETE route

	// Add more routes and controllers as needed
}
