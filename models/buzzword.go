package models

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"tech-buzzword-service/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var buzzword *Buzzword
var prevBuzzwords []*Buzzword

type Buzzword struct {
	ID          primitive.ObjectID `bson:"_id"`
	Buzzword    string             `bson:"Buzzword"`
	Definition  string             `bson:"Definition"`
	Examples    []string           `bson:"Examples"`
	HasBeenSaid bool               `bson:"HasBeenSaid"`
	Date        time.Time          `bson:"Date"`
}

func Init() {
	coll := db.GetColl()
	InitBuzzword(coll)
	InitPreviousBuzzwords(coll)
}

func InitBuzzword(coll *mongo.Collection) {
	fmt.Println("Initializing Buzzword")
	ifBuzzwordTodayExists := bson.D{{
		Key:   "Date",
		Value: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local),
	},
		{
			Key:   "HasBeenSaid",
			Value: true,
		}}
	var result *Buzzword
	err := coll.FindOne(context.TODO(), ifBuzzwordTodayExists).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No documents found, need to set new Buzzword")
			randomBuzz, randNumErr := GetRandomBuzzword(coll)
			if randNumErr != nil {
				panic(randNumErr)
			}
			result = randomBuzz
		} else {
			panic(err)
		}
	} else {
		fmt.Printf("Found word for today: %s\n", result.Buzzword)
	}
	buzzword = result
}

func InitPreviousBuzzwords(coll *mongo.Collection) {
	fmt.Println("Initializing Previous Buzzwords")
	prevBuzzwordsFilter := bson.D{
		{
			Key:   "HasBeenSaid",
			Value: true,
		},
		{
			Key:   "Date",
			Value: bson.D{{Key: "$lte", Value: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)}},
		},
	}

	cursor, err := coll.Find(context.TODO(), prevBuzzwordsFilter)
	if err != nil {
		panic(err)
	}
	var result []*Buzzword
	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}
	prevBuzzwords = result
}

func (b Buzzword) GetBuzzword() *Buzzword {
	return buzzword
}

func (b Buzzword) GetPreviousBuzzwords() []*Buzzword {
	return prevBuzzwords
}

func GetPotentialBuzzwords(coll *mongo.Collection) []Buzzword {
	fmt.Println("Getting list of potential buzzwords")
	filter := bson.D{{Key: "HasBeenSaid", Value: false}}
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
	fmt.Println("Update New Buzzword in Mongo")
	filter := bson.D{{Key: "_id", Value: word.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "HasBeenSaid", Value: true}, {Key: "Date", Value: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)}}}}
	result, _ := coll.UpdateOne(context.TODO(), filter, update)
	fmt.Printf("Documents matched: %v\n", result.MatchedCount)
}

func GetRandomBuzzword(coll *mongo.Collection) (*Buzzword, error) {
	fmt.Println("Getting new random buzzword")
	possibleBuzzwords := GetPotentialBuzzwords(coll)
	if len(possibleBuzzwords) < 1 {
		return nil, errors.New("area calculation failed, radius is less than zero")
	}
	randomNum := rand.Intn(len(possibleBuzzwords))
	newBuzzword := &possibleBuzzwords[randomNum]
	fmt.Printf("New Buzzword: %s\n", newBuzzword.Buzzword)
	UpdateNewBuzzword(coll, *newBuzzword)
	return newBuzzword, nil
}
