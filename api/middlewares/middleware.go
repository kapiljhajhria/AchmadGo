package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"

	// "github.com/gofiber/websocket/v2"
	models "github.com/samhj/AchmadGo/api/models"
	// router "github.com/samhj/AchmadGo/api/routes"
)

//AppMiddleWares ...
func AppMiddleWares(s *models.Server) {

	//setup headers
	Headers(s)

	//setup validators
	UserValidators(s)

	// Setup the middleware to retrieve the data sent in first GET request
	// s.App.Use(func(c *fiber.Ctx) error {
	// 	// IsWebSocketUpgrade returns true if the client
	// 	// requested upgrade to the WebSocket protocol.
	// 	if websocket.IsWebSocketUpgrade(c) {
	// 		c.Locals("allowed", true)
	// 		return c.Next()
	// 	}
	// 	return fiber.ErrUpgradeRequired
	// })

	//cache some requests like site settings etc.
	s.App.Use("/api/sitesettings/getAll", cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   10 * time.Minute,
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
