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

type Ghcard struct {
	ID              int           `json:"id"`
	FullName        string        `json:"full_name"`
	Nationality     string        `json:"nationality"`
	Dob             string        `json:"dob"`
	Sex             string        `json:"sex"`
	IdNumber        string        `json:"id_number"`
	DocumentNo      string        `json:"document_no"`
	Height          int           `json:"height"`
	PlaceOfIssuance string        `json:"place_of_issurance"`
	ExpiryDate      time.Duration `json:"expiry_date"`
}

type UserTin struct {
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

type User struct {
	Id                    uint       `json:"id"`
	FirstName             string     `json:"firstname"`
	LastName              string     `json:"last_name"`
	Name                  string     `json:"name"`
	Email                 string     `json:"email"`
	Age                   int        `json:"age"`
	Address               string     `json:"address"`
	Phone                 string     `json:"phone"`
	MonthlyIncome         string     `json:"monthly_income"`
	EmploymentStatus      bool       `json:"employment_status"`
	EmploymentDuration    time.Month `json:"employment_duration"`
	NumberOfDependents    int        `json:"number_of_dep"`
	MaritalStatus         bool       `json:"marital_status"`
	EducationalBackground string     `json:"educational_background"`
	HomeOwnershipStatus   bool       `json:"home_status"`
	GhcardId              uint       `json:"ghcard_id"`
	Ghcard                Ghcard     `json:"ghcard" gorm:"foreignKey:GhcardId"`
	UserTinId             uint       `json:"usertin_id"`
	UserTin               UserTin    `json:"usertin" gorm:"foreignKey:UserTinId"`
	BankId                uint       `json:"bank_id"`
	Bank                  Bank       `json:"bank" gorm:"foreignKey:BankId"`
	MomoId                uint       `json:"momo_id"`
	Momo                  Momo       `json:"momo" gorm:"foreignKey:MomoId"`
	LoansId               uint       `json:"loans_id"`
	Loans                 Loans      `json:"loans" gorm:"foreignKey:LoansId"`
}

func GenerateUsersHandler(c *fiber.Ctx) error {
	users := GenerateUsers()
	return c.JSON(users)
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

	app.Get("/generate_users", GenerateUsersHandler)
	app.Get("/export_to_csv", ExportToCSV)

	log.Fatal(app.Listen(":8080"))
}

// func ExportToCSVHandler(c *fiber.Ctx) error {
// 	return ExportToCSV(c)
// }

func GenerateUsers() []User {
	var users []User
	domain := "@test.com"
	ghindex := "GHA"
	for i := 0; i < 100; i++ {
		var user User
		err := faker.FakeData(&user)
		if err != nil {
			// Handle the error or skip this user
			continue
		}
		user.Phone = generateRandomPhoneNumber(10)
		user.Momo.MobileNumber = generateRandomPhoneNumber(10)
		user.Bank.AccountNumber = generateRandomAccountNumber(14)
		user.Ghcard.Sex = getRandomGender()
		user.Bank.BankName = getRandomBankName()
		user.Momo.Network = getRandomMomoName()
		user.UserTin.Tin = generateRandomPhoneNumber(10)
		user.Ghcard.IdNumber = ghindex + generateRandomPhoneNumber(10)
		user.Ghcard.DocumentNo = generateRandomPhoneNumber(7)
		user.FirstName = getRandomFirstName()
		user.LastName = getRandomLastName()
		user.Name = user.FirstName + " " + user.LastName
		user.Email = user.Name + domain
		user.Ghcard.FullName = user.Name
		user.Bank.AccountHolder = user.Name
		user.Address = getRandomAddress()
		user.MonthlyIncome = getRandomIncome()
		user.Ghcard.Nationality = getRandomNationality()
		user.Ghcard.PlaceOfIssuance = getRandomIssurance()
		user.Loans.AverageLoanAmount = getRandomIncome()
		user.Loans.LoanAmount = getRandomIncome()
		user.Loans.TotalAmountOwed = getRandomIncome()

		users = append(users, user)
	}

	return users
}
func getRandomGender() string {
	genders := []string{"Male", "Female"}
	return genders[rand.Intn(len(genders))]
}

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
	firstNames := []string{"James", "Robert", "John", "Michael", "David", "William", "Richard", "Joseph", "Thomas", "Christopher", "Charles", "Daniel", "Matthew", "Anthony", "Mark", "Donald", "Steven", "Andrew", "Paul", "Kenneth", "Mary", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara", "Susan"}
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
	users := GenerateUsers()

	// for i := 0; i < 100; i++ {
	// 	var user User
	// 	err := faker.FakeData(&user)
	// 	if err != nil {
	// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	// 	}
	// 	users = append(users, user)
	// }

	// Create and open a CSV file for writing
	file, err := os.Create("user_data.csv")
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
		"Name",
		"Email",
		"Age",
		"Address",
		"Phone",
		"MonthlyIncome",
		"EmploymentStatus",
		"EmploymentDuration",
		"NumberOfDependents",
		"MaritalStatus",
		"EducationalBackground",
		"HomeOwnershipStatus",
		"Nationality",
		"Dob",
		"Sex",
		"IdNumber",
		"DocumentNo",
		"Height",
		"PlaceOfIssurance",
		"Tin",
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
	for _, user := range users {
		record := []string{
			fmt.Sprintf("%d", user.Id),
			user.Name,
			user.Email,
			fmt.Sprintf("%d", user.Age),
			user.Address,
			user.Phone,
			user.MonthlyIncome,
			strconv.FormatBool(user.EmploymentStatus),
			fmt.Sprintf("%d", user.EmploymentDuration),
			fmt.Sprintf("%d", user.NumberOfDependents),
			strconv.FormatBool(user.MaritalStatus),
			user.EducationalBackground,
			strconv.FormatBool(user.HomeOwnershipStatus),
			user.Ghcard.Nationality,
			user.Ghcard.Dob,
			user.Ghcard.Sex,
			user.Ghcard.IdNumber,
			user.Ghcard.DocumentNo,
			fmt.Sprintf("%d", user.Ghcard.Height),
			user.Ghcard.PlaceOfIssuance,
			user.UserTin.Tin,
			user.Bank.BankName,
			user.Bank.AccountHolder,
			user.Bank.AccountNumber,
			user.Momo.Network,
			user.Momo.MobileNumber,
			user.Loans.LoanAmount,
			strconv.FormatBool(user.Loans.LoanStatus),
			fmt.Sprintf("%d", user.Loans.LoanTerm),
			user.Loans.LoanPurpose,
			fmt.Sprintf("%d", user.Loans.NumberOfPreviousLoans),
			fmt.Sprintf("%d", user.Loans.NumberOfRepaidLoans),
			fmt.Sprintf("%d", user.Loans.NumberOfDefaultedLoans),
			fmt.Sprintf("%d", user.Loans.NumberOfLatePayments),
			user.Loans.AverageLoanAmount,
			user.Loans.TotalAmountOwed,
			fmt.Sprintf("%f", user.Loans.AverageInterestRateOfPreviousLoans),
			fmt.Sprintf("%d", user.Loans.TimeSinceLastLoan)}
		if err := writer.Write(record); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	}

	c.Response().Header.Set("Content-Type", "text/csv")
	c.Response().Header.Set("Content-Disposition", "attachment; filename=user_data.csv")

	return c.SendFile("user_data.csv")
}
