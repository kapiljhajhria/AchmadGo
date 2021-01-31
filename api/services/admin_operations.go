package services

import (
	"context"
	"strings"

	"github.com/samhj/AchmadGo/api/config"
	"github.com/samhj/AchmadGo/api/models"
	resp "github.com/samhj/AchmadGo/api/responses"
	"github.com/samhj/AchmadGo/api/utils"
	"github.com/samhj/AchmadGo/api/validations"
	"go.mongodb.org/mongo-driver/bson"
)

//NewsLetter ...
type NewsLetter struct {
	Image     string `bson:"header_image" json:"header_image" xml:"header_image" form:"header_image"`
	Subject   string `bson:"subject" json:"subject" xml:"subject" form:"subject"`
	Recipient string `bson:"recipient" json:"recipient" xml:"recipient" form:"recipient"`
	Message   string `bson:"message" json:"message" xml:"message" form:"message"`
}

//SendNewsLetterEmail ...
func SendNewsLetterEmail(s *models.Server) error {

	//create a new instance of the User struct as user
	newsletter := NewsLetter{}
	s.Ctx.BodyParser(&newsletter)

	usersArr, err := GetUsers(newsletter.Recipient, s)
	if err != nil {
		//user does not exist
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = "Error: " + err.Error()
		return resp.JSON(s.Resp)
	}

	go func() {
		for _, user := range usersArr {
			utils.SendEmail(&models.Email{
				ReceiverEmail: user.Email,
				ReceiverName:  user.FullName,
				Subject:       newsletter.Subject,
				FileName:      "newsletter",
				SenderEmail:   config.GeneralSenderEmail,
				Replacer: map[string]string{
					"%username%": strings.Split(user.FullName, " ")[0],
					"%image%":    newsletter.Image,
					"%subject%":  newsletter.Subject,
					"%message%":  newsletter.Message,
				},
			})
		}
	}()

	s.Resp.Ctx = s.Ctx
	s.Resp.StatusCd = 200
	s.Resp.Msg = "Newsletter Sent Successfully!"
	s.Resp.Data = nil
	s.Resp.Succ = true
	return resp.JSON(s.Resp)
}

//GetUsers ...
func GetUsers(recipient string, s *models.Server) ([]validations.User, error) {
	var users []validations.User
	var filter = bson.M{}

	if recipient == "all" {
		filter = bson.M{}
	} else if recipient == "verified" {
		filter = bson.M{"isVerified": true}
	} else if recipient == "unverified" {
		filter = bson.M{"isVerified": nil}
	} else if recipient == "admin" {
		filter = bson.M{"status": "admin"}
	} else {
		filter = bson.M{"status": "admin"}
		// filter = bson.M{"status": bson.M{"$ne": "admin"}}
	}

	cursor, err := s.Coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}
