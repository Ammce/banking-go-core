package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ammce/go-banking-core/service"
	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomers()
	w.Header().Add("Content-type", "application/json")
	fmt.Println(r.Header.Get("Content-type"))
	json.NewEncoder(w).Encode(&customers)
}

func (ch *CustomerHandlers) getCustomerById(w http.ResponseWriter, r *http.Request) {
	customerId64, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 32)
	customerId32 := int32(customerId64)
	customer, err := ch.service.GetCustomerById(customerId32)
	if err != nil {
		writeReponse(w, err.Code, err.AsMessage())
	}
	writeReponse(w, http.StatusOK, &customer)
}

func writeReponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
