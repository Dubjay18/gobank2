package app

import (
	"encoding/json"
	"fmt"
	"github.com/Dubjay18/gobank2/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string `json:"fullname"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func WriteJson(w http.ResponseWriter, i interface{}, code ...int) {
	if code == nil {

		code = append(code, http.StatusOK)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code[0])
	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		return
	}
}
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "hello world!!")
}
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomers(status)
	if err != nil {
		WriteJson(w, err.AsMessage(), err.Code)
		return
	}
	WriteJson(w, customers)

}
func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	customer, err := ch.service.GetCustomer(customerId)
	if err != nil {
		WriteJson(w, err.AsMessage(), err.Code)
		return
	}
	WriteJson(w, customer)

}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	fmt.Println(w, "customer_id:", customerId)
}

func createCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}
