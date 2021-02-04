package models

import "time"

//UserSettings ...
type UserSettings struct {
	AllowNotifications      bool `bson:"allow_notifications" json:"allow_notifications" xml:"allow_notifications" form:"allow_notifications"`
	ShowLatestBlogUpdates   bool `bson:"show_latest_blog_updates" json:"show_latest_blog_updates" xml:"show_latest_blog_updates" form:"show_latest_blog_updates"`
	DisplayMyPostsToMeFirst bool `bson:"display_my_posts_to_me_first" json:"display_my_posts_to_me_first" xml:"display_my_posts_to_me_first" form:"display_my_posts_to_me_first"`
	KeepAPublicProfile      bool `bson:"keep_a_public_profile" json:"keep_a_public_profile" xml:"keep_a_public_profile" form:"keep_a_public_profile"`
	KeepMyBlogPostsPrivate  bool `bson:"keep_my_blog_posts_private" json:"keep_my_blog_posts_private" xml:"keep_my_blog_posts_private" form:"keep_my_blog_posts_private"`
}

//User ...
type User struct {
	ID           string       `bson:"_id,omitempty" json:"user_id" xml:"user_id" form:"user_id"`
	FullName     string       `bson:"fullname,omitempty" json:"fullname" xml:"fullname" form:"fullname"`
	ProfileImage string       `bson:"profile_image,omitempty" json:"profile_image" xml:"profile_image" form:"profile_image"`
	Email        string       `bson:"email,omitempty" json:"email" xml:"email" form:"email"`
	Gender       string       `bson:"gender,omitempty" json:"gender" xml:"gender" form:"gender"`
	DOB          string       `bson:"dob,omitempty" json:"dob" xml:"dob" form:"dob"`
	Country      string       `bson:"country,omitempty" json:"country" xml:"country" form:"country"`
	Password     string       `bson:"password,omitempty" json:"password" xml:"password" form:"password"`
	Settings     UserSettings `bson:"settings,omitempty" json:"settings" xml:"settings" form:"settings"`
	Status       string       `bson:"status,omitempty" json:"status" xml:"status" form:"status"`
	IsVerified   bool         `bson:"isverified,omitempty" json:"isverified" xml:"isverified" form:"isverified"`
	Token        string       `bson:"token,omitempty" json:"token" xml:"token" form:"token"`
	CreatedAt    time.Time    `json:"created_at" bson:"created_at,omitempty"`
}
