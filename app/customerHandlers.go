package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Ammce/go-banking-core/dto/customerDTO"
	"github.com/Ammce/go-banking-core/service"
	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) createCustomer(w http.ResponseWriter, r *http.Request) {
	var body customerDTO.CreateCustomer
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		writeReponse(w, http.StatusBadRequest, err)
	} else {
		r, err := ch.service.CreateCustomer(body)
		if err != nil {
			writeReponse(w, err.Code, err)
		} else {
			writeReponse(w, http.StatusOK, r)
		}
	}
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomers(status)
	if err != nil {
		writeReponse(w, err.Code, err.AsMessage())
	} else {
		writeReponse(w, http.StatusOK, &customers)
	}
}

func (ch *CustomerHandlers) getCustomerById(w http.ResponseWriter, r *http.Request) {
	customerId64, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 32)
	customerId32 := int32(customerId64)
	customer, err := ch.service.GetCustomerById(customerId32)
	if err != nil {
		writeReponse(w, err.Code, err.AsMessage())
	} else {
		writeReponse(w, http.StatusOK, &customer)
	}

}

func writeReponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
