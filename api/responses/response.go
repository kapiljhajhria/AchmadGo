package resp

/** This package contains app wide re-usable api response **/

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samhj/AchmadGo/api/models"
)

//JSON ...
func JSON(resp *models.Response) error {

	return resp.Ctx.Status(resp.StatusCd).JSON(&fiber.Map{
		"success": resp.Succ,
		"message": resp.Msg,
		"data":    string(resp.Data),
	})
}
