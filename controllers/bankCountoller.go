package controllers

import (
	"errors"
	"math"
	"strconv"

	"xactscore/database"
	"xactscore/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AllBanks(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 10
	// offset := (page - 1) * limit
	var total int64
	var banks []models.Bank
	// database.Database.Db.Preload("User").Offset(offset).Limit(limit).Find(&banks)
	database.Database.Db.Model(&models.Bank{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": banks,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func CreateBank(c *fiber.Ctx) error {
	var bank models.Bank
	if err := c.BodyParser(&bank); err != nil {
		return err
	}
	if err := database.Database.Db.Create(&bank).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating bank account"})
	}
	return c.JSON(bank)
}

func CreateUserBank(c *fiber.Ctx) error {
	var bank models.BankUser
	if err := c.BodyParser(&bank); err != nil {
		return err
	}
	userIDParam := c.Params("id")
	parsedUserID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	var existingUser models.User
	if result := database.Database.Db.First(&existingUser, parsedUserID); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}
	bank.UserID = uint(parsedUserID)
	if err := database.Database.Db.Create(&bank).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating bank account"})
	}
	// if err := database.Database.Db.Model(&existingUser).Association("Bank").Append(&bank); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error associating bank with user"})
	// }
	return c.JSON(bank)
}
func GetUserBankDetails(c *fiber.Ctx) error {
	// Extract user ID from the request parameters
	userIDParam := c.Params("id")
	parsedUserID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Fetch user details including associated banks
	var user models.User
	if result := database.Database.Db.Preload("BankUsers").Preload("BankUsers.Bank").First(&user, parsedUserID); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	// Return the user's bank details
	return c.JSON(user.BankUsers)
}

// func AllBanks(c *fiber.Ctx) error {
// 	page, _ := strconv.Atoi(c.Query("page", "1"))
// 	limit := 10
// 	offset := (page - 1) * limit
// 	var total int64
// 	var bank []models.Bank

// 	database.Database.Db.Preload("Bank").Offset(offset).Limit(limit).Find(&bank)
// 	database.Database.Db.Model(&models.Bank{}).Count(&total)

// 	return c.JSON(fiber.Map{
// 		"data": bank,
// 		"meta": fiber.Map{
// 			"total":     total,
// 			"page":      page,
// 			"last_page": math.Ceil(float64(int(total) / limit)),
// 		},
// 	})

// }

// func CreateBank(c *fiber.Ctx) error {
// 	var bank models.Bank

// 	if err := c.BodyParser(&bank); err != nil {
// 		return err
// 	}

// 	database.Database.Db.Create(&bank)

// 	return c.JSON(bank)

// }

// func CreateUserBank(c *fiber.Ctx) error {
// 	var bank models.Bank

// 	// Parse the bank details from the request body
// 	if err := c.BodyParser(&bank); err != nil {
// 		return err
// 	}

// 	// Get the user ID from the request parameter or any other source
// 	userIDParam := c.Params("id")
// 	parsedUserID, err := strconv.Atoi(userIDParam)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
// 	}

// 	// Check if the user exists
// 	var existingUser models.User
// 	if result := database.Database.Db.First(&existingUser, parsedUserID); result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
// 	}

// 	// Associate the bank with the user without fetching the full user details
// 	bank.UserID = uint(parsedUserID)

// 	// Save the bank account without updating the user
// 	if err := database.Database.Db.Create(&bank).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating bank account"})
// 	}

// 	// Add the bank to the user's Banks slice
// 	existingUser.Bank = append(existingUser.Bank, bank)

// 	// Save the user with the updated Banks slice
// 	if err := database.Database.Db.Save(&existingUser).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error updating user's Banks slice"})
// 	}

// 	return c.JSON(bank)
// }

func DeleteBank(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// Fetch the existing bank details from the database
	var bank models.Bank
	if result := database.Database.Db.First(&bank, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Bank not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	// Delete the bank record
	result := database.Database.Db.Delete(&bank)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
