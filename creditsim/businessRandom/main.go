package main

import (
	"encoding/csv"
	"strconv"

	// "strings"
	"math/rand"
	"time"

	// "encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
)

type Loans struct {
	ID                                 int           `json:"id"`
	LoanAmount                         string        `json:"loan_amount"`
	LoanStatus                         bool          `json:"loan_status"`
	LoanTerm                           time.Month    `json:"loan_term"`
	LoanPurpose                        string        `json:"loan_purpose"`
	NumberOfPreviousLoans              int           `json:"number_of_previous_loans"`
	NumberOfRepaidLoans                int           `json:"number_of_repaid_loans"`
	NumberOfDefaultedLoans             int           `json:"number_of_defaulted_loans"`
	NumberOfLatePayments               int           `json:"number_of_late_payment"`
	AverageLoanAmount                  string        `json:"average_loan_amount"`
	TotalAmountOwed                    string        `json:"total_amount_owed"`
	AverageInterestRateOfPreviousLoans float32       `json:"airopl"`
	TimeSinceLastLoan                  time.Duration `json:"tsll"`
}

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

type Bank struct {
	Id            uint   `json:"id"`
	BankName      string `json:"bank_name"`
	AccountHolder string `json:"account_holder"`
	AccountNumber string `json:"account_number"`
}

type Momo struct {
	Id           uint   `json:"id"`
	Network      string `json:"network"`
	MobileNumber string `json:"mobile_number"`
}

type Business struct {
	Id          uint   `json:"id"`
	CompanyName string `json:"company_name"`
	Email       string `json:"email" gorm:"unique"`
	Phone       string `json:"phone" gorm:"unique"`
	Address     string `json:"address"`
	Password    []byte `json:"-"`
	BusRegId    uint   `json:"busreg_id"`
	BusReg      BusReg `json:"busreg" gorm:"foreignKey:BusRegId"`
	BusTinId    uint   `json:"bustin_id"`
	BusTin      BusTin `json:"bustin" gorm:"foreignKey:BusTinId"`
	BankId      uint   `json:"bank_id"`
	Bank        Bank   `json:"bank" gorm:"foreignKey:BankId"`
	MomoId      uint   `json:"momo_id"`
	Momo        Momo   `json:"momo" gorm:"foreignKey:MomoId"`
	LoansId     uint   `json:"loans_id"`
	Loans       Loans  `json:"loans" gorm:"foreignKey:LoansId"`
}

func GenerateBusinessHandler(c *fiber.Ctx) error {
	business := GenerateBusiness()
	return c.JSON(business)
}

// func ExportToCSVHandler(c *fiber.Ctx) error {
//     users := GenerateUsers()

//     // Create and open a CSV file for writing
//     file, err := os.Create("user_data.csv")
//     if err != nil {
//         return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
//     }
//     defer file.Close()

//     // rest of the code...

//	    return c.SendFile("user_data.csv")
//	}
func main() {
	app := fiber.New()

	app.Get("/generate_business", GenerateBusinessHandler)
	app.Get("/export_to_csv", ExportToCSV)

	log.Fatal(app.Listen(":8080"))
}

// func ExportToCSVHandler(c *fiber.Ctx) error {
// 	return ExportToCSV(c)
// }

func GenerateBusiness() []Business {
	var businesses []Business
	domain := "@test.com"
	ghindex := "GHA"
	for i := 0; i < 100; i++ {
		var business Business
		err := faker.FakeData(&business)
		if err != nil {
			// Handle the error or skip this user
			continue
		}
		business.Phone = generateRandomPhoneNumber(10)
		business.Momo.MobileNumber = generateRandomPhoneNumber(10)
		business.Bank.AccountNumber = generateRandomAccountNumber(14)
		// business.Ghcard.Sex = getRandomGender()
		business.Bank.BankName = getRandomBankName()
		business.Momo.Network = getRandomMomoName()
		business.BusTin.Tin = generateRandomPhoneNumber(10)
		business.BusReg.RegNumber = ghindex + generateRandomPhoneNumber(10)
		business.BusReg.DocumentNo = generateRandomPhoneNumber(7)
		business.CompanyName = getRandomFirstName()
		business.Email = business.CompanyName + domain
		business.Bank.AccountHolder = business.CompanyName
		business.Address = getRandomAddress()
		// business.MonthlyIncome = getRandomIncome()
		// business.Ghcard.Nationality = getRandomNationality()
		// business.Ghcard.PlaceOfIssuance = getRandomIssurance()
		business.Loans.AverageLoanAmount = getRandomIncome()
		business.Loans.LoanAmount = getRandomIncome()
		business.Loans.TotalAmountOwed = getRandomIncome()

		businesses = append(businesses, business)
	}

	return businesses
}

// func getRandomGender() string {
// 	genders := []string{"Male", "Female"}
// 	return genders[rand.Intn(len(genders))]
// }

func getRandomBankName() string {
	bankNames := []string{"Stanbic Bank", "Standard Chartered", "Fidelity Bank", "EcoBank", "ABSA"}
	return bankNames[rand.Intn(len(bankNames))]
}

