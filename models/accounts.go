package models

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	// UserID        uint
	// User          User   `gorm:"foreignKey:UserID"`
	BankName string `json:"bank_name"`
}
type BankUser struct {
	gorm.Model
	UserID uint `json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"user"`
	BankID uint `json:"bank_id"`
	// Bank          Bank   `gorm:"foreignKey:BankID" json:"bank"`
	AccountHolder string `json:"account_holder"`
	AccountNumber int    `json:"account_number"`
}

type Momo struct {
	gorm.Model
	// UserID       uint   `json:"user_id"`
	Network      string `json:"network"`
	MobileNumber int    `json:"mobile_number"`
}

func (b Bank) Value() (driver.Value, error) {
	return b.ID, nil
}

// Implement sql.Scanner interface
func (b *Bank) Scan(value interface{}) error {
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
