package customerDTO

import "github.com/Ammce/go-banking-core/errs"

type CreateCustomer struct {
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

func (c CreateCustomer) Validate() *errs.AppError {
	if len(c.Zipcode) > 7 {
		return errs.NewValidationError("Invalid length of the zipcode")
	}
	return nil
}
