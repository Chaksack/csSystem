package controllers

import (
	"math"
	"strconv"

	"xactscore/database"
	"xactscore/models"

	"github.com/gofiber/fiber/v2"
)

func AllStaff(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 10
	offset := (page - 1) * limit
	var total int64
	var staffs []models.Staff

	database.Database.Db.Preload("Role").Offset(offset).Limit(limit).Find(&staffs)
	database.Database.Db.Model(&models.Staff{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": staffs,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})

}

func CreateStaff(c *fiber.Ctx) error {
	var staff models.Staff

	if err := c.BodyParser(&staff); err != nil {
		return err
	}

	staff.SetPassword("1234")

	database.Database.Db.Create(&staff)

	return c.JSON(staff)

}

func GetStaff(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	staff := models.Staff{
		Id: uint(id),
	}
	database.Database.Db.Preload("Role").Find(&staff)
	return c.JSON(staff)
}

func UpdateStaff(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	staff := models.Staff{
		Id: uint(id),
	}
	if err := c.BodyParser(&staff); err != nil {
		return err
	}
	database.Database.Db.Model(&staff).Updates(staff)
	return c.JSON(staff)
}

func DeleteStaff(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	staff := models.Staff{
		Id: uint(id),
	}

	database.Database.Db.Delete(&staff)
	return nil
}
