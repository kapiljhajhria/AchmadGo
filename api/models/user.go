package models

//User ...
type User struct {
	ID         string `bson:"_id,omitempty" json:"user_id" xml:"user_id" form:"user_id"`
	FullName   string `bson:"fullname,omitempty" json:"fullname" xml:"fullname" form:"fullname"`
	ProfileImage   string `bson:"profile_image,omitempty" json:"profile_image" xml:"profile_image" form:"profile_image"`
	Email      string `bson:"email,omitempty" json:"email" xml:"email" form:"email"`
	Gender      string `bson:"gender,omitempty" json:"gender" xml:"gender" form:"gender"`
	DOB      string `bson:"dob,omitempty" json:"dob" xml:"dob" form:"dob"`
	Country      string `bson:"country,omitempty" json:"country" xml:"country" form:"country"`
	Password   string `bson:"password,omitempty" json:"password" xml:"password" form:"password"`
	Status   string `bson:"status,omitempty" json:"status" xml:"status" form:"status"`
	IsVerified bool   `bson:"isverified,omitempty" json:"isverified" xml:"isverified" form:"isverified"`
	Token      string `bson:"token,omitempty" json:"token" xml:"token" form:"token"`
}
