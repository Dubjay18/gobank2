package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string `json:"fullname"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "hello world!!")
}
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "John", City: "New York", Zipcode: "10001"},
		{Name: "John", City: "New York", Zipcode: "10001"},
		{Name: "John", City: "New York", Zipcode: "10001"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer_id := vars["customer_id"]
	fmt.Println(w, "customer_id:", customer_id)
}

func createCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}
