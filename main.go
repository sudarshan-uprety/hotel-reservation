package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudarshan-uprety/hotel-reservation/api"
	"github.com/sudarshan-uprety/hotel-reservation/db"
	"github.com/sudarshan-uprety/hotel-reservation/initializers"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	db.ConnectDatabase()
	initializers.SyncDatabase()

	// handle initialization
	userHandler := api.NewUserHandler(db.NewPostgresUserStore(db.DB))

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Put("/user/:id", userHandler.HandlePutUser)

	app.Listen(":8000")
}
