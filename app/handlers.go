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

func WriteJson(w http.ResponseWriter, i interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		return
	}
}
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "hello world!!")
}
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{Name: "John", City: "New York", Zipcode: "10001"},
	//	{Name: "John", City: "New York", Zipcode: "10001"},
	//	{Name: "John", City: "New York", Zipcode: "10001"},
	//}
	customers, _ := ch.service.GetAllCustomers()
	WriteJson(w, customers)

}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	fmt.Println(w, "customer_id:", customerId)
}

func createCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}
