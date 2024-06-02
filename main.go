package main

import (
	"log"

	"github.com/amaldevm19/go_matrix_tna/database"
	router "github.com/amaldevm19/go_matrix_tna/router/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/handlebars/v2"
)

func GetHomePage(c *fiber.Ctx) error {
	// Render index
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}

func main() {
	engine := handlebars.New("./views", ".hbs")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	database.ConnectDB()
	cosec_db := database.COSEC_DB
	defer cosec_db.Close()

	proxy_db := database.TNA_PROXY_DB
	defer proxy_db.Close()

	app.Use(cors.New())
	app.Use(logger.New())

	api := app.Group("/api")
	router.SetupBranchRoutes(api)
	router.SetupDepartmentRoutes(api)

	app.Get("/", GetHomePage)
	log.Fatal(app.Listen(":8000"))

}
