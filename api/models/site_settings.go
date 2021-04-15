package models

//MagAuthor ...
type MagAuthor struct {
	UserID   string `bson:"user_id,omitempty" json:"user_id" xml:"user_id" form:"user_id"`
	UserData User   `bson:"user_data,omitempty" json:"user_data" xml:"user_data" form:"user_data"`
	Status   string `bson:"status,omitempty" json:"status" xml:"status" form:"status"`
	Role     string `bson:"role,omitempty" json:"role" xml:"role" form:"role"`
}

//Directory ...
type Directory struct {
	AuthorID  string `bson:"author_id,omitempty" json:"author_id" xml:"author_id" form:"author_id"`
	Directory string `bson:"directory,omitempty" json:"directory" xml:"directory" form:"directory"`
}

//Magazine ... visibility
type Magazine struct {
	Title       string      `bson:"title,omitempty" json:"title" xml:"title" form:"title"`
	PubType     string      `bson:"pub_type,omitempty" json:"pub_type" xml:"pub_type" form:"pub_type"`
	Image       string      `bson:"image,omitempty" json:"image" xml:"image" form:"image"`
	MagID       string      `bson:"magId,omitempty" json:"magId" xml:"magId" form:"magId"`
	Status      string      `bson:"status,omitempty" json:"status" xml:"status" form:"status"`
	Visibility  string      `bson:"visibility,omitempty" json:"visibility" xml:"visibility" form:"visibility"`
	Directories []string    `bson:"directories,omitempty" json:"directories" xml:"directories" form:"directories"`
	Authors     []MagAuthor `bson:"authors,omitempty" json:"authors" xml:"authors" form:"authors"`
}

//SiteSettings ...
type SiteSettings struct {
	MaxUploadSize string      `bson:"max_upload_size,omitempty" json:"max_upload_size" xml:"max_upload_size" form:"max_upload_size"`
	AppStoreAppID string      `bson:"apple_store_app_id,omitempty" json:"apple_store_app_id" xml:"apple_store_app_id" form:"apple_store_app_id"`
	Directories   []Directory `bson:"directories,omitempty" json:"directories" xml:"directories" form:"directories"`
	Magazines     []Magazine  `bson:"magazines" json:"magazines" xml:"magazines" form:"magazines"`
}
