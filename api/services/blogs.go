package services

import (
	"context"
	"encoding/json"

	"github.com/samhj/AchmadGo/api/models"
	resp "github.com/samhj/AchmadGo/api/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//GetBlogs ...
func GetBlogs(s *models.Server) error {
	var blogs []models.Blog

	lookupStage := bson.D{{"$lookup", bson.D{{"from", "users"}, {"localField", "author_id"}, {"foreignField", "_id"}, {"as", "author"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$author"}, {"preserveNullAndEmptyArrays", false}}}}

	blogsLoadedCursor, err := s.Coll.Aggregate(context.TODO(), mongo.Pipeline{lookupStage, unwindStage})
	if err != nil {
		s.Resp.Data = nil
		s.Resp.StatusCd = 400
		s.Resp.Succ = false
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = "Error: " + err.Error()
		return resp.JSON(s.Resp)
	}

	if err = blogsLoadedCursor.All(context.TODO(), &blogs); err != nil {
		s.Resp.Data = nil
		s.Resp.StatusCd = 400
		s.Resp.Succ = false
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = "Error: " + err.Error()
		return resp.JSON(s.Resp)
	}

	for k,_ := range blogs{
		blogs[k].Author.Email = ""
		blogs[k].Author.DOB = ""
		blogs[k].Author.Token = ""
		blogs[k].Author.Password = ""
	}

	res, _ := json.Marshal(blogs)

	s.Resp.Ctx = s.Ctx
	s.Resp.StatusCd = 200
	s.Resp.Msg = "Blogs Fetched Successfully!"
	s.Resp.Data = res
	s.Resp.Succ = true
	return resp.JSON(s.Resp)
}
