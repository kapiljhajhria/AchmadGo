package router

import (
	controllers "github.com/samhj/AchmadGo/api/controllers"
	models "github.com/samhj/AchmadGo/api/models"
)

// SetupRoutes ...
func SetupRoutes(s *models.Server) {

	//Setup all controller routes in this project here
	controllers.UserController(s)
	controllers.SettingsController(s)

}
