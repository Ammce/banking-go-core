package accountDTO

import (
	"database/sql"
	"time"
)

type CreateAccountResponse struct {
	ID          uint         `json:"id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
	OpeningDate string       `json:"opening_date"`
	AccountType string       `json:"account_type"`
	Amount      float64      `json:"amount"`
	Status      string       `json:"status"`
	CustomerID  uint         `json:"customer_id"`
}
