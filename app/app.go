package app

import (
	"fmt"
	"github.com/Dubjay18/gobank2/domain"
	"github.com/Dubjay18/gobank2/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
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

func getDbClient() *sqlx.DB {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	constr := "user=" + dbUser + " dbname=" + dbName + " password=" + dbPass + " host=" + dbHost + " port=" + dbPort + " sslmode=disable"

	db, err := sqlx.Open("postgres", constr)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)

	}
	return db
}

func Start() {
	//mux := http.NewServeMux()
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDB(dbClient)
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}
	r := mux.NewRouter()
	r.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	r.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	port := os.Getenv("SERVER_PORT")
	address := os.Getenv("SERVER_ADDRESS")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), r))
}
