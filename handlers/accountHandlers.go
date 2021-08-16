package handlers

import (
	"encoding/json"
	"net/http"

	account "github.com/Ammce/go-banking-core/domain/Account"
	"github.com/Ammce/go-banking-core/dto/accountDTO"
)

type AccountHandlers struct {
	service account.AccountService
}

func (ah *AccountHandlers) createAccount(w http.ResponseWriter, r *http.Request) {
	customerId := getIdFromRequest(r, "id")
	var body accountDTO.CreateAccountDTO
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writeReponse(w, http.StatusBadRequest, err.Error())
	} else {
		body.CustomerID = customerId
		account, appErr := ah.service.CreateAccount(&body)
		if appErr != nil {
			writeReponse(w, appErr.Code, appErr.Message)
		} else {
			writeReponse(w, http.StatusOK, account)
		}
	}
}
