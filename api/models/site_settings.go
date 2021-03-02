package models

//Magazine ...
type Magazine struct {
	Title      string `bson:"title" json:"title" xml:"title" form:"title"`
	Image   string `bson:"image,omitempty" json:"image" xml:"image" form:"image"`
	MagID string`bson:"magId" json:"magId" xml:"magId" form:"magId"`
	Directories   []string `bson:"directories,omitempty" json:"directories" xml:"directories" form:"directories"`
}

//SiteSettings ...
type SiteSettings struct {
	MaxUploadSize string   `bson:"max_upload_size,omitempty" json:"max_upload_size" xml:"max_upload_size" form:"max_upload_size"`
	AppStoreAppID string   `bson:"apple_store_app_id,omitempty" json:"apple_store_app_id" xml:"apple_store_app_id" form:"apple_store_app_id"`
	Directories   []string `bson:"directories,omitempty" json:"directories" xml:"directories" form:"directories"`
	Magazines   []Magazine `bson:"magazines,omitempty" json:"magazines" xml:"magazines" form:"magazines"`
}
