package models

import "gorm.io/gorm"

type Ghcard struct {
	gorm.Model
	UserID          uint   `json:"user_id"`
	FullName        string `json:"full_name"`
	Nationality     string `json:"nationality"`
	Dob             string `json:"dob"`
	Sex             string `json:"sex"`
	IdNumber        string `json:"id_number"`
	DocumentNo      string `json:"document_no"`
	Height          int    `json:"height"`
	PlaceOfIssuance string `json:"place_of_issurance"`
	ExpiryDate      string `json:"expiry_date"`
}

type UserTin struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Tin    string `json:"tin"`
}
