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

func getIdFromRequest(r *http.Request, variable string) uint {
	Id64, _ := strconv.ParseInt(mux.Vars(r)[variable], 10, 32)
	Id32 := int32(Id64)
	return uint(Id32)
}
