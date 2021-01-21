package db

import (
	"context"

	"github.com/samhj/AchmadGo/api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	models "github.com/samhj/AchmadGo/api/models"

	"fmt"
)

//Client connection to mongo
func Client() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Config("MONGO_CLUSTER_CONNECTION")))

	if err != nil {

		fmt.Println("Token Generation Failed:", err.Error())
		return nil

	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {

		fmt.Println("DB Connection Failed:", err.Error())
		return nil
	}

	fmt.Println("Connected to MongoDB!")

	//if connection is okay, return the client
	return client

}

//GetCollection ...
func GetCollection(s *models.Server,collection string) *mongo.Collection {
	return s.DB.Database(config.Config("DB")).Collection(collection)
}

//DeferDB ...
func DeferDB(s *models.Server) {
	s.DB.Disconnect(context.TODO())
}
