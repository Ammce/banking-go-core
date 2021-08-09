package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerId:[0-9]+}", getCustomerById).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}