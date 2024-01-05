package controllers

import (
	// "crypto/rand"
	// "fmt"
	// "math/big"
	"strconv"
	"time"

	// "github.com/go-gomail/gomail"
	// "github.com/twilio/twilio-go"
	// twilioApi "github.com/twilio/twilio-go/rest/api/v2010"

	"xactscore/database"
	"xactscore/models"
	"xactscore/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	newUser := models.User{
		FirstName:   data["first_name"],
		LastName:    data["last_name"],
		Email:       data["email"],
		PhoneNumber: data["phone_number"],
		// OTPData: data[""],
		RoleId: 1,
	}

	newUser.SetPassword(data["password"])
	database.Database.Db.Create(&newUser)

	return c.JSON(newUser)
}
func StaffRegister(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	staff := models.Staff{
		Email:  data["email"],
		RoleId: 1,
	}

	staff.SetPassword(data["password"])
	database.Database.Db.Create(&staff)

	return c.JSON(staff)
}
func BusinessRegister(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	business := models.Business{
		CompanyName: data["company_name"],
		Email:       data["email"],
		PhoneNumber: data["phone_number"],
		RoleId:      1,
	}

	business.SetPassword(data["password"])
	database.Database.Db.Create(&business)

	return c.JSON(business)
}

// func Login(c *fiber.Ctx) error {
// 	var data map[string]string

// 	if err := c.BodyParser(&data); err != nil {
// 		return err
// 	}

// 	var user models.User

// 	database.Database.Db.Where("email = ?", data["email"]).First(&user)

// 	if user.Id == 0 {
// 		c.Status(fiber.StatusNotFound)
// 		return c.JSON(fiber.Map{
// 			"message": "user not found",
// 		})
// 	}

// 	if err := user.ComparePassword(data["password"]); err != nil {
// 		c.Status(fiber.StatusBadRequest)
// 		return c.JSON(fiber.Map{
// 			"message": "incorrect password",
// 		})
// 	}
// 	token, err := utils.GenerateJwt(strconv.Itoa(int(user.Id)))

// 	if err != nil {
// 		c.Status(fiber.StatusInternalServerError)
// 		return c.JSON(fiber.Map{
// 			"message": "could not login",
// 		})
// 	}

// 	cookie := fiber.Cookie{
// 		Name:     "jwt",
// 		Value:    token,
// 		Expires:  time.Now().Add(time.Hour * 24),
// 		HTTPOnly: true,
// 	}

// 	c.Cookie(&cookie)

// 	return c.JSON(fiber.Map{
// 		"message": "success",
// 	})

// }

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	if result := database.Database.Db.Where("email = ?", data["email"]).First(&user); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": "User not found",
			})
		}
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	token, err := utils.GenerateJwt(strconv.Itoa(int(user.ID)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func StaffLogin(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var staff models.Staff

	database.Database.Db.Where("email = ?", data["email"]).First(&staff)

	if staff.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "staff not found",
		})
	}

	if err := staff.ComparePassword(data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	token, err := utils.GenerateJwt(strconv.Itoa(int(staff.Id)))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func BusinessLogin(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var business models.Business

	database.Database.Db.Where("email = ?", data["email"]).First(&business)

	if business.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "business not found",
		})
	}

	if err := business.ComparePassword(data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	token, err := utils.GenerateJwt(strconv.Itoa(int(business.Id)))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := utils.ParseJwt(cookie)

	var user models.User

	database.Database.Db.Where("id = ?", id).First(&user)

	return c.JSON(user)
}

func Staff(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := utils.ParseJwt(cookie)

	var staff models.Staff

	database.Database.Db.Where("id = ?", id).First(&staff)

	return c.JSON(staff)
}

func Business(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := utils.ParseJwt(cookie)

	var business models.Business

	database.Database.Db.Where("id = ?", id).First(&business)

	return c.JSON(business)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// func UpdateInfo(c *fiber.Ctx) error {
// 	var data map[string]string

// 	if err := c.BodyParser(&data); err != nil {
// 		return err
// 	}

// 	cookie := c.Cookies("jwt")

// 	id, _ := utils.ParseJwt(cookie)

// 	userId, _ := strconv.Atoi(id)

// 	staffId, _ := strconv.Atoi(id)

// 	businessId, _ := strconv.Atoi(id)

// 	staff := models.Staff{
// 		Id:        uint(staffId),
// 		FirstName: data["first_name"],
// 		LastName:  data["last_name"],
// 		Email:     data["email"],
// 	}

// 	user := models.User{
// 		Id:        uint(userId),
// 		FirstName: data["first_name"],
// 		LastName:  data["last_name"],
// 		Email:     data["email"],
// 	}

// 	business := models.Business{
// 		Id:          uint(businessId),
// 		CompanyName: data["company_name"],
// 		PhoneNumber: data["phone_number"],
// 		Email:       data["email"],
// 	}

// database.Database.Db.Updates(user, staff, business)

// return c.JSON(user, staff, business)
// Update records individually
// 	if err := database.Database.Db.Model(&user).Updates(user).Error; err != nil {
// 		return err
// 	}

// 	if err := database.Database.Db.Model(&staff).Updates(staff).Error; err != nil {
// 		return err
// 	}

// 	if err := database.Database.Db.Model(&business).Updates(business).Error; err != nil {
// 		return err
// 	}

// 	return c.JSON(fiber.Map{
// 		"user":     user,
// 		"staff":    staff,
// 		"business": business,
// 	})
// }

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password does not match",
		})
	}

	cookie := c.Cookies("jwt")

	id, err := utils.ParseJwt(cookie)
	if err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	userId, _ := strconv.Atoi(id)

	// Fetch the user from the database
	var user models.User
	if err := database.Database.Db.First(&user, userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.Status(404)
			return c.JSON(fiber.Map{
				"message": "User not found",
			})
		}
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "Error fetching user",
		})
	}

	// Update the user's password
	user.SetPassword(data["password"])

	// Save the updated user record
	if err := database.Database.Db.Save(&user).Error; err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": "Error updating user",
		})
	}

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// func UpdatePassword(c *fiber.Ctx) error {
// 	var data map[string]string

// 	if err := c.BodyParser(&data); err != nil {
// 		return err
// 	}

// 	if data["password"] != data["password_confirm"] {
// 		c.Status(400)
// 		return c.JSON(fiber.Map{
// 			"message": "password does not match",
// 		})
// 	}

// 	cookie := c.Cookies("jwt")

// 	id, _ := utils.ParseJwt(cookie)
// 	userId, _ := strconv.Atoi(id)
// 	staffId, _ := strconv.Atoi(id)
// 	businessId, _ := strconv.Atoi(id)

// 	staff := models.Staff{
// 		Id: uint(staffId),
// 	}

// 	business := models.Business{
// 		Id: uint(businessId),
// 	}

// 	user := models.User{
// 		Id: uint(userId),
// 	}

// 	user.SetPassword(data["password"])
// 	staff.SetPassword(data["password"])
// 	business.SetPassword(data["password"])

// 	// Update records individually
// 	if err := database.Database.Db.Model(&user).Updates(user).Error; err != nil {
// 		return err
// 	}

// 	if err := database.Database.Db.Model(&staff).Updates(staff).Error; err != nil {
// 		return err
// 	}

// 	if err := database.Database.Db.Model(&business).Updates(business).Error; err != nil {
// 		return err
// 	}

// 	return c.JSON(fiber.Map{
// 		"user":     user,
// 		"staff":    staff,
// 		"business": business,
// 	})
// }
