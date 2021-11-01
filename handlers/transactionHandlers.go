package handlers

import (
	"encoding/json"
	"net/http"

	transaction "github.com/Ammce/go-banking-core/domain/Transaction"
)

type TransactionHandlers struct {
	service transaction.TransactionService
}

func (ts TransactionHandlers) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var body transaction.Transaction
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writeReponse(w, http.StatusBadRequest, err)
	} else {
		r, err := ts.service.CreateTransaction(body)
		if err != nil {
			writeReponse(w, http.StatusBadGateway, err)
		} else {
			writeReponse(w, http.StatusOK, r)
		}
	}

}
