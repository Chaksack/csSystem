package models

import "golang.org/x/crypto/bcrypt"

type Business struct {
	Id          uint   `json:"id"`
	CompanyName string `json:"company_name"`
	Email       string `json:"email" gorm:"unique"`
	PhoneNumber string `json:"phone_number" gorm:"unique"`
	Password    []byte `json:"-"`
	RoleId      uint   `json:"role_id"`
	Role        Role   `json:"role" gorm:"foreignKey:RoleId"`
	BusRegId    uint   `json:"busreg_id"`
	BusReg      BusReg `json:"busreg" gorm:"foreignKey:BusRegId"`
	BusTinId    uint   `json:"bustin_id"`
	BusTin      BusTin `json:"bustin" gorm:"foreignKey:BusTinId"`
	BankId      uint   `json:"bank_id"`
	Bank        Bank   `json:"bank" gorm:"foreignKey:BankId"`
	MomoId      uint   `json:"momo_id"`
	Momo        Momo   `json:"momo" gorm:"foreignKey:MomoId"`
}

func (business *Business) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("1234"), 14)
	business.Password = hashedPassword
}

func (business *Business) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(business.Password, []byte(password))
}
