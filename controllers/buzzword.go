package controllers

import (
	"tech-buzzword-service/models"
	"tech-buzzword-service/util"

	"github.com/gin-gonic/gin"

	"net/http"
)

type BuzzwordController struct{}

var buzzwordModel = new(models.Buzzword)

func (b BuzzwordController) RetrieveBuzzword(c *gin.Context) {
	buzzword := buzzwordModel.GetBuzzword()
	if buzzword == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve buzzword", "error": "No buzzword"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Buzzword of the day!", "buzzword": buzzword})
}

func (b BuzzwordController) RetrievePreviousBuzzwords(c *gin.Context) {
	buzzwords := buzzwordModel.GetPreviousBuzzwords()
	if buzzwords == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve previous buzzwords", "error": "error"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Previous Buzzwords!", "previous_buzzwords": buzzwords})
}

func (b BuzzwordController) SetNewBuzzword(c *gin.Context) {
	currentBuzzword := buzzwordModel.GetBuzzword().Buzzword
	buzzwordModel.SetNewBuzzword()
	newBuzzword := buzzwordModel.GetBuzzword().Buzzword

	if util.SecureCompare(currentBuzzword, newBuzzword) == 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Buzzword didn't update", "error": "error"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Buzzwords updated"})
}
