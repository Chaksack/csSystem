package models

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserBankDetails struct {
	gorm.Model
	UserID        uint   `json:"user_id"`
	User          *User  `gorm:"foreignKey:UserID" json:"user"`
	BankID        uint   `json:"bank_id"`
	Bank          *Bank  `gorm:"foreignKey:BankID" json:"bank"`
	AccountHolder string `json:"account_holder"`
	AccountNumber int    `json:"account_number"`
}

func (b *UserBankDetails) Scan(value interface{}) error {
	if value == nil {
		return errors.New("Scan: value is nil")
	}

	// Perform type assertion to extract the actual value
	if v, ok := value.(uint); ok {
		b.ID = (v)
		return nil
	}

	return fmt.Errorf("Scan: unable to convert value %v to Bank", value)
}

func (b UserBankDetails) Value() (driver.Value, error) {
	return b.ID, nil
}
