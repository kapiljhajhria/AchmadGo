package validations

import (
	"errors"
	"html"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/samhj/AchmadGo/api/config"
	models "github.com/samhj/AchmadGo/api/models"
)

//User alias of the User model in models.User
type User models.User

//BeforeCreate initiates the hashing of the user's password
func (u *User) BeforeCreate() error {
	hashedPassword, err := config.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

//Prepare sanitizes the values gotten from the client
func (u *User) Prepare() {
	u.FullName = html.EscapeString(strings.TrimSpace(u.FullName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
}

//IsValid checks if the supplied data from the client are valid/acceptable
func (u *User) IsValid(action string) error {
	switch strings.ToLower(action) {
	case "token":
		if u.ID == "" {
			return errors.New("User ID is required")
		}
		if u.Token == "" {
			return errors.New("Token is required")
		}

	case "id":
		if u.ID == "" {
			return errors.New("User ID is required")
		}
	case "changepassword":
		if u.Password == "" {
			return errors.New("Password is required")
		}
		if len(u.Password) < 6 {
			return errors.New("Password length must equal or greater than 6")
		}
		if u.Token == "" {
			return errors.New("Token is required")
		}
		if u.ID == "" {
			return errors.New("User ID is required")
		}
		
	case "email":
		if u.Email == "" {
			return errors.New("User Email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email is invalid")
		}

	case "update":
		if u.FullName == "" {
			return errors.New("FullName is required")
		}
		if u.Email == "" {
			return errors.New("User Email is required")
		}

	case "login":
		if u.Password == "" {
			return errors.New("Password is required")
		}
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if u.Email != ""{
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				return errors.New("Email is invalid")
			}
		}
		if len(u.Password) < 6 {
			return errors.New("Password length must equal or greater than 6")
		}

	default:
		if u.FullName == "" {
			return errors.New("FullName is required")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
		if len(u.Password) < 6 {
			return errors.New("Password length must equal or greater than 6")
		}
		if u.Status == "" {
			return errors.New("User status is required")
		}
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email is invalid")
		}
	}

	return nil
}
