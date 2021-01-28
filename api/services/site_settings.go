package services

import (
	"context"
	"encoding/json"

	"github.com/samhj/AchmadGo/api/config"
	"github.com/samhj/AchmadGo/api/models"
	resp "github.com/samhj/AchmadGo/api/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const settingsID = "600dc081051e9b3081b37ca6"

//GetSiteSettings ...
func GetSiteSettings(s *models.Server) error {

	objectID, err := primitive.ObjectIDFromHex(settingsID)
	if err != nil {
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.InvalidID
		return resp.JSON(s.Resp)
	}

	filter := bson.M{"_id": objectID}

	var settingsObj models.SiteSettings

	err = s.Coll.FindOne(context.TODO(), filter).Decode(&settingsObj)
	if err != nil {
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.NoRecord
		return resp.JSON(s.Resp)
	}

	res, _ := json.Marshal(settingsObj)
	s.Resp.Ctx = s.Ctx
	s.Resp.StatusCd = 200
	s.Resp.Msg = config.SettingsFetched
	s.Resp.Data = res
	s.Resp.Succ = true
	return resp.JSON(s.Resp)
}

//UpdateSiteSettings ...
func UpdateSiteSettings(s *models.Server) error {

	//create a new instance of the SiteSettings struct as settings
	settings := models.SiteSettings{}
	s.Ctx.BodyParser(&settings)

	objectID, err := primitive.ObjectIDFromHex(settingsID)
	if err != nil {
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.InvalidID
		return resp.JSON(s.Resp)
	}

	filter := bson.M{"_id": objectID}

	update := bson.M{
		"$set": &settings,
	}

	_, err = s.Coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.ErrorMSG + " " + err.Error()
		return resp.JSON(s.Resp)
	}

	//return response
	s.Resp.Ctx = s.Ctx
	s.Resp.StatusCd = 200
	s.Resp.Msg = config.UpdateSuccess
	s.Resp.Data = nil
	s.Resp.Succ = true

	return resp.JSON(s.Resp)
}
