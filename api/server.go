package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samhj/AchmadGo/api/config"
	"github.com/samhj/AchmadGo/api/config/db"
	middlewares "github.com/samhj/AchmadGo/api/middlewares"

	models "github.com/samhj/AchmadGo/api/models"
	router "github.com/samhj/AchmadGo/api/routes"
)

//Init ...
func Init() {

	s := models.Server{
		DB:  db.Client(),
		App: fiber.New(),
	}

	s.App.Static("/", "./public")

	//Setup app wide middlewares
	middlewares.AppMiddleWares(&s)

	//setup app wide routes
	router.SetupRoutes(&s)

	//close mongodb connection when the main function terminates
	defer db.DeferDB(&s)

	//start app
	err := s.App.Listen(":" + config.Config("PORT"))
	if err != nil {
		panic("panic error: " + err.Error())
	}

}
