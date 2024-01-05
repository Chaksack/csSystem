package controllers

import (
	"strconv"
	"xactscore/database"
	"xactscore/models"

	"github.com/gofiber/fiber/v2"
)

func UpdateUserTin(c *fiber.Ctx) error {
	// Parse the request parameter for the user ID
	userIDParam := c.Params("id")
	parsedUserID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var userTinUpdate models.UserTin
	if err := c.BodyParser(&userTinUpdate); err != nil {
		return err
	}

	var user models.User
	if err := database.Database.Db.First(&user, parsedUserID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting user"})
	}

	var existingGhcard models.Ghcard
	if err := database.Database.Db.First(&existingGhcard, user.GhcardId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting ghcard"})
	}

	var existingUserTin models.UserTin
	if err := database.Database.Db.First(&existingUserTin, user.UserTinId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting usertin"})
	}

	// Update the existing ghcard and usertin records with the new data
	// database.Database.Db.Model(&existingGhcard).Updates(userTinUpdate.Ghcard)
	// database.Database.Db.Model(&existingUserTin).Updates(userTinUpdate.UserTin)

	return c.JSON(fiber.Map{"success": true})
}

func AddUserKyc(c *fiber.Ctx) error {
	var ghcard models.Ghcard

	// Parse the bank details from the request body
	if err := c.BodyParser(&ghcard); err != nil {
		return err
	}

	// Get the user ID from the request parameter or any other source
	userIDParam := c.Params("id")
	parsedUserID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Create the ghcard account for the user
	ghcard.UserID = uint(parsedUserID)
	if err := database.Database.Db.Create(&ghcard).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating ghcard account"})
	}

	return c.JSON(ghcard)
}

func AddUserTin(c *fiber.Ctx) error {
	var tin models.UserTin

	// Parse the bank details from the request body
	if err := c.BodyParser(&tin); err != nil {
		return err
	}

	// Get the user ID from the request parameter or any other source
	userIDParam := c.Params("id")
	parsedUserID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Create the tin account for the user
	tin.UserID = uint(parsedUserID)
	if err := database.Database.Db.Create(&tin).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating tin account"})
	}

	return c.JSON(tin)
}
