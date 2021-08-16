package routes

import (
	"net/http"

	"github.com/Ammce/go-banking-core/handlers"
	"github.com/gorilla/mux"
)

func NewAccountRoutes(r *mux.Router, ah handlers.AccountHandlers) {
	r.HandleFunc("/", ah.CreateAccount).Methods(http.MethodPost)
}
