package models

//SiteSettings ...
type SiteSettings struct {
	MaxUploadSize string `bson:"max_upload_size,omitempty" json:"max_upload_size" xml:"max_upload_size" form:"max_upload_size"`
	AppStoreAppID string `bson:"apple_store_app_id,omitempty" json:"apple_store_app_id" xml:"apple_store_app_id" form:"apple_store_app_id"`
}