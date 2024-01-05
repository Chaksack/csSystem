package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName             string  `json:"first_name"`
	LastName              string  `json:"last_name"`
	Email                 string  `json:"email" gorm:"unique"`
	PhoneNumber           string  `json:"phone_number" gorm:"unique"`
	Age                   int     `json:"age"`
	MonthlyIncome         float32 `json:"monthly_income"`
	EmploymentStatus      string  `json:"employment_status"`
	EmploymentDuration    int     `json:"employment_duration"`
	NumberOfDependents    int     `json:"number_of_dep"`
	MaritalStatus         string  `json:"marital_status"`
	EducationalBackground string  `json:"educational_background"`
	HomeOwnershipStatus   string  `json:"home_status"`
	Password              []byte  `json:"-"`
	RoleId                uint
	Role                  Role       `json:"role" gorm:"foreignKey:RoleId"`
	GhcardId              uint       `json:"-"`
	Ghcard                Ghcard     `json:"ghcard" gorm:"foreignKey:GhcardId"`
	UserTinId             uint       `json:"-"`
	UserTin               UserTin    `json:"usertin" gorm:"foreignKey:UserTinId"`
	BankUsers             []BankUser `json:"bank_users"`
	Momo                  []Momo     `gorm:"many2many:user_momos"`
	Loan                  []Loans    `gorm:"many2many:user_loans"`
	CreditId              uint
	Credit                Credit `json:"credit" gorm:"foreignKey:CreditId"`
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return nil
}

func (user *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	return err
}
