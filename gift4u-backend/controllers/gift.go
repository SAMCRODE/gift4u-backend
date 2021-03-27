package controllers

import (
	"net/http"
	"strconv"

	"github.com/SAMCRODE/gift4u-backend/models"

	"github.com/gin-gonic/gin"
)

type GiftController struct{}

var giftModel = new(models.Gift)

// Gift godoc
// @Summary Save gift
// @Description save a new gift
// @Param   gift body models.Gift true "gift object"
// @Accept  json
// @Produce  json
// @Success 200
// @Router /gift/create [post]
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

// Gift godoc
// @Summary Search gift
// @Description Search gift identified by id
// @Param   id path int true "Gift id"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Gift
// @Router /gift/search/:id [get]
func (g GiftController) SearchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please try again later"})
		return
	}

	gift, err := models.SearchGiftById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": gift})
}

// Gift godoc
// @Summary Search a collection of gift given a page
// @Description Search gifts page
// @Param   page path int true "Page to be retrieved"
// @Accept  json
// @Produce  json
// @Success 200 {array} []models.Gift
// @Router /gift/collection/:page [get]
func (g GiftController) SearchGiftsByPage(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please try again later1"})
		return
	}

	gifts, err := models.SearchGiftsPaginated(page)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later2"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"gifts": gifts})
}
