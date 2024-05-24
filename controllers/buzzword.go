package controllers

import (
	"context"
	"fmt"
	"math/rand"
	"tech-buzzword-service/db"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"net/http"
)

var buzzword *Buzzword

type Buzzword struct {
	ID          primitive.ObjectID `bson:"_id"`
	Buzzword    string             `bson:"buzzword"`
	Definition  string             `bson:"definition"`
	Examples    []string           `bson:"examples"`
	HasBeenSaid bool               `bson:"hasBeenSaid"`
	Date        time.Time          `bson:"Date"`
}

func (b Buzzword) Retrieve(c *gin.Context) {

	if buzzword == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve buzzword", "error": "No buzzword"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Buzzword of the day!", "buzzword": buzzword})
}

func InitBuzzword() {
	fmt.Println("Initializing Cron")
	coll := db.GetColl()

	ifBuzzwordTodayExists := bson.D{{Key: "date", Value: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)}}
	var result *Buzzword
	err := coll.FindOne(context.TODO(), ifBuzzwordTodayExists).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No documents found, need to set new Buzzword")
			result = GetRandomBuzzword(coll)
		} else {
			panic(err)
		}
	} else {
		fmt.Printf("Found word for today: %s\n", result.Buzzword)
		buzzword = result
	}

}

func GetBuzzword() *Buzzword {
	return buzzword
}

func GetPotentialBuzzwords(coll *mongo.Collection) []Buzzword {
	filter := bson.D{{Key: "hasBeenSaid", Value: false}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []Buzzword
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func UpdateNewBuzzword(coll *mongo.Collection, word Buzzword) {
	filter := bson.D{{Key: "_id", Value: word.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "hasBeenSaid", Value: true}, {Key: "date", Value: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)}}}}
	result, _ := coll.UpdateOne(context.TODO(), filter, update)
	fmt.Printf("Documents matched: %v\n", result.MatchedCount)
}

func GetRandomBuzzword(coll *mongo.Collection) *Buzzword {
	possibleBuzzwords := GetPotentialBuzzwords(coll)
	randomNum := rand.Intn(len(possibleBuzzwords) - 1)
	newBuzzword := &possibleBuzzwords[randomNum]
	fmt.Printf("New Buzzword: %s\n ", newBuzzword.Buzzword)
	UpdateNewBuzzword(coll, *newBuzzword)
	return newBuzzword
}
