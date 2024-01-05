package controllers

// import (
// 	"encoding/json"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"strings"
// 	"testing"
// 	"xactscore/database"
// 	"xactscore/models"

// 	// "xactscore/utils"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/assert"
// )

// func setupTestApp() *fiber.App {
// 	app := fiber.New()
// 	database.ConnectDb()

// 	// Add routes with your controllers here
// 	app.Post("/register", Register)
// 	app.Post("/staff/register", StaffRegister)
// 	app.Post("/business/register", BusinessRegister)
// 	app.Post("/login", Login)
// 	app.Post("/staff/login", StaffLogin)
// 	app.Post("/business/login", BusinessLogin)
// 	app.Get("/user", User)
// 	app.Get("/staff", Staff)
// 	app.Get("/business", Business)
// 	app.Post("/logout", Logout)
// 	app.Put("/update/info", UpdateInfo)
// 	app.Put("/update/password", UpdatePassword)

// 	return app
// }

// func performRequest(app *fiber.App, method, path string, body string) (*http.Response, string) {
// 	req := httptest.NewRequest(method, path, strings.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	resp, err := app.Test(req)
// 	if err != nil {
// 		panic(err)
// 	}

// 	bodyBytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return resp, string(bodyBytes)
// }

// func TestRegister(t *testing.T) {
// 	app := setupTestApp()
// 	defer database.Database.Db.Migrator().DropTable(&models.User{})
// 	defer database.Database.Db.Migrator().DropTable(&models.Staff{})
// 	defer database.Database.Db.Migrator().DropTable(&models.Business{})

// 	payload := `{"first_name": "John", "last_name": "Doe", "email": "john.doe@example.com", "phone_number": "1234567890", "password": "password123"}`
// 	resp, body := performRequest(app, "POST", "/register", payload)

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// 	assert.Contains(t, body, "John")
// 	assert.Contains(t, body, "Doe")
// 	// Add more assertions based on your expectations
// }

// // Similar tests for StaffRegister and BusinessRegister can be added here.

// // Add tests for Login, StaffLogin, and BusinessLogin

// func TestUpdateInfo(t *testing.T) {
// 	app := setupTestApp()
// 	defer database.Database.Db.Migrator().DropTable(&models.User{})
// 	defer database.Database.Db.Migrator().DropTable(&models.Staff{})
// 	defer database.Database.Db.Migrator().DropTable(&models.Business{})

// 	// Register a user to update info
// 	registerPayload := `{"first_name": "John", "last_name": "Doe", "email": "john.doe@example.com", "phone_number": "1234567890", "password": "password123"}`
// 	resp, _ := performRequest(app, "POST", "/register", registerPayload)
// 	var registeredUser models.User
// 	err := json.Unmarshal([]byte(resp.Body), &registeredUser)
// 	assert.NoError(t, err)

// 	// Login to get a JWT token
// 	loginPayload := `{"email": "john.doe@example.com", "password": "password123"}`
// 	resp, body := performRequest(app, "POST", "/login", loginPayload)
// 	var loginResponse map[string]string
// 	err = json.Unmarshal([]byte(body), &loginResponse)
// 	assert.NoError(t, err)
// 	jwtToken := loginResponse["message"]

// 	// Update info using the JWT token
// 	updateInfoPayload := `{"first_name": "Updated", "last_name": "Info", "email": "updated.email@example.com"}`
// 	resp, body = performRequest(app, "PUT", "/update/info", updateInfoPayload)
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)

// 	var updatedUser models.User
// 	err = json.Unmarshal([]byte(body), &updatedUser)
// 	assert.NoError(t, err)

// 	assert.Equal(t, registeredUser.ID, updatedUser.ID)
// 	assert.Equal(t, "Updated", updatedUser.FirstName)
// 	assert.Equal(t, "Info", updatedUser.LastName)
// 	assert.Equal(t, "updated.email@example.com", updatedUser.Email)
// }

// // Similar tests for Staff and Business controllers can be added here.

// // Add tests for UpdatePassword, User, Staff, Business, and Logout

// func TestMain(m *testing.M) {
// 	// Run the tests with exit code
// 	os.Exit(m.Run())
// }
