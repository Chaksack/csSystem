package models

import (
	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	Details  []UserBankDetails `json:"details"`
	BankName string            `json:"bank_name"`
}

type Momo struct {
	gorm.Model
	// UserID       uint   `json:"user_id"`
	Network      string `json:"network"`
	MobileNumber int    `json:"mobile_number"`
}

// func (b Bank) Value() (driver.Value, error) {
// 	return b.ID, nil
// }

// // Implement sql.Scanner interface
// func (b *Bank) Scan(value interface{}) error {
// 	if value == nil {
// 		return errors.New("Scan: value is nil")
// 	}

// 	// Perform type assertion to extract the actual value
// 	if v, ok := value.(uint); ok {
// 		b.ID = (v)
// 		return nil
// 	}

// 	return fmt.Errorf("Scan: unable to convert value %v to Bank", value)
// }
