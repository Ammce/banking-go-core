package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string
	City    string
	Zipcode int16
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customer := []Customer{{Name: "Amel", City: "Belgrade", Zipcode: 11000}}
	w.Header().Add("Content-type", "application/json")
	fmt.Println(r.Header.Get("Content-type"))
	json.NewEncoder(w).Encode(&customer)
}
