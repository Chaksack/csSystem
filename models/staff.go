package models

import "golang.org/x/crypto/bcrypt"

type Staff struct {
	Id            uint        `json:"id"`
	FirstName     string      `json:"first_name"`
	LastName      string      `json:"last_name"`
	Email         string      `json:"email" gorm:"unique"`
	Password      []byte      `json:"-"`
	RoleId        uint        `json:"role_id"`
	Role          Role        `json:"role" gorm:"foreignKey:RoleId"`
	PermissionsId uint        `json:"permissions_id"`
	Permissions   Permissions `json:"permissions" gorm:"foreignKey:PermissionsId"`
}

// func (staff *Staff) SetPassword(password string) {
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("1234"), 14)
// 	staff.Password = hashedPassword
// }

// func (staff *Staff) ComparePassword(password string) error {
// 	return bcrypt.CompareHashAndPassword(staff.Password, []byte(password))
// }

func (staff *Staff) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	staff.Password = hashedPassword
	return nil
}

func (staff *Staff) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword(staff.Password, []byte(password))
	return err
}
