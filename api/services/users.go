package services

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/samhj/AchmadGo/api/auth"
	"github.com/samhj/AchmadGo/api/config"
	"github.com/samhj/AchmadGo/api/models"
	resp "github.com/samhj/AchmadGo/api/responses"
	"github.com/samhj/AchmadGo/api/utils"
	"github.com/samhj/AchmadGo/api/validations"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

//Login ...
func Login(s *models.Server) error {

	//create a new instance of the User struct as user
	user := validations.User{}
	//populate user with the data in the body of the request.
	s.Ctx.BodyParser(&user)

	//sanitize the fields gotten from the client
	user.Prepare()

	userPassword := user.Password

	filter := bson.M{"email": user.Email}

	//check if user exist
	user, err := GetUser(filter, s)
	if err != nil {
		//user does not exist
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.MatchFailed
		return resp.JSON(s.Resp)
	}

	//validate password and return error if it is not correct
	err = config.VerifyPassword(user.Password, userPassword)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		//invalid password
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.MatchFailed
		return resp.JSON(s.Resp)
	}

	token, _ := auth.CreateToken(user.ID)
	user.Token = token
	user.Password = ""

	res, _ := json.Marshal(user)
	s.Resp.Ctx = s.Ctx
	s.Resp.StatusCd = 200
	s.Resp.Msg = config.LoginSuccess
	s.Resp.Data = res
	s.Resp.Succ = true
	return resp.JSON(s.Resp)

}

//Register ...
func Register(s *models.Server) error {

	//create a new instance of the User struct as user
	user := validations.User{}
	//populate user with the data in the body of the request.
	s.Ctx.BodyParser(&user)

	//sanitize the fields gotten from the client
	user.Prepare()

	userPassword := user.Password

	filter := bson.M{"email": user.Email}

	_, err := GetUser(filter, s)
	if err == nil {
		//user already exists
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.EmailNotAvailable
		return resp.JSON(s.Resp)
	}

	//hash user's password
	hashPass, _ := config.Hash(userPassword)
	user.Password = string(hashPass)
	user.Settings = models.UserSettings{nil,true, true, false, true, false}
	user.CreatedAt = time.Now()

	result, err := s.Coll.InsertOne(context.TODO(), user)
	if err != nil {
		//failed to create user
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.FailedToCreateUser
		return resp.JSON(s.Resp)
	}

	//set user ID
	user.ID = string(result.InsertedID.(primitive.ObjectID).Hex())

	token, _ := auth.CreateToken(user.ID)
	user.Token = token
	user.Password = ""

	res, _ := json.Marshal(user)
	s.Resp.Ctx = s.Ctx
	s.Resp.StatusCd = 200
	s.Resp.Msg = config.SignUpSuccess
	s.Resp.Data = res
	s.Resp.Succ = true

	//send welcome email
	go func() {
		emailModel := &models.Email{
			ReceiverEmail: user.Email,
			ReceiverName:  user.FullName,
			Subject:       config.SignUpSubject,
			FileName:      "welcome",
			SenderEmail:   config.GeneralSenderEmail,
			Replacer: map[string]string{
				"%username%":         strings.Split(user.FullName, " ")[0],
				"%verification_url%": config.VerificationURL + user.ID,
			},
		}

		utils.SendEmail(emailModel)
	}()

	//return response
	return resp.JSON(s.Resp)

}

//ResendWelcomEmail ...
func ResendWelcomEmail(s *models.Server) error {

	//create a new instance of the User struct as user
	user := validations.User{}
	s.Ctx.BodyParser(&user)
	user.Prepare()

	filter := bson.M{"email": user.Email}

	user, err := GetUser(filter, s)
	if err != nil {
		//user does not exist
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.InvalidUser
		return resp.JSON(s.Resp)
	}

	s.Resp.Ctx = s.Ctx
	s.Resp.StatusCd = 200
	s.Resp.Msg = config.EmailSent
	s.Resp.Data = nil
	s.Resp.Succ = true

	//re-send welcome email
	err = utils.SendEmail(&models.Email{
		ReceiverEmail: user.Email,
		ReceiverName:  user.FullName,
		Subject:       config.SignUpSubject,
		FileName:      "welcome",
		SenderEmail:   config.GeneralSenderEmail,
		Replacer: map[string]string{
			"%username%":         strings.Split(user.FullName, " ")[0],
			"%verification_url%": config.VerificationURL + user.ID,
		},
	})

	if err != nil {
		s.Resp.StatusCd = 400
		s.Resp.Msg = config.CantSendEmail
		s.Resp.Succ = false
	}

	//return response
	return resp.JSON(s.Resp)

}

