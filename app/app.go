package app

import (
	"github.com/Dubjay18/gobank2/domain"
	"github.com/Dubjay18/gobank2/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	//mux := http.NewServeMux()

	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	r := mux.NewRouter()
	r.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
