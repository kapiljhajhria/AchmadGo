package services

import (
	"context"

	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//TOBJ ...
type TOBJ struct {
	Token  string `bson:"token"`
	UserID string `bson:"user_id"`
	Type   string `bson:"type"`
}

//AddToken ...
func AddToken(userID string, tokenColl *mongo.Collection, ttype string) string {

	tokenOBJ, err := GetToken(userID, ttype, tokenColl)
	if err == nil {
		//token exists. so delete it
		DeleteToken(tokenOBJ.Token, tokenColl)
	}

	//generate and store new token
	tokenOBJ = TOBJ{uuid.New().String(), userID, ttype}
	_, err = tokenColl.InsertOne(context.TODO(), tokenOBJ)
	if err != nil {
		fmt.Println("Failed to store token:", err.Error())
	}

	return tokenOBJ.Token
}

//GetToken ...
func GetToken(userID string, ttype string, tokenColl *mongo.Collection) (TOBJ, error) {
	filter := bson.M{"user_id": userID, "type": ttype}

	var tokenOBJ TOBJ

	err := tokenColl.FindOne(context.TODO(), filter).Decode(&tokenOBJ)

	return tokenOBJ, err
}

//DeleteToken ...
func DeleteToken(token string, tokenColl *mongo.Collection) {
	filter := bson.M{"token": token}
	_, err := tokenColl.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println("Failed to delete token:", err.Error())
	}
}
