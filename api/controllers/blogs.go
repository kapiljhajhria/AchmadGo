package controllers //declare package name here

import (
	"github.com/samhj/AchmadGo/api/config"

	"github.com/gofiber/fiber/v2"
	"github.com/samhj/AchmadGo/api/config/db"
	"github.com/samhj/AchmadGo/api/models"
	"github.com/samhj/AchmadGo/api/services"
)

//BlogsController for settings related operations
func BlogsController(s *models.Server) {

	resp := &models.Response{
		StatusCd: 400,
		Succ:     false,
		Data:     nil,
		Msg:      "",
	}

	collection := db.GetCollection(s, config.Config("BLOGS"))

	s.App.Get(config.GetAPIBase()+"blogs/getAll",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			return services.GetBlogs(s)

		})
	s.App.Patch(config.GetAPIBase()+"blogs/updateSingle",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			return services.UpdateBlog(s)

		})
	s.App.Delete(config.GetAPIBase()+"blogs/deleteSingle",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			return services.DeleteBlog(s)

		})

	s.App.Patch(config.GetAPIBase()+"blogs/addSingle",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			return services.AddBlog(s)

		})

}