//SendRecoveryEmail ...
func SendRecoveryEmail(s *models.Server) error {
	//create a new instance of the User struct as user
	user := validations.User{}
	s.Ctx.BodyParser(&user)
	user.Prepare()

	filter := bson.M{"email": user.Email}

	user, err := GetUser(filter, s)
	if err != nil {
		//user does not exist
		s.Resp.Data = nil
		s.Resp.Succ = true
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.EmailSentRecovery
		return resp.JSON(s.Resp)
	}

	s.Resp.Ctx = s.Ctx
	s.Resp.StatusCd = 200
	s.Resp.Msg = config.EmailSentRecovery
	s.Resp.Data = nil
	s.Resp.Succ = true

	token := AddToken(user.ID, s.TokenColl, "p-recovery")

	//send password recovery email
	err = utils.SendEmail(&models.Email{
		ReceiverEmail: user.Email,
		ReceiverName:  user.FullName,
		Subject:       config.PasswordRecoverySubject,
		FileName:      "account_recovery",
		SenderEmail:   config.GeneralSenderEmail,
		Replacer: map[string]string{
			"%username%":     strings.Split(user.FullName, " ")[0],
			"%recovery_url%": config.RecoveryURL + user.ID + "&t=" + token,
		},
	})

	if err != nil {
		s.Resp.StatusCd = 400
		s.Resp.Msg = config.CantSendEmail
		s.Resp.Succ = false
	}

	//return response
	return resp.JSON(s.Resp)
}

//CheckPasswordRecoveryToken ...
func CheckPasswordRecoveryToken(s *models.Server) error {

	//create a new instance of the User struct as user
	user := validations.User{}
	s.Ctx.BodyParser(&user)
	user.Prepare()

	//check if token exist/valid
	tokenOBJ, err := GetToken(user.ID, "p-recovery", s.TokenColl)
	if err != nil || tokenOBJ.Token != user.Token {
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.InvalidToken
		return resp.JSON(s.Resp)
	}

	//return response
	s.Resp.Ctx = s.Ctx
	s.Resp.StatusCd = 200
	s.Resp.Msg = config.TokenValid
	s.Resp.Data = nil
	s.Resp.Succ = true

	return resp.JSON(s.Resp)
}

//UpdatePassword ...
func UpdatePassword(s *models.Server) error {

	//create a new instance of the User struct as user
	user := validations.User{}
	s.Ctx.BodyParser(&user)
	user.Prepare()

	//check if token exist/valid
	tokenOBJ, err := GetToken(user.ID, "p-recovery", s.TokenColl)
	if err != nil || tokenOBJ.Token != user.Token {
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.InvalidToken
		return resp.JSON(s.Resp)
	}

	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.InvalidID
		return resp.JSON(s.Resp)
	}

	filter := bson.M{"_id": objectID}
	hashPass, _ := config.Hash(user.Password)

	update := bson.M{
		"$set": validations.User{Password: string(hashPass)},
	}

	_, err = s.Coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.ErrorMSG
		return resp.JSON(s.Resp)
	}

	//delete token
	DeleteToken(user.Token, s.TokenColl)

	//return response
	s.Resp.Ctx = s.Ctx
	s.Resp.StatusCd = 200
	s.Resp.Msg = config.UpdateSuccess
	s.Resp.Data = nil
	s.Resp.Succ = true

	return resp.JSON(s.Resp)
}

//VerifyAccount ...
func VerifyAccount(s *models.Server) error {

	//create a new instance of the User struct as user
	user := validations.User{}
	s.Ctx.BodyParser(&user)
	user.Prepare()

	objectID, err := primitive.ObjectIDFromHex(user.ID)
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
		"$set": validations.User{IsVerified: true},
	}

	_, err = s.Coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.ErrorMSG
		return resp.JSON(s.Resp)
	}

	//return response
	s.Resp.Ctx = s.Ctx
	s.Resp.StatusCd = 200
	s.Resp.Msg = config.AccountVerified
	s.Resp.Data = nil
	s.Resp.Succ = true

	return resp.JSON(s.Resp)
}

//IsUserVerified ...
func IsUserVerified(s *models.Server) error {
	//create a new instance of the User struct as user
	user := validations.User{}
	s.Ctx.BodyParser(&user)
	user.Prepare()

	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.InvalidID
		return resp.JSON(s.Resp)
	}

	filter := bson.M{"_id": objectID}

	user, err = GetUser(filter, s)
	if err != nil {
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.ErrorMSG
		return resp.JSON(s.Resp)
	}

	if !user.IsVerified {
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.AccountPendingVerification
		return resp.JSON(s.Resp)
	}

	//return response
	s.Resp.Ctx = s.Ctx
	s.Resp.StatusCd = 200
	s.Resp.Msg = config.AccountVerificationSuccess
	s.Resp.Data = nil
	s.Resp.Succ = true

	return resp.JSON(s.Resp)
}

//UpdateProfile ...
func UpdateProfile(s *models.Server) error {
	//create a new instance of the User struct as user
	user := validations.User{}
	s.Ctx.BodyParser(&user)
	user.Prepare()

	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.InvalidID
		return resp.JSON(s.Resp)
	}

	filter := bson.M{"_id": objectID}

	//remove the user ID from the user struct
	user.ID = ""

	update := bson.M{
		"$set": &user,
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

//GetUser ...
func GetUser(filter bson.M, s *models.Server) (validations.User, error) {

	var userObj validations.User

	err := s.Coll.FindOne(context.TODO(), filter).Decode(&userObj)

	return userObj, err
}
