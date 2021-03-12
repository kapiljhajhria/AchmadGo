package models

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

//Server ...
type Server struct {
	DB  *mongo.Client
	App *fiber.App
	Ctx *fiber.Ctx
	Resp *Response
	Coll *mongo.Collection
	UsersColl *mongo.Collection
	TokenColl *mongo.Collection
}