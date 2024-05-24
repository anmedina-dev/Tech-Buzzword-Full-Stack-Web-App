package controllers

import (
	"fmt"
	"tech-buzzword-service/buzzword"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"net/http"
)

type Buzzword struct {
	ID          primitive.ObjectID `bson:"_id"`
	Buzzword    string             `bson:"buzzword"`
	Definition  string             `bson:"definition"`
	Examples    []string           `bson:"examples"`
	HasBeenSaid bool               `bson:"hasBeenSaid"`
	Date        time.Time          `bson:"Date"`
}

func (b Buzzword) RetrieveBuzzword(c *gin.Context) {
	buzzword := buzzword.GetBuzzword()
	if buzzword == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve buzzword", "error": "No buzzword"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Buzzword of the day!", "buzzword": buzzword})
}

func (b Buzzword) RetrievePreviousBuzzwords(c *gin.Context) {
	buzzwords := buzzword.GetPreviousBuzzwords()
	if buzzwords == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve previous buzzwords", "error": "error"})
		c.Abort()
		return
	}
	fmt.Println(buzzwords)
	c.JSON(http.StatusOK, gin.H{"message": "Previous Buzzwords!", "previous_buzzwords": buzzwords})
}
