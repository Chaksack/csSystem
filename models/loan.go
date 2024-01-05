package models

import "gorm.io/gorm"

type Loans struct {
	gorm.Model
	LoanAmount                         float32 `json:"loan_amount"`
	LoanStatus                         bool    `json:"loan_status"`
	LoanTerm                           int     `json:"loan_term"`
	LoanPurpose                        string  `json:"loan_purpose"`
	NumberOfPreviousLoans              int     `json:"number_of_previous_loans"`
	NumberOfRepaidLoans                int     `json:"number_of_repaid_loans"`
	NumberOfDefaultedLoans             int     `json:"number_of_defaulted_loans"`
	NumberOfLatePayments               int     `json:"number_of_late_payment"`
	AverageLoanAmount                  float32 `json:"average_loan_amount"`
	TotalAmountOwed                    float32 `json:"total_amount_owed"`
	AverageInterestRateOfPreviousLoans float32 `json:"airopl"`
	TimeSinceLastLoan                  int     `json:"tsll"`
}
