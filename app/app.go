package app

import (
	"fmt"
	"github.com/Dubjay18/gobank2/domain"
	"github.com/Dubjay18/gobank2/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func GetEnvVar() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

}

func SanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variables not defined...")
	}
}

func Start() {
	//mux := http.NewServeMux()

	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	r := mux.NewRouter()
	r.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	port := os.Getenv("SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), r))
}
