package controllers

import (
	"gift4u/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GiftController struct{}

var giftModel = new(models.Gift)

func (g GiftController) Save(c *gin.Context) {
	gift := models.Gift{}

	if errA := c.ShouldBind(&gift); errA != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a valid gift object"})
		return
	}

	err := gift.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Gift Save"})
}
