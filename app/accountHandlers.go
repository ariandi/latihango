package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"latihan1/dto"
	"latihan1/services"
	"net/http"
)

type AccountHandlers struct {
	service services.AccountService
}

func (h AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, err2 := h.service.NewAccount(request)
		if err2 != nil {
			WriteResponse(w, err2.Code, err2.Message)
		} else {
			WriteResponse(w, http.StatusCreated, account)
		}
	}
}