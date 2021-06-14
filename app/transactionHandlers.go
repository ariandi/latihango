package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"latihan1/dto"
	"latihan1/logger"
	"latihan1/services"
	"net/http"
)

type TransactionHandlers struct {
	service services.TransactionService
}

func (h TransactionHandlers) NewTransaction(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	var request dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responseHandler(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		trx, err2 := h.service.NewTransaction(request)
		if err2 != nil {
			logger.Error("error" + err2.Message)
			responseHandler(w, err2.Code, err2.AsMessage())
		} else {
			responseHandler(w, http.StatusCreated, trx)
		}
	}
}