package controllers

import (
	"math"
	"strconv"

	"xactscore/database"
	"xactscore/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AllUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 10
	offset := (page - 1) * limit
	var total int64
	var users []models.User

	database.Database.Db.Preload("Role").Offset(offset).Limit(limit).Find(&users)
	database.Database.Db.Model(&models.User{}).Count(&total)
	result := database.Database.Db.Preload("BankDetails").Preload("BankDetails.Bank").Find(&users)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.JSON(fiber.Map{
		"data": users,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func CreateUser(c *fiber.Ctx) error {
	var newUser models.User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Create a new user in the database
	result := database.Database.Db.Create(&newUser)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.JSON(newUser)
}

// func GetUser(c *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(c.Params("id"))

//		user := models.User{
//			ID: uint(id),
//		}
//		database.Database.Db.Preload("Role").Find(&user)
//		return c.JSON(user)
//	}
func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var user models.User
	if result := database.Database.Db.Where("id = ?", id).Preload("Role").Preload("BankDetails").First(&user); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.JSON(user)
}

type UserUpdate struct {
	Email            string  `json:"email"`
	PhoneNumber      string  `json:"phone_number" gorm:"unique"`
	MonthlyIncome    float32 `json:"monthly_income"`
	EmploymentStatus string  `json:"employment_status"`
	MaritalStatus    string  `json:"marital_status"`
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var updatedUser UserUpdate
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Fetch the existing user details from the database
	var existingUser models.User
	if result := database.Database.Db.First(&existingUser, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	// Update the user details with the new data
	result := database.Database.Db.Model(&existingUser).Updates(updatedUser)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	// Return the updated user details
	return c.JSON(existingUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	result := database.Database.Db.Delete(&models.User{}, id)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
