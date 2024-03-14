package app

import (
	"encoding/json"
	"github.com/Dubjay18/gobank2/dto"
	"github.com/Dubjay18/gobank2/service"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteJson(w, err.Error(), http.StatusBadRequest)

	} else {
		request.CustomerId = customerId
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			WriteJson(w, appError.AsMessage(), appError.Code)
		} else {
			WriteJson(w, account, http.StatusCreated)
		}
	}
}

func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	var request dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteJson(w, err.Error(), http.StatusBadRequest)
	} else {
		request.AccountId = accountId
		transaction, appError := h.service.MakeTransaction(request)
		if appError != nil {
			WriteJson(w, appError.AsMessage(), appError.Code)
		} else {
			WriteJson(w, transaction, http.StatusCreated)
		}
	}
}
