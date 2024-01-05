package main

import (
	// "embed"
	"log"
	"xactscore/database"
	"xactscore/routes"

	"github.com/gofiber/fiber/v2/middleware/cors"
	// "html/template"
	"github.com/gofiber/template/html/v2"

	"github.com/gofiber/fiber/v2"
	// "io/fs"
)

// var content embed.FS

func main() {
	database.ConnectDb()
	app := fiber.New()
	//prod version
	engine := html.New("./home/ubuntu/views", ".html")
	//test version
	// engine := html.New("./views", ".html")

	app = fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static files
	//prod version
	app.Static("/assets", "./home/ubuntu/assets")
	//test version
	// app.Static("/assets", "./assets")

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))

	routes.Setup(app)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Welcome to xactscore")
	// })

	log.Fatal(app.Listen(":3000"))
}

// func init() {
// 	engine := template.Must(template.New("").ParseGlob("views/*.html"))
// 	fiber.Views(engine)
// }
