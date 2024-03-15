package structs

import "time"

// Customer struct represents a customer
type Customer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	IDCardNumber string `json:"id_card_number"`
	Address      string `json:"address"`
	CIF          string `json:"cif"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// User struct represents a admin/user
type UserAdmin struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Saldo struct represents a saldo entry
type Saldo struct {
	ID              int     `json:"id"`
	CustomerID      int     `json:"customer_id"`
	Saldo           float64 `json:"saldo"`
	TransactionDate time.Time
}

// Transaction struct represents a transaction
type Transaction struct {
	ID              int     `json:"id"`
	SaldoID         int     `json:"saldo_id"`
	Volume          float64 `json:"volume"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate time.Time
}

// Report struct represents a report
type Report struct {
	ID              int     `json:"id"`
	TransactionID   int     `json:"transaction_id"`
	VolumeIn        float64 `json:"volume_in"`
	VolumeOut       float64 `json:"volume_out"`
	TransactionDate time.Time
}
