package services

import (
	"time"
	"transactions/repositories"
	"transactions/repositories/schema"
)

func GetTransactions(orgId string, limit int, before *time.Time) *[]schema.Transaction {
	if before == nil {
		before = new(time.Time)
		*before = time.Now()
	}
	transactions := repositories.GetTransactionsFromDB(orgId, limit, *before)
	return transactions
}
