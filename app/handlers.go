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
	customer, _ := ch.service.GetCustomerById(customerId32)
	w.Header().Add("Content-type", "application/json")
	fmt.Println(r.Header.Get("Content-type"))
	json.NewEncoder(w).Encode(&customer)
}
