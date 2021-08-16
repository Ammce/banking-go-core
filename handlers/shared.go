package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func writeReponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func getIdFromRequest(r *http.Request, variable string) int32 {
	customerId64, _ := strconv.ParseInt(mux.Vars(r)[variable], 10, 32)
	customerId32 := int32(customerId64)
	return customerId32
}
