package controllers //declare package name here

import (
	"github.com/samhj/AchmadGo/api/config"

	"github.com/gofiber/fiber/v2"
	"github.com/samhj/AchmadGo/api/config/db"
	"github.com/samhj/AchmadGo/api/models"
	"github.com/samhj/AchmadGo/api/services"
)

//SettingsController for settings related operations
func SettingsController(s *models.Server) {

	resp := &models.Response{
		StatusCd: 400,
		Succ:     false,
		Data:     nil,
		Msg:      "",
	}

	collection := db.GetCollection(s, config.Config("SETTINGS"))
	usersColl := db.GetCollection(s, config.Config("USERS"))

	//handle sign-up
	s.App.Get(config.GetAPIBase()+"sitesettings/getAll",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			s.UsersColl = usersColl
			return services.GetSiteSettings(s)
		})

	s.App.Patch(config.GetAPIBase()+"sitesettings/update",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			s.UsersColl = usersColl
			return services.UpdateSiteSettings(s)

		})

}
