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

func AllMomos(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 10
	offset := (page - 1) * limit
	var total int64
	var momos []models.Momo

	database.Database.Db.Preload("User").Offset(offset).Limit(limit).Find(&momos)
	database.Database.Db.Model(&models.Momo{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": momos,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})

}

func CreateMomo(c *fiber.Ctx) error {
	var momo models.Momo

	if err := c.BodyParser(&momo); err != nil {
		return err
	}

	database.Database.Db.Create(&momo)

	return c.JSON(momo)
}

func CreateUserMomo(c *fiber.Ctx) error {
	var momo models.Momo

	// Parse the bank details from the request body
	if err := c.BodyParser(&momo); err != nil {
		return err
	}

	// Get the user ID from the request parameter or any other source
	// userIDParam := c.Params("id")
	// parsedUserID, err := strconv.Atoi(userIDParam)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	// }

	// Create the momo account for the user
	// momo.UserID = uint(parsedUserID)
	// if err := database.Database.Db.Create(&momo).Error; err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating momo account"})
	// }

	return c.JSON(momo)
}
func GetMomo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var momo models.Momo
	if err := database.Database.Db.Preload("User").First(&momo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Momo not found",
			})
		}
		// Handle other errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	return c.JSON(momo)
}

func UpdateMomo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var updatedMomo models.Momo
	if err := c.BodyParser(&updatedMomo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Fetch the existing bank details from the database
	var existingMomo models.Momo
	if result := database.Database.Db.First(&existingMomo, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Momo not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	// Update the bank details with the new data
	database.Database.Db.Model(&existingMomo).Updates(updatedMomo)

	return c.JSON(existingMomo)
}

func DeleteMomo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// Fetch the existing bank details from the database
	var existingMomo models.Momo
	if result := database.Database.Db.First(&existingMomo, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Momo not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	// Delete the bank record
	result := database.Database.Db.Delete(&existingMomo)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
