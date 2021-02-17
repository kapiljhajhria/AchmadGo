package services

import (
	"context"
	"encoding/json"

	"github.com/samhj/AchmadGo/api/config"
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

	for k, _ := range blogs {
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

//UpdateBlog ...
func UpdateBlog(s *models.Server) error {
	//create a new instance of the Blog struct as blog
	blog := models.Blog{}
	s.Ctx.BodyParser(&blog)

	filter := bson.M{"_id": blog.ID}

	blog.ID = ""

	update := bson.M{
		"$set": &blog,
	}

	_, err := s.Coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = config.ErrorMSG + " " + err.Error()
		return resp.JSON(s.Resp)
	}

	//return response
	s.Resp.Ctx = s.Ctx
	s.Resp.StatusCd = 200
	s.Resp.Msg = config.UpdateSuccess
	s.Resp.Data = nil
	s.Resp.Succ = true

	return resp.JSON(s.Resp)
}

//DeleteBlog ...
func DeleteBlog(s *models.Server) error {
	//create a new instance of the Blog struct as blog
	blog := models.Blog{}
	//populate blog with the data in the body of the request.
	s.Ctx.BodyParser(&blog)

	filter := bson.M{"_id": blog.ID}

	_, err := s.Coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		//blog does not exist
		s.Resp.Data = nil
		s.Resp.StatusCd = 400
		s.Resp.Succ = false
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = "Error: " + err.Error()
		return resp.JSON(s.Resp)
	}

	//return response
	s.Resp.Ctx = s.Ctx
	s.Resp.StatusCd = 200
	s.Resp.Msg = "Blog Deleted Successfully!"
	s.Resp.Data = nil
	s.Resp.Succ = true

	return resp.JSON(s.Resp)
}

//AddBlog ...
func AddBlog(s *models.Server) error {
	//create a new instance of the Blog struct as blog
	blog := models.Blog{}
	//populate blog with the data in the body of the request.
	s.Ctx.BodyParser(&blog)

	_, err := s.Coll.InsertOne(context.TODO(), blog)
	if err != nil {
		//failed to create blog
		s.Resp.Data = nil
		s.Resp.Succ = false
		s.Resp.StatusCd = 400
		s.Resp.Ctx = s.Ctx
		s.Resp.Msg = "Failed To Create Blog At This Moment!"
		return resp.JSON(s.Resp)
	}

	//set blog ID
	s.Resp.Ctx = s.Ctx
	s.Resp.StatusCd = 200
	s.Resp.Msg = "Blog Created Successfully!"
	s.Resp.Data = nil
	s.Resp.Succ = true

	//return response
	return resp.JSON(s.Resp)
}
