package dtos

type MidtransInput struct {
	CustomerAddress    CustomerAddress    `json"customer_address"`
	TransactionDetails TransactionDetails `json:"transaction_details"`
	CustomerDetail     CustomerDetail     `json:"customer_detail"`
	Items              Items              `json:"items"`
}

type CustomerAddress struct {
	FName       string
	LName       string
	Phone       string
	Address     string
	City        string
	Postcode    string
	CountryCode string
}

type TransactionDetails struct {
	OrderID  string
	GrossAmt int
}

type CustomerDetail struct {
	FName string
	LName string
	Email string
	Phone string
}

type Items struct {
	ID    int
	Price int
	Qty   int
	Name  string
}
