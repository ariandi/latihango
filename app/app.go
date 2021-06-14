package app

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"latihan1/domain"
	"latihan1/logger"
	"latihan1/services"
	"log"
	"net/http"
	"os"
	"time"
)

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	trxRepositoryDb := domain.NewTransactionRepositoryDb(dbClient)

	// ch := CustomerHandlers{services.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{services.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandlers{services.NewAccountService(accountRepositoryDb)}
	th := TransactionHandlers{services.NewTransactionService(trxRepositoryDb)}

	// Define Routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/account/{account_id:[0-9]+}/transaction", th.NewTransaction).Methods(http.MethodPost)

	// Define Server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	logger.Info("address : " + address)
	logger.Info("port : " + port)

	log.Fatal(http.ListenAndServe(address+":"+port, router))
}

func getDbClient() *sqlx.DB  {
	client, err := sqlx.Open("mysql", "root:433205ari@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
