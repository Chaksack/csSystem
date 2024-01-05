package models

type BusReg struct {
	ID              int    `json:"id"`
	RegNumber       string `json:"reg_number"`
	DocumentNo      string `json:"document_no"`
	PlaceOfIssuance string `json:"place_of_issurance"`
	ExpiryDate      string `json:"expiry_date"`
}

type BusTin struct {
	ID  int    `json:"id"`
	Tin string `json:"tin"`
}
