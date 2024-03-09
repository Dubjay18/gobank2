package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	//mux := http.NewServeMux()
	r := mux.NewRouter()
	r.HandleFunc("/greet", greet).Methods(http.MethodGet)
	r.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers", createCustomers).Methods(http.MethodPost)
	r.HandleFunc("/customers/{customer_id}", getCustomer).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
