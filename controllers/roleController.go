package controllers

import (
	"strconv"

	"xactscore/database"
	"xactscore/models"

	"github.com/gofiber/fiber/v2"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.Database.Db.Find(&roles)

	return c.JSON(roles)

}

// func CreateRole(c *fiber.Ctx) error {
// 	var roleDTO fiber.Map

// 	if err := c.BodyParser(&roleDTO); err != nil {
// 		return err
// 	}

// 	// Check if 'permissions' is not nil
// 	permissionsInterface := roleDTO["permissions"]
// 	if permissionsInterface == nil {
// 		return c.Status(fiber.StatusBadRequest).SendString("Permissions data is missing")
// 	}

// 	// Now, you can safely assert it to a slice of interfaces
// 	list, ok := permissionsInterface.([]interface{})
// 	if !ok {
// 		return c.Status(fiber.StatusBadRequest).SendString("Invalid 'permissions' data format")
// 	}

// 	permissions := make([]models.Permission, len(list))

// 	for i, permissionId := range list {
// 		// id, err := strconv.Atoi(permissionId.(string))
// 		idStr, ok := permissionId.(string)
// 		if !ok {
// 			return c.Status(fiber.StatusBadRequest).SendString("Permission ID is not a string")
// 		}

// 		id, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			return c.Status(fiber.StatusBadRequest).SendString("Invalid permission ID format")
// 		}

// 		if err != nil {
// 			return c.Status(fiber.StatusBadRequest).SendString("Invalid permission ID")
// 		}

// 		permissions[i] = models.Permission{
// 			Id: uint(id),
// 		}
// 	}

// 	role := models.Role{
// 		Name:       roleDTO["name"].(string),
// 		Permission: permissions,
// 	}

// 	database.Database.Db.Create(&role)

// 	return c.JSON(role)
// }

func CreateRole(c *fiber.Ctx) error {
	var roleDTO fiber.Map

	if err := c.BodyParser(&roleDTO); err != nil {
		return err
	}

	// Check if 'permissions' is not nil, and if it is, initialize it as an empty slice
	permissionsInterface := roleDTO["permissions"]
	if permissionsInterface == nil {
		permissionsInterface = []interface{}{}
	}

	// Now, you can safely assert it to a slice of interfaces
	list, ok := permissionsInterface.([]interface{})
	if !ok {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid 'permissions' data format")
	}

	permissions := make([]models.Permissions, len(list))

	for i, permissionId := range list {
		idStr, ok := permissionId.(string)
		if !ok {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid permission ID format")
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid permission ID format")
		}
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Failed to get permission name")
		}

		permissions[i] = models.Permissions{
			Id: uint(id),
			// Name: string(name),
		}
	}

	role := models.Role{
		// ID:   roleDTO["id"].(uint),
		Name:        roleDTO["name"].(string),
		Permissions: permissions,
	}

	database.Database.Db.Create(&role)

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}
	database.Database.Db.Preload("Permissions").Find(&role)
	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDTO fiber.Map

	if err := c.BodyParser(&roleDTO); err != nil {
		return err
	}

	list := roleDTO["permissions"].([]interface{})

	permissions := make([]models.Permissions, len(list))
	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permissions{
			Id: uint(id),
		}
	}
	var result interface{}

	database.Database.Db.Table("role_permissions").Where("role_id", id).Delete(&result)

	role := models.Role{
		Id:          uint(id),
		Name:        roleDTO["name"].(string),
		Permissions: permissions,
	}

	database.Database.Db.Model(&role).Updates(role)
	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.Database.Db.Delete(&role)
	return nil
}
