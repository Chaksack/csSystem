package controllers

import (
	"strconv"

	"xactscore/database"
	"xactscore/models"

	"github.com/gofiber/fiber/v2"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permissions

	database.Database.Db.Find(&permissions)

	return c.JSON(permissions)

}

func CreatePermission(c *fiber.Ctx) error {
	var permission models.Permissions

	if err := c.BodyParser(&permission); err != nil {
		return err
	}

	database.Database.Db.Create(&permission)

	return c.JSON(permission)

}

func GetPermission(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	Permissions := models.Permissions{
		Id: uint(id),
	}
	database.Database.Db.Find(&Permissions)
	return c.JSON(Permissions)
}

func UpdatePermission(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	permission := models.Permissions{
		Id: uint(id),
	}
	if err := c.BodyParser(&permission); err != nil {
		return err
	}
	database.Database.Db.Model(&permission).Updates(permission)
	return c.JSON(permission)
}

func DeletePermission(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	permission := models.Permissions{
		Id: uint(id),
	}

	database.Database.Db.Delete(&permission)
	return nil
}
