package models

type Filter string

type Title string

type TransactionFilter struct {
	filter Filter
	title  Title
}
