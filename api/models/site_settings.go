package models

//MagAuthor ...
type MagAuthor struct {
	ID       string `bson:"_id,omitempty" json:"author_id" xml:"author_id" form:"author_id"`
	UserID   string `bson:"user_id,omitempty" json:"user_id" xml:"user_id" form:"user_id"`
	UserData User   `bson:"user_data,omitempty" json:"user_data" xml:"user_data" form:"user_data"`
	Status   string `bson:"status,omitempty" json:"status" xml:"status" form:"status"`
	Role     string `bson:"role,omitempty" json:"role" xml:"role" form:"role"`
}

//Directory ...
type Directory struct {
	ID        string `bson:"_id,omitempty" json:"dir_id" xml:"dir_id" form:"dir_id"`
	AuthorID  string `bson:"author_id,omitempty" json:"author_id" xml:"author_id" form:"author_id"`
	Directory string `bson:"directory,omitempty" json:"directory" xml:"directory" form:"directory"`
}

//Magazine ...
type Magazine struct {
	ID          string      `bson:"_id,omitempty" json:"mag_id" xml:"mag_id" form:"mag_id"`
	Title       string      `bson:"title,omitempty" json:"title" xml:"title" form:"title"`
	Image       string      `bson:"image,omitempty" json:"image" xml:"image" form:"image"`
	MagID       string      `bson:"magId,omitempty" json:"magId" xml:"magId" form:"magId"`
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
