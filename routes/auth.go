package routes

import (
	"gofiber-auth/handlers"
	"github.com/gofiber/fiber/v2"
	"gofiber-auth/middleware"
)

func Setup(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)
	auth.Post("/forgot", handlers.ForgotPassword)
	auth.Post("/reset", handlers.ResetPassword)
	auth.Get("/message", handlers.GetMessage)

	app.Get("/protected-route", middleware.Protected(), func(c *fiber.Ctx) error {
		userEmail := c.Locals("userEmail").(string)
		return c.JSON(fiber.Map{
			"message": "Access granted to protected route!",
			"email":   userEmail,
		})
	})

	
}
