package app

import (
	"log"
	"net/http"

	customerAdapter "github.com/Ammce/go-banking-core/adapters/Customer"
	"github.com/Ammce/go-banking-core/domain"
	customerPort "github.com/Ammce/go-banking-core/domain/Customer"
	"github.com/Ammce/go-banking-core/logger"
	"github.com/Ammce/go-banking-core/service"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	router := mux.NewRouter()

	dbClient := createDatabase()

	customerRepositoryDB := customerAdapter.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountingRepositoryDB(dbClient)
	transactionRepositoryDB := domain.NewTransactionRepositoryDB(dbClient)

	ch := CustomerHandlers{service: customerPort.NewCustomerService(customerRepositoryDB)}
	ah := AccountHandlers{service: service.NewAccountService(accountRepositoryDB)}
	th := TransactionHandlers{service: service.NewTransactionService(transactionRepositoryDB)}

	router.HandleFunc("/transactions", th.createTransaction).Methods(http.MethodPost)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.deleteCustomerById).Methods(http.MethodDelete)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.updateCustomer).Methods(http.MethodPatch)
	router.HandleFunc("/customers/{id:[0-9]+}/account", ah.createAccount).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func createDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=banking port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&domain.Customer{})
	// db.AutoMigrate(&domain.Customer{}, &domain.Account{}, domain.Transaction{})

	if err != nil {
		logger.Error("error connectin go the database")
		panic(err)
	}
	return db
}
