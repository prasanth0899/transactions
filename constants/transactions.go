package constants

type Filter string

const (
	date        Filter = "date"
	amount             = "amount"
	method             = "method"
	sentFrom           = "sentFrom"
	status             = "status"
	attachments        = "attachments"
)

type Title string

const (
	fromDate   Title = "fromDate"
	beforeDate Title = "beforeDate"
	direction  Title = "direction"
)
