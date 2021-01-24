package controllers //declare package name here

import (
	"github.com/samhj/AchmadGo/api/config"

	"github.com/gofiber/fiber/v2"
	"github.com/samhj/AchmadGo/api/config/db"
	"github.com/samhj/AchmadGo/api/models"
	"github.com/samhj/AchmadGo/api/services"
)

//UserController for user related operations
func UserController(s *models.Server) {

	resp := &models.Response{
		StatusCd: 400,
		Succ:     false,
		Data:     nil,
		Msg:      "",
	}

	collection := db.GetCollection(s, config.Config("USERS"))
	tokenColl := db.GetCollection(s, config.Config("TOKENS"))

	//handle sign-in
	s.App.Post(config.GetAPIBase()+"user/login",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			return services.Login(s)
		})

	//handle sign-up
	s.App.Post(config.GetAPIBase()+"user/register",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			return services.Register(s)
		})

	//handle resending of welcome email
	s.App.Patch(config.GetAPIBase()+"user/resendWelcomEmail",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			return services.ResendWelcomEmail(s)
		})

	//send password recovery email
	s.App.Patch(config.GetAPIBase()+"user/sendRecoveryEmail",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			s.TokenColl = tokenColl
			return services.SendRecoveryEmail(s)
		})

	//change user's password
	s.App.Patch(config.GetAPIBase()+"user/changePassword",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			s.TokenColl = tokenColl
			return services.UpdatePassword(s)
		})

	//verify account
	s.App.Patch(config.GetAPIBase()+"user/verifyAccount",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			return services.VerifyAccount(s)
		})

	//verify account
	s.App.Patch(config.GetAPIBase()+"user/isUserVerified",
		func(c *fiber.Ctx) error {
			s.Ctx = c
			s.Resp = resp
			s.Coll = collection
			return services.IsUserVerified(s)
		})

	/**check if token and user id exists in the tokens collection
	***before showing the change of password form inputs
	**/
	s.App.Patch(config.GetAPIBase()+"user/checkPasswordRecoveryToken",
	func(c *fiber.Ctx) error {
		s.Ctx = c
		s.Resp = resp
		s.Coll = collection
		s.TokenColl = tokenColl
		return services.CheckPasswordRecoveryToken(s)
	})

	s.App.Patch(config.GetAPIBase()+"user/updateProfile",
	func(c *fiber.Ctx) error {
		s.Ctx = c
		s.Resp = resp
		s.Coll = collection
		return services.UpdateProfile(s)

	})

}
