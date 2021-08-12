package app

import (
	"encoding/json"
	"net/http"

	"github.com/Ammce/go-banking-core/dto"
	"github.com/Ammce/go-banking-core/service"
	"github.com/gorilla/mux"
)

type AccountHandlers struct {
	service service.AccountService
}

func (ah *AccountHandlers) createAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var body dto.CreateAccountDTO
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writeReponse(w, http.StatusBadRequest, err.Error())
	} else {
		body.CustomerId = customerId
		account, appErr := ah.service.CreateAccount(&body)
		if appErr != nil {
			writeReponse(w, appErr.Code, appErr.Message)
		} else {
			writeReponse(w, http.StatusOK, account)
		}
	}
}
