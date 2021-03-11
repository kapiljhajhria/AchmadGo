package models

//MagAuthor
type MagAuthor struct{
	UserID string `bson:"user_id,omitempty" json:"user_id" xml:"user_id" form:"user_id"`
	Status string `bson:"status,omitempty" json:"status" xml:"status" form:"status"`
	Role string `bson:"role,omitempty" json:"role" xml:"role" form:"role"`
}

//Magazine ...
type Magazine struct {
	Title      string `bson:"title,omitempty" json:"title" xml:"title" form:"title"`
	Image   string `bson:"image,omitempty" json:"image" xml:"image" form:"image"`
	MagID string`bson:"magId,omitempty" json:"magId" xml:"magId" form:"magId"`
	Directories   []string `bson:"directories,omitempty" json:"directories" xml:"directories" form:"directories"`
	Author 	 []MagAuthor `bson:"authors" json:"authors" xml:"authors" form:"authors"`
}

//SiteSettings ...
type SiteSettings struct {
	MaxUploadSize string   `bson:"max_upload_size,omitempty" json:"max_upload_size" xml:"max_upload_size" form:"max_upload_size"`
	AppStoreAppID string   `bson:"apple_store_app_id,omitempty" json:"apple_store_app_id" xml:"apple_store_app_id" form:"apple_store_app_id"`
	Directories   []string `bson:"directories,omitempty" json:"directories" xml:"directories" form:"directories"`
	Magazines   []Magazine `bson:"magazines" json:"magazines" xml:"magazines" form:"magazines"`
}
