package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	models "github.com/samhj/AchmadGo/api/models"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

//Headers ...
func Headers(s *models.Server) {

	s.App.Use(cors.New())

	//used for recovering from any panics thrown
	s.App.Use(recover.New())

	s.App.Use(func(c *fiber.Ctx) error{
		c.Set("Content-Type", "application/json")
		return c.Next()
	})

}
