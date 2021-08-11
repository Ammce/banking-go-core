package app

import (
	"log"
	"net/http"

	"github.com/Ammce/go-banking-core/domain"
	"github.com/Ammce/go-banking-core/service"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerREpositoryDB())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
