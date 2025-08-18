// Package model provides the Recipient model struct
package models

import "fmt"

// Recipient holds details about an email recipient/vendor
type Recipient struct {
	Name           string
	Email          string
	Amount         float64
	Month          string
	TINBeneficiary string
	TINPayer       string
	SerialNo       string
	Date           string
}

func (r Recipient) String() string {
	return fmt.Sprintf(
		"Recipient{Name: %s, Email: %s, Amount: %f, Month: %s, TINBeneficiary: %s, TINPayer: %s, SerialNo: %s, Date: %s}",
		r.Name, r.Email, r.Amount, r.Month, r.TINBeneficiary, r.TINPayer, r.SerialNo, r.Date,
	)
}
