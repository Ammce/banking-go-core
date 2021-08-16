package handlers

import (
	"encoding/json"
	"net/http"

	customer "github.com/Ammce/go-banking-core/domain/Customer"
	"github.com/Ammce/go-banking-core/dto/customerDTO"
)

type CustomerHandlers struct {
	service customer.CustomerService
}

func (ch CustomerHandlers) CreateCustomer(w http.ResponseWriter, r *http.Request) {
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

func (ch CustomerHandlers) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var body customerDTO.CreateCustomer
	err := json.NewDecoder(r.Body).Decode(&body)
	customerId := getIdFromRequest(r, "id")
	if err != nil {
		writeReponse(w, http.StatusBadRequest, err)
	} else {
		r, err := ch.service.UpdateCustomer(customerId, body)
		if err != nil {
			writeReponse(w, err.Code, err)
		} else {
			writeReponse(w, http.StatusOK, r)
		}
	}
}

func (ch CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomers(status)
	if err != nil {
		writeReponse(w, err.Code, err.AsMessage())
	} else {
		writeReponse(w, http.StatusOK, &customers)
	}
}

func (ch *CustomerHandlers) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	customerId := getIdFromRequest(r, "id")
	customer, err := ch.service.GetCustomerById(customerId)
	if err != nil {
		writeReponse(w, err.Code, err.AsMessage())
	} else {
		writeReponse(w, http.StatusOK, &customer)
	}
}

func (ch *CustomerHandlers) DeleteCustomerById(w http.ResponseWriter, r *http.Request) {
	customerId := getIdFromRequest(r, "id")
	err := ch.service.DeleteCustomerById(customerId)
	if err != nil {
		writeReponse(w, err.Code, err.AsMessage())
	} else {
		writeReponse(w, http.StatusOK, nil)
	}
}

func NewCustomerHandlers(service customer.CustomerService) CustomerHandlers {
	return CustomerHandlers{
		service: service,
	}
}
