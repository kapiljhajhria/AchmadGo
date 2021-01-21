package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samhj/AchmadGo/api/config"
	"github.com/samhj/AchmadGo/api/models"
	resp "github.com/samhj/AchmadGo/api/responses"
	"github.com/samhj/AchmadGo/api/validations"
)

//UserValidators ...
func UserValidators(s *models.Server) {

	s.App.Use(config.GetAPIBase()+"user/login", func(c *fiber.Ctx) error {

		return ValidateUser("login", c)
	})

	s.App.Use(config.GetAPIBase()+"user/register", func(c *fiber.Ctx) error {

		return ValidateUser("", c)
	})

	s.App.Use(config.GetAPIBase()+"user/resendWelcomEmail", func(c *fiber.Ctx) error {

		return ValidateUser("email", c)
	})

	s.App.Use(config.GetAPIBase()+"user/sendRecoveryEmail", func(c *fiber.Ctx) error {

		return ValidateUser("email", c)
	})

	s.App.Use(config.GetAPIBase()+"user/changePassword", func(c *fiber.Ctx) error {

		return ValidateUser("changepassword", c)
	})

	s.App.Use(config.GetAPIBase()+"user/verifyAccount", func(c *fiber.Ctx) error {

		return ValidateUser("id", c)
	})

	s.App.Use(config.GetAPIBase()+"user/isUserVerified", func(c *fiber.Ctx) error {

		return ValidateUser("id", c)
	})

	s.App.Use(config.GetAPIBase()+"user/checkPasswordRecoveryToken", func(c *fiber.Ctx) error {

		return ValidateUser("token", c)
	})

}

//ValidateUser ...
func ValidateUser(action string, c *fiber.Ctx) error {

	user := validations.User{}
	c.BodyParser(&user)

	//sanitize the fields gotten from the client
	user.Prepare()

	err := user.IsValid(action)

	//throw error if there's any invalid field value(s)
	if err != nil {
		res := &models.Response{
			Ctx:      c,
			StatusCd: 400,
			Succ:     false,
			Data:     nil,
			Msg:      err.Error(),
		}
		return resp.JSON(res)
	}

	return c.Next()

}
