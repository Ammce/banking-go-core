package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ammce/go-banking-core/service"
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