func getRandomAddress() string {
	address := []string{"Kasoa", "Weija", "East Legon", "Legon", "Tema"}
	return address[rand.Intn(len(address))]
}
func getRandomIssurance() string {
	issurance := []string{"Kasoa", "Weija", "East Legon", "Legon", "Tema"}
	return issurance[rand.Intn(len(issurance))]
}
func getRandomFirstName() string {
	firstNames := []string{"Company1", "Company2", "Company3", "Company4", "Company5", "Company6", "Company7", "Company8", "Company9", "Company10", "Company11", "Company12", "Company13", "Company14", "Company15", "Company16", "Company17", "Company18", "Company19", "Company20", "Company21", "Company22", "Company23", "Company24", "Company25", "Company26", "Company27"}
	return firstNames[rand.Intn(len(firstNames))]
}

func getRandomLastName() string {
	lastNames := []string{"James", "Robert", "John", "Michael", "David", "William", "Richard", "Joseph", "Thomas", "Christopher", "Charles", "Daniel", "Matthew", "Anthony", "Mark", "Donald", "Steven", "Andrew", "Paul", "Kenneth", "Mary", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara", "Susan"}
	return lastNames[rand.Intn(len(lastNames))]
}
func getRandomMomoName() string {
	momoNames := []string{"Mtn", "Vodafone", "AT"}
	return momoNames[rand.Intn(len(momoNames))]
}
func getRandomNationality() string {
	nationality := []string{"Ghanaian", "Nigerian"}
	return nationality[rand.Intn(len(nationality))]
}

func getRandomIncome() string {
	income := []string{"750", "900", "1000", "1500", "2000", "2500", "3000", "3500", "4000", "4500", "5000", "5500", "6000", "6500", "7000", "7500", "8000", "8500", "9000", "9500", "10000"}
	return income[rand.Intn(len(income))]
}

func generateRandomPhoneNumber(length int) string {
	rand.Seed(time.Now().UnixNano())

	var result string
	for i := 0; i < length; i++ {
		// Generate a random digit from 0 to 9
		digit := rand.Intn(10)
		result += fmt.Sprint(digit)
	}

	return result
}

func generateRandomAccountNumber(length int) string {
	rand.Seed(time.Now().UnixNano())

	var result string
	for i := 0; i < length; i++ {
		// Generate a random digit from 0 to 9
		digit := rand.Intn(14)
		result += fmt.Sprint(digit)
	}

	return result
}

func ExportToCSV(c *fiber.Ctx) error {
	businesses := GenerateBusiness()

	// for i := 0; i < 100; i++ {
	// 	var user User
	// 	err := faker.FakeData(&user)
	// 	if err != nil {
	// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	// 	}
	// 	users = append(users, user)
	// }

	// Create and open a CSV file for writing
	file, err := os.Create("business_data.csv")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the CSV header
	header := []string{
		"Id",
		"CompanyName",
		"Email",
		"Address",
		"Phone",
		"BusReg",
		"BusTin",
		"DocumentNo",
		"BankName",
		"AccountHolder",
		"AccountNumber",
		"Network",
		"MobileNumber",
		"LoanAmount",
		"LoanStatus",
		"LoanTerm",
		"LoanPurpose",
		"NumberofPreviousLoans",
		"NumberOfRepaidLoans",
		"NumberOfDefaultedLoans",
		"NumberOfLatePayments",
		"AverageLoanAmount",
		"TotalAmountOwed",
		"AverageInterestRateOfPreviousLoans",
		"TimeSinceLastLoan"}
	if err := writer.Write(header); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Write user data to the CSV file
	for _, business := range businesses {
		record := []string{
			fmt.Sprintf("%d", business.Id),
			business.CompanyName,
			business.Email,
			business.Address,
			business.Phone,
			business.BusReg.RegNumber,
			business.BusReg.DocumentNo,
			business.BusTin.Tin,
			business.Bank.BankName,
			business.Bank.AccountHolder,
			business.Bank.AccountNumber,
			business.Momo.Network,
			business.Momo.MobileNumber,
			business.Loans.LoanAmount,
			strconv.FormatBool(business.Loans.LoanStatus),
			fmt.Sprintf("%d", business.Loans.LoanTerm),
			business.Loans.LoanPurpose,
			fmt.Sprintf("%d", business.Loans.NumberOfPreviousLoans),
			fmt.Sprintf("%d", business.Loans.NumberOfRepaidLoans),
			fmt.Sprintf("%d", business.Loans.NumberOfDefaultedLoans),
			fmt.Sprintf("%d", business.Loans.NumberOfLatePayments),
			business.Loans.AverageLoanAmount,
			business.Loans.TotalAmountOwed,
			fmt.Sprintf("%f", business.Loans.AverageInterestRateOfPreviousLoans),
			fmt.Sprintf("%d", business.Loans.TimeSinceLastLoan)}
		if err := writer.Write(record); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	}

	c.Response().Header.Set("Content-Type", "text/csv")
	c.Response().Header.Set("Content-Disposition", "attachment; filename=user_data.csv")

	return c.SendFile("business_data.csv")
}
