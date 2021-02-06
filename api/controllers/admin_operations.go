package controllers //declare package name here

import (
	"github.com/samhj/AchmadGo/api/config"

	"github.com/gofiber/fiber/v2"
	"github.com/samhj/AchmadGo/api/config/db"
	"github.com/samhj/AchmadGo/api/models"
	"github.com/samhj/AchmadGo/api/services"
)

//AdminOperationsController for settings related operations
func AdminOperationsController(s *models.Server) {

	resp := &models.Response{
		StatusCd: 400,
		Succ:     false,
		Data:     nil,
		Msg:      "",
	}

	collection := db.GetCollection(s, config.Config("USERS"))

	s.App.Patch(config.GetAPIBase()+"admin/sendnewsletter",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			return services.SendNewsLetterEmail(s)

		})
	s.App.Get(config.GetAPIBase()+"users/getAll",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			return services.GetAllUsers(s)

		})

	s.App.Delete(config.GetAPIBase()+"users/deleteUser",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			return services.DeleteUser(s)

		})

		s.App.Patch(config.GetAPIBase()+"users/updateUser",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			return services.UpdateSingleUser(s)

		})

}
