package routes

import (
	"net/http"

	"github.com/Ammce/go-banking-core/handlers"
	"github.com/gorilla/mux"
)

func NewCustomerRoutes(r *mux.Router, ch handlers.CustomerHandlers) {
	r.HandleFunc("/", ch.GetAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/", ch.CreateCustomer).Methods(http.MethodPost)
	r.HandleFunc("/{id:[0-9]+}", ch.GetCustomerById).Methods(http.MethodGet)
	r.HandleFunc("/{id:[0-9]+}", ch.DeleteCustomerById).Methods(http.MethodDelete)
	r.HandleFunc("/{id:[0-9]+}", ch.UpdateCustomer).Methods(http.MethodPatch)
}
