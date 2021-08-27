package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"

	transactionDelivery "github.com/hezbymuhammad/payment-gateway/transaction/delivery/http"
	transactionRepo "github.com/hezbymuhammad/payment-gateway/transaction/repository/sqlite"
	transactionUsecase "github.com/hezbymuhammad/payment-gateway/transaction/usecase"

	merchantDelivery "github.com/hezbymuhammad/payment-gateway/merchant/delivery/http"
	merchantRepo "github.com/hezbymuhammad/payment-gateway/merchant/repository/sqlite"
	merchantUsecase "github.com/hezbymuhammad/payment-gateway/merchant/usecase"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	dbDriver := viper.GetString("database.driver")
	dbFile := viper.GetString("database.file")
	log.Println("database driver: " + dbDriver)
	log.Println("database file: " + dbFile)
	dbConn, err := sql.Open(dbDriver, dbFile)
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	mr := merchantRepo.NewMerchantRepository(dbConn)
	mu := merchantUsecase.NewMerchantUsecase(mr)
	tr := transactionRepo.NewTransactionRepository(dbConn)
	tu := transactionUsecase.NewTransactionUsecase(mr, tr)
	merchantDelivery.NewMerchantHandler(e, mu)
	transactionDelivery.NewTransactionHandler(e, tu)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
