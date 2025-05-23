package main

import (
	"net/http"

	"github.com/likhithkp/banking-ledger-service/config"
	"github.com/likhithkp/banking-ledger-service/db/mongo"
	"github.com/likhithkp/banking-ledger-service/db/psql"
	"github.com/likhithkp/banking-ledger-service/routers"
	"github.com/likhithkp/banking-ledger-service/services"
)

func main() {
	config.Config()

	mongo.InitMongo()
	psql.GetDB()
	defer psql.CloseDB()

	go services.ConsumeTransaction("kafka:9092", "transactionGroup", "transactions")

	mux := http.NewServeMux()
	routers.Router(mux)
	http.ListenAndServe(":3001", mux)
}
