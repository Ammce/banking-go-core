package app

import (
	"encoding/json"
	"net/http"

	"github.com/Ammce/go-banking-core/domain"
	"github.com/Ammce/go-banking-core/service"
)

type TransactionHandlers struct {
	service service.TransactionService
}

func (th TransactionHandlers) createTransaction(w http.ResponseWriter, r *http.Request) {
	var body domain.Transaction
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writeReponse(w, http.StatusBadRequest, err.Error())
	} else {
		t, err := th.service.Create(body)
		if err != nil {
			writeReponse(w, err.Code, err.AsMessage())
		} else {
			writeReponse(w, http.StatusOK, *t)
		}
	}
}
