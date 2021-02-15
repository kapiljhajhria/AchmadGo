package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Blog ...
type Blog struct {
	Title     string             `bson:"title,omitempty" json:"title" xml:"title" form:"title"`
	Directory string             `bson:"directory,omitempty" json:"directory" xml:"directory" form:"directory"`
	Content   string             `bson:"content,omitempty" json:"content" xml:"content" form:"content"`
	AuthorID  primitive.ObjectID `bson:"author_id,omitempty" json:"author_id" xml:"author_id" form:"author_id"`
	Image     string             `bson:"image,omitempty" json:"image" xml:"image" form:"image"`
	Author    User               `bson:"author,omitempty" json:"author" xml:"author" form:"author"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}
