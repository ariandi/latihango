package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"latihan1/services"
	"net/http"
)

type CustomerHandlers struct {
	service services.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers (w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{"Ari", "Bogor", "36160"},
	//	{"Susi", "Tangerang", "36163"},
	//	{"Foi", "Bogor", "36163"},
	//	{"Puta", "Tangerang", "36163"},
	//	{"Mail", "Tangerang", "36163"},
	//}

	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomer(status)
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, customers)
	}
}


func (ch *CustomerHandlers) getCustomer (w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{"Ari", "Bogor", "36160"},
	//	{"Susi", "Tangerang", "36163"},
	//	{"Foi", "Bogor", "36163"},
	//	{"Puta", "Tangerang", "36163"},
	//	{"Mail", "Tangerang", "36163"},
	//}
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, customer)
	}
}

func WriteResponse(w http.ResponseWriter, code int, data interface{})  {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}