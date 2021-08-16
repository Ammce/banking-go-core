package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	customer "github.com/Ammce/go-banking-core/domain/Customer"
	"github.com/Ammce/go-banking-core/dto/customerDTO"
	"github.com/gorilla/mux"
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
	customerId := getCustomerIdFromRequest(r)
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
	customerId64, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 32)
	customerId32 := int32(customerId64)
	customer, err := ch.service.GetCustomerById(customerId32)
	if err != nil {
		writeReponse(w, err.Code, err.AsMessage())
	} else {
		writeReponse(w, http.StatusOK, &customer)
	}
}

func (ch *CustomerHandlers) DeleteCustomerById(w http.ResponseWriter, r *http.Request) {
	customerId64, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 32)
	customerId32 := int32(customerId64)
	err := ch.service.DeleteCustomerById(customerId32)
	if err != nil {
		writeReponse(w, err.Code, err.AsMessage())
	} else {
		writeReponse(w, http.StatusOK, nil)
	}
}

func writeReponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func getCustomerIdFromRequest(r *http.Request) int32 {
	customerId64, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 32)
	customerId32 := int32(customerId64)
	return customerId32
}

func NewCustomerHandlers(service customer.CustomerService) CustomerHandlers {
	return CustomerHandlers{
		service: service,
	}
}
