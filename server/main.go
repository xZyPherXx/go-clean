package main

import (
	"fmt"
	"pwa/config"
	"pwa/internal/database"
	"pwa/internal/handler"
	"pwa/internal/repository"
	"pwa/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// type Product struct {
// 	ID          int     `json:"id"`
// 	Name        string  `json:"name"`
// 	Description string  `json:"description"`
// 	Price       float64 `json:"price"`
// 	ImageURL    string  `json:"image_url"`
// }

// var products = []Product{
// 	{ID: 1, Name: "Rainbow Six Siege", Description: "Tactical, FPS", Price: 14.00, ImageURL: "/images/rainbow-six-siege.png"},
// 	{ID: 2, Name: "Stardew Valley", Description: "Farming, Simulation", Price: 9.00, ImageURL: "/images/stardew-valley.png"},
// 	{ID: 3, Name: "theHunter: Call of the Wild™", Description: "Simulation, Open-world, Hunting", Price: 11.99, ImageURL: "/images/the-hunter.png"},
// 	{ID: 4, Name: "Minecraft", Description: "Sandbox, Open-world", Price: 29.00, ImageURL: "/images/minecraft.png"},
// 	{ID: 5, Name: "Sea of Thieves", Description: "Action, Adventure", Price: 38.00, ImageURL: "/images/sea-of-thieves.png"},
// 	{ID: 6, Name: "Ghost Recon® Breakpoint", Description: "Tactical, FPS, Action, Adventure", Price: 45.00, ImageURL: "/images/ghost-recon.png"},
// 	{ID: 7, Name: "Farm Together 2", Description: "Farming, Simulation", Price: 11.00, ImageURL: "/images/farm-together.png"},
// 	{ID: 8, Name: "Elden Ring", Description: "Action, RPG, Adventure", Price: 50.99, ImageURL: "/images/elden-ring.png"},
// }

func main() {

	config := config.LoadConfig()
	database := database.InitDatabase(config.Database)
	repository := repository.NewRepository(database)
	usecase := usecase.NewUsecase(repository)
	handler := handler.NewHandler(usecase)

	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: "*", AllowHeaders: "Origin, Content-Type, Accept"}))
	// app.Get("/api/products", func(c *fiber.Ctx) error {

	// 	return c.Status(fiber.StatusOK).JSON(products)
	// })
	app.Get("/api/products", handler.GetProducts)
	app.Get("/api/products/:id", handler.GetProductById)
	app.Post("/api/products", handler.CreateProduct)
	app.Delete("/api/products/:id", handler.DeleteProduct)
	app.Put("/api/products/:id", handler.UpdateProductById)

	// ! App defualt HTTP
	app.Listen(fmt.Sprintf(":%s", config.App.Port))
	// ! App HTTPS
	// err := app.ListenTLS(fmt.Sprintf(":%s", config.App.Port), config.App.FL, config.App.PK)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
