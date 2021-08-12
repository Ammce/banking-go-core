package app

import (
	"log"
	"net/http"

	"github.com/Ammce/go-banking-core/domain"
	"github.com/Ammce/go-banking-core/logger"
	"github.com/Ammce/go-banking-core/service"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	router := mux.NewRouter()

	dbClient := createDatabase()

	customerRepositoryDB := domain.NewCustomerREpositoryDB(dbClient)

	ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDB)}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func createDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=banking port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate()

	if err != nil {
		logger.Error("error connectin go the database")
		panic(err)
	}
	return db
}
