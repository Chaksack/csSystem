package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Define an endpoint for calculating credit score
	app.Post("/calculate-credit-score", CalculateCreditScore)

	app.Listen(":3030")
}

func CalculateCreditScore(c *fiber.Ctx) error {
	// Parse input JSON
	var input struct {
		AccountCreation int
		KYC             bool
		AddAccount      bool
		AgeOfAccount    int
		Income          float64
		OnTimeCredit    float64
		LateCredit      float64
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Calculate credit score (a simple example)
	creditScore := CalculateScore(input.AccountCreation, input.KYC, input.AddAccount, input.AgeOfAccount, input.Income, input.OnTimeCredit, input.LateCredit)

	// Return the credit score as a response
	return c.JSON(fiber.Map{"credit_score": creditScore})
}

// Calculate credit score (a simple example)
func CalculateScore(accountCreation int, kyc bool, addAccount bool, ageOfAccount int, income float64, onTimeCredit float64, lateCredit float64) float64 {
	// You can define your own scoring algorithm here
	score := 0.0

	// Example: Add points based on various criteria
	if accountCreation >= 12 {
		score += 20.0
	}

	if kyc {
		score += 10.0
	}

	if addAccount {
		score += 5.0
	}

	if ageOfAccount >= 24 {
		score += 15.0
	}

	// Example: Income-based scoring
	if income >= 5000.0 {
		score += 10.0
	}

	// Example: Credit history scoring
	score += onTimeCredit - lateCredit

	// Ensure the score is within a certain range (e.g., 300-850)
	if score < 300 {
		score = 300
	} else if score > 850 {
		score = 950
	}

	return score
}
