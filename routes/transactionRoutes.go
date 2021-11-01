package routes

import (
	"net/http"

	"github.com/Ammce/go-banking-core/handlers"
	"github.com/gorilla/mux"
)

func NewTransactionRoutes(r *mux.Router, ch handlers.TransactionHandlers) {
	r.HandleFunc("/", ch.CreateTransaction).Methods(http.MethodPost)
}
