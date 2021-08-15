package customerDTO

import (
	"database/sql"
	"time"
)

type CustomerResponse struct {
	ID          uint         `json:"id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
	Name        string       `json:"name"`
	City        string       `json:"city"`
	Zipcode     string       `json:"zipcode"`
	DateofBirth string       `json:"date_of_birth"`
	Status      string       `json:"status"`
}
