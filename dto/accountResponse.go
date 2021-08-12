package dto

import (
	"github.com/google/uuid"
)

type AccountResponse struct {
	AccountId uuid.UUID `json:"account_id"`
}
