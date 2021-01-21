package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	models "github.com/samhj/AchmadGo/api/models"
	// router "github.com/samhj/AchmadGo/api/routes"
)

//AppMiddleWares ...
func AppMiddleWares(s *models.Server) {

	//setup headers
	Headers(s)
	
	//setup validators
	UserValidators(s)

	//cache some requests like site settings etc.
	s.App.Use("/api/settings/getAll", cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   20 * time.Minute,
		CacheControl: true,
	}))

	// s.App.Add("/api",router.SetupRoutes(sesrver))

	// s.App.Use("/api", func(c *fiber.Ctx) error{
	// 	return router.SetupRoutes(s)
	// })

	// s.App.Use("/", func(c *fiber.Ctx) error {
	// 	return c.Status(200).SendFile("./public/index.html")
	// })

}
