package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Blog ...
type Blog struct {
	ID          string             `bson:"_id,omitempty" json:"blog_id" xml:"blog_id" form:"blog_id"`
	Title       string             `bson:"title,omitempty" json:"title" xml:"title" form:"title"`
	Directory   string             `bson:"directory,omitempty" json:"directory" xml:"directory" form:"directory"`
	Content     string             `bson:"content,omitempty" json:"content" xml:"content" form:"content"`
	AuthorID    primitive.ObjectID `bson:"author_id,omitempty" json:"author_id" xml:"author_id" form:"author_id"`
	Image       string             `bson:"image,omitempty" json:"image" xml:"image" form:"image"`
	Status      string             `bson:"status,omitempty" json:"status" xml:"status" form:"status"`
	Likes       string             `bson:"likes,omitempty" json:"likes" xml:"likes" form:"likes"`
	Views       string             `bson:"views,omitempty" json:"views" xml:"views" form:"views"`
	Author      User               `bson:"author,omitempty" json:"author" xml:"author" form:"author"`
	PublishedAt string             `bson:"published_at,omitempty" json:"published_at" bson:"published_at,omitempty"`
	LastUpdated string             `bson:"last_updated,omitempty" json:"last_updated" bson:"last_updated,omitempty"`
}
