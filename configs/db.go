package configs

import (
	"github.com/go-pg/pg"
	"github.com/spf13/viper"
	"log"
	"os"
	"transactions/repositories/schema"
)

func Connect() *pg.DB {
	viper.ReadInConfig()
	opts := &pg.Options{
		User:     viper.GetString("PG_USER"),
		Password: viper.GetString("PG_PASS"),
		Addr:     viper.GetString("PG_ADDR"),
		Database: viper.GetString("PG_DATABASE"),
	}
	var db = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	return db
}

func CreateTables() {
	db := Connect()
	if schema.CreateTransactionsTable(db) != nil {
		os.Exit(100)
	}
}
