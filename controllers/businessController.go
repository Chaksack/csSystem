package controllers

import (
	"math"
	"strconv"

	"xactscore/database"
	"xactscore/models"

	"github.com/gofiber/fiber/v2"
)

func AllBusiness(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 10
	offset := (page - 1) * limit
	var total int64
	var business []models.Business

	database.Database.Db.Preload("Role").Offset(offset).Limit(limit).Find(&business)
	database.Database.Db.Model(&models.Business{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": business,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})

}

func CreateBusiness(c *fiber.Ctx) error {
	var business models.Business

	if err := c.BodyParser(&business); err != nil {
		return err
	}

	business.SetPassword("1234")

	database.Database.Db.Create(&business)

	return c.JSON(business)

}

func GetBusiness(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	business := models.Business{
		Id: uint(id),
	}
	database.Database.Db.Preload("Role").Find(&business)
	return c.JSON(business)
}

func UpdateBusiness(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	business := models.Business{
		Id: uint(id),
	}
	if err := c.BodyParser(&business); err != nil {
		return err
	}
	database.Database.Db.Model(&business).Updates(business)
	return c.JSON(business)
}

func DeleteBusiness(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	business := models.Business{
		Id: uint(id),
	}

	database.Database.Db.Delete(&business)
	return nil
}
