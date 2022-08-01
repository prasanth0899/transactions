package schema

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
	"time"
)

type Transaction struct {
	ID                     string    `json:"id"`
	OrgId                  string    `json:"org_id"`
	CustomerId             string    `json:"customer_id"`
	FromAccount            string    `json:"from_account_number"`
	FromBankName           string    `json:"from_bank_name"`
	ToAccount              string    `json:"to_account_number"`
	ToBankName             string    `json:"to_bank_name"`
	TransactionType        string    `json:"debit/credit"`
	Amount                 string    `json:"amount"`
	ReferenceID            string    `json:"reference_id"`
	Status                 string    `json:"status"`
	TransactionMode        string    `json:"tx_type"`
	FromCountry            string    `json:"from_country"`
	ToCountry              string    `json:"to_country"`
	IsCurrencyConversion   bool      `json:"is_currency_conversion"`
	ActualConversionRate   string    `json:"actual_conversion_rate"`
	CustomerConversionRate string    `json:"customer_conversion_rate"`
	TransactionFees        string    `json:"transaction_fees"`
	MiscellaneousFees      string    `json:"miscellaneous_fees"`
	InputCurrency          string    `json:"input_currency"`
	OutputCurrency         string    `json:"output_currency"`
	Comments               string    `json:"comments"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

func CreateTransactionsTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Transaction{}, opts)
	if createError != nil {
		log.Printf("Error while creating transactions table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Transactions table created")
	return nil
}
