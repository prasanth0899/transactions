package repositories

import (
	"log"
	"time"
	"transactions/configs"
	"transactions/repositories/schema"
)

func GetTransactionsFromDB(orgId string, limit int, before time.Time) *[]schema.Transaction {
	var transactions *[]schema.Transaction
	db := configs.Connect()
	defer db.Close()
	err := db.Model(transactions).
		Where("org_id = ?", orgId).
		Where("created_at < ?", before).
		Order("created_at DSC").
		Limit(limit).Select()
	if err != nil {
		log.Printf("error while fetching data from postgres")
	}
	return transactions
}
