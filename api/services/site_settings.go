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

//UpdateData ...
type UpdateData struct {
	MgID     string          `json:"mgID" xml:"mgID" form:"mgID"`
	Type     string          `json:"type" xml:"type" form:"type"`
	From     string          `json:"from" xml:"from" form:"from"`
	Magazine models.Magazine `json:"mg" xml:"mg" form:"mg"`
}

//GetSiteSettings ...
func GetSiteSettings(s *models.Server) error {

	settingsObj, err := GetSettings("direct",s)
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

	r := len(settings.Magazines)

	update := bson.M{
		"$set": &settings,
	}

	sOBJ, _ := GetSettings("update",s)

	if r == 0 {
		//i.e no magazine was attached
		updateData := new(UpdateData)
		//populate updateData with the data in the body of the request.
		err := s.Ctx.BodyParser(updateData)

		if err == nil {

			if updateData.From == "magazine" {

				if updateData.Type == "delete" {

					settings.Magazines = removeMagazineObjByPropVal(sOBJ.Magazines, updateData.MgID)

				} else if updateData.Type == "update" {

					newMagsList := removeMagazineObjByPropVal(sOBJ.Magazines, updateData.MgID)
					newMagsList = append(newMagsList, updateData.Magazine)
					settings.Magazines = newMagsList

				} else if updateData.Type == "add" {
					update = bson.M{"$push": bson.M{"magazines": updateData.Magazine}}

				}

			} else {
				settings.Magazines = sOBJ.Magazines
			}
		}

	} else {

		settings.Magazines = sOBJ.Magazines
	}

	_, err := s.Coll.UpdateOne(context.TODO(), bson.M{}, update)
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

func removeMagazineObjByPropVal(mags []models.Magazine, magID string) []models.Magazine {

	newMagsList := []models.Magazine{}

	for _, mag := range mags {
		if mag.MagID != magID {
			newMagsList = append(newMagsList, mag)
		}
	}

	return newMagsList
}

//GetSettings ... contributor
func GetSettings(from string, s *models.Server) (models.SiteSettings, error) {

	var settingsObj models.SiteSettings
	sChan := make(chan models.SiteSettings)

	defer close(sChan)

	err := s.Coll.FindOne(context.TODO(), bson.M{}).Decode(&settingsObj)

	getter := func() {
		mags := settingsObj.Magazines

		newS := s
		newS.Coll = s.UsersColl

		if len(mags) > 0 {
			for i, mag := range mags {

				if len(mag.Authors) > 0 {

					for i, author := range mag.Authors {
						objectID, _ := primitive.ObjectIDFromHex(author.UserID)

						filter := bson.M{"_id": objectID}

						user, err := GetUser(filter, newS)

						if err == nil {
							//copy data from validations.User into models.User
							//using models.User(user)
							author.UserData = models.User(user)
							author.UserData.Password = ""
						}
						mag.Authors[i] = author
					}
					mags[i] = mag
				}
			}

			settingsObj.Magazines = mags
		}
	}

	if from == "direct" {

		getter()

	} else {

		go getter()
	}

	return settingsObj, err
}
