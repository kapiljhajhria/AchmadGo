package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Blog ...
type Blog struct {
	ID       string             `bson:"_id,omitempty" json:"_id" xml:"_id" form:"_id"`
	Title       string             `bson:"title,omitempty" json:"title" xml:"title" form:"title"`
	Directory   string             `bson:"directory,omitempty" json:"directory" xml:"directory" form:"directory"`
	Content     string             `bson:"content,omitempty" json:"content" xml:"content" form:"content"`
	AuthorID    primitive.ObjectID `bson:"author_id,omitempty" json:"author_id" xml:"author_id" form:"author_id"`
	Image       string             `bson:"image,omitempty" json:"image" xml:"image" form:"image"`
	Status      string             `bson:"status,omitempty" json:"status" xml:"status" form:"status"`
	Author      User               `bson:"author,omitempty" json:"author" xml:"author" form:"author"`
	PublishedAt   string             `json:"published_at" bson:"published_at,omitempty"`
	LastUpdated string             `json:"last_updated" bson:"last_updated,omitempty"`
}
