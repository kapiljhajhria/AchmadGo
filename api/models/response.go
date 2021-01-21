package models

import(
	"github.com/gofiber/fiber/v2"
)

//Response ...
type Response struct {
	Ctx *fiber.Ctx
	StatusCd int
	Succ bool
	Msg string
	Data []byte
}